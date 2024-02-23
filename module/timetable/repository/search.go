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

var (
	cache           = make(map[shareddomain.AcademicYear][]*timetabledomain.Course)
	mu              sync.Mutex
	CourseCacheTime time.Duration = time.Duration(appenv.COURSE_CACHE_HOURS) * time.Hour
)

func (r *impl) GetCoursesWithCache(ctx context.Context, year shareddomain.AcademicYear) (courses []*timetabledomain.Course, err error) {
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

	totalNumCourses := lo.Reduce(lo.Keys(cache), func(totalNumCourses int, year shareddomain.AcademicYear, _ int) int {
		return totalNumCourses + len(cache[year])
	}, 0)

	if totalNumCourses+len(courses) > 100_000 {
		minYear := lo.Min(lo.Keys(cache))
		delete(cache, minYear)
	}

	cache[year] = courses

	go func() {
		time.Sleep(CourseCacheTime)
		mu.Lock()
		delete(cache, year)
		mu.Unlock()
	}()

	return
}

func (r *impl) SearchCourses(ctx context.Context, conds timetableport.SearchCoursesConds) ([]*timetabledomain.Course, error) {
	courses, err := r.GetCoursesWithCache(ctx, conds.Year)
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
		return lo.Every(conds.Schedules.FullyIncluded, course.Schedules)
	})
	courses = lo.Filter(courses, func(course *timetabledomain.Course, _ int) bool {
		intersect := lo.Intersect(course.Schedules, conds.Schedules.PartiallyOverlapped)
		return len(intersect) > 0
	})

	// Apply offset
	courses = courses[lo.Clamp(conds.Offset, 0, len(courses)):]

	// Apply limit
	courses = courses[:lo.Clamp(conds.Limit, 0, len(courses))]

	return base.MapByClone(courses), nil
}
