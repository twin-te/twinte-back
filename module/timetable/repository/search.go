package timetablerepository

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/appenv"
	"github.com/twin-te/twinte-back/base"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
	timetableport "github.com/twin-te/twinte-back/module/timetable/port"
)

func (r *impl) SearchCourses(ctx context.Context, conds timetableport.SearchCoursesConds) ([]*timetabledomain.Course, error) {
	courses, err := r.getCoursesWithCache(ctx, conds.Year)
	if err != nil {
		return nil, err
	}

	// Filter by keywords
	courses = lo.Filter(courses, func(course *timetabledomain.Course, _ int) bool {
		return lo.EveryBy(conds.Keywords, func(keyword string) bool {
			return strings.Contains(course.Name.String(), keyword)
		})
	})

	// Filter by code prefixes
	courses = lo.Filter(courses, func(course *timetabledomain.Course, _ int) bool {
		return lo.EveryBy(conds.CodePrefixes.Included, func(code string) bool {
			return strings.HasPrefix(course.Name.String(), code)
		})
	})
	courses = lo.Filter(courses, func(course *timetabledomain.Course, _ int) bool {
		return lo.EveryBy(conds.CodePrefixes.Excluded, func(code string) bool {
			return !strings.HasPrefix(course.Name.String(), code)
		})
	})

	// Filter by schedules
	courses = lo.Filter(courses, func(course *timetabledomain.Course, _ int) bool {
		return lo.EveryBy(conds.Schedules.FullyIncluded, func(s1 timetabledomain.Schedule) bool {
			return lo.SomeBy(course.Schedules, func(s2 timetabledomain.Schedule) bool {
				return s1.Module == s2.Module && s1.Day == s2.Day && s1.Period == s2.Period
			})
		})
	})
	courses = lo.Filter(courses, func(course *timetabledomain.Course, _ int) bool {
		return lo.SomeBy(conds.Schedules.PartiallyOverlapped, func(s1 timetabledomain.Schedule) bool {
			return lo.SomeBy(course.Schedules, func(s2 timetabledomain.Schedule) bool {
				return s1.Module == s2.Module && s1.Day == s2.Day && s1.Period == s2.Period
			})
		})
	})

	// Apply offset
	courses = courses[lo.Clamp(conds.Offset, 0, len(courses)):]

	// Apply limit
	courses = courses[:lo.Clamp(conds.Limit, 0, len(courses))]

	return base.MapByClone(courses), nil
}

var (
	cache                = make(map[shareddomain.AcademicYear][]*timetabledomain.Course)
	mu                   sync.Mutex
	courseCacheTime      time.Duration = time.Duration(appenv.COURSE_CACHE_HOURS) * time.Hour
	maxNumCoursesToCache               = 100_000
)

func (r *impl) getCoursesWithCache(ctx context.Context, year shareddomain.AcademicYear) (courses []*timetabledomain.Course, err error) {
	mu.Lock()
	defer mu.Unlock()

	courses, ok := cache[year]
	if ok {
		return
	}

	courses, err = r.ListCourses(ctx, timetableport.ListCoursesConds{
		Year: &year,
	}, sharedport.LockNone)
	if err != nil {
		return
	}

	cache[year] = courses

	for len(cache) != 0 {
		totalNumCourses := lo.Reduce(lo.Keys(cache), func(totalNumCourses int, year shareddomain.AcademicYear, _ int) int {
			return totalNumCourses + len(cache[year])
		}, 0)

		if totalNumCourses <= maxNumCoursesToCache {
			break
		}

		minYear := lo.Min(lo.Keys(cache))
		delete(cache, minYear)
	}

	go func() {
		time.Sleep(courseCacheTime)
		mu.Lock()
		delete(cache, year)
		mu.Unlock()
	}()

	return
}
