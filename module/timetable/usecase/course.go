package timetableusecase

import (
	"context"
	"errors"

	authdomain "github.com/twin-te/twinte-back/module/auth/domain"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
	timetablemodule "github.com/twin-te/twinte-back/module/timetable"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
	timetableport "github.com/twin-te/twinte-back/module/timetable/port"
)

func (uc *impl) GetCoursesByIDs(ctx context.Context, ids []idtype.CourseID) ([]*timetabledomain.Course, error) {
	return uc.r.ListCourses(ctx, timetableport.ListCoursesConds{
		IDs: &ids,
	}, sharedport.LockNone)
}

func (uc *impl) GetCoursesByCodes(ctx context.Context, year shareddomain.AcademicYear, codes []timetabledomain.Code) ([]*timetabledomain.Course, error) {
	return uc.r.ListCourses(ctx, timetableport.ListCoursesConds{
		Year:  &year,
		Codes: &codes,
	}, sharedport.LockNone)
}

func (uc *impl) SearchCourses(ctx context.Context, in timetablemodule.SearchCoursesIn) ([]*timetabledomain.Course, error) {
	return uc.r.SearchCourses(ctx, timetableport.SearchCoursesConds(in))
}

func (uc *impl) UpdateCoursesBasedOnKDB(ctx context.Context, year shareddomain.AcademicYear) error {
	if err := uc.a.Authorize(ctx, authdomain.PermissionExecuteBatchJob); err != nil {
		return err
	}

	courseWithoutIDs, err := uc.g.GetCourseWithoutIDsFromKDB(ctx, year)
	if err != nil {
		return err
	}

	for _, courseWithoutID := range courseWithoutIDs {
		err = uc.r.Transaction(ctx, func(rtx timetableport.Repository) error {
			savedCourse, err := rtx.FindCourse(ctx, timetableport.FindCourseConds{
				Year: year,
				Code: courseWithoutID.Code,
			}, sharedport.LockExclusive)

			isErrNotFound := errors.Is(err, sharedport.ErrNotFound)

			if err != nil && !isErrNotFound {
				return err
			}

			if isErrNotFound {
				newCourse, err := uc.f.NewCourse(courseWithoutID)
				if err != nil {
					return err
				}
				return rtx.CreateCourses(ctx, newCourse)
			}

			if courseWithoutID.LastUpdatedAt.After(savedCourse.LastUpdatedAt) {
				savedCourse.BeforeUpdateHook()
				savedCourse.Update(timetabledomain.CourseDataToUpdate{
					Name:              &courseWithoutID.Name,
					Instructors:       &courseWithoutID.Instructors,
					Credit:            &courseWithoutID.Credit,
					Overview:          &courseWithoutID.Overview,
					Remarks:           &courseWithoutID.Remarks,
					LastUpdatedAt:     &courseWithoutID.LastUpdatedAt,
					HasParseError:     &courseWithoutID.HasParseError,
					IsAnnual:          &courseWithoutID.IsAnnual,
					RecommendedGrades: &courseWithoutID.RecommendedGrades,
					Methods:           &courseWithoutID.Methods,
					Schedules:         &courseWithoutID.Schedules,
				})
				return rtx.UpdateCourse(ctx, savedCourse)
			}

			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
