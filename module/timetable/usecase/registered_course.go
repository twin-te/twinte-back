package timetableusecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/apperr"
	"github.com/twin-te/twinte-back/base"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	sharederr "github.com/twin-te/twinte-back/module/shared/err"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
	timetablemodule "github.com/twin-te/twinte-back/module/timetable"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
	timetableport "github.com/twin-te/twinte-back/module/timetable/port"
)

func (uc *impl) CreateRegisteredCoursesByCodes(ctx context.Context, year shareddomain.AcademicYear, codes []timetabledomain.Code) ([]*timetabledomain.RegisteredCourse, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	courses, err := uc.r.ListCourses(ctx, timetableport.ListCoursesConds{
		Year:  &year,
		Codes: &codes,
	}, sharedport.LockNone)
	if err != nil {
		return nil, err
	}

	courseIDs := base.Map(courses, func(course *timetabledomain.Course) idtype.CourseID {
		return course.ID
	})

	courseIDToCode := lo.SliceToMap(courses, func(course *timetabledomain.Course) (idtype.CourseID, timetabledomain.Code) {
		return course.ID, course.Code
	})

	savedTargetRegisteredCourses, err := uc.r.ListRegisteredCourses(ctx, timetableport.ListRegisteredCoursesConds{
		UserID:    &userID,
		Year:      &year,
		CourseIDs: &courseIDs,
	}, sharedport.LockNone)
	if err != nil {
		return nil, err
	}

	if len(savedTargetRegisteredCourses) != 0 {
		msg := "the courses with these codes are already registered: "
		for i, savedsavedTargetRegisteredCourse := range savedTargetRegisteredCourses {
			if i != 0 {
				msg += ","
			}
			msg += courseIDToCode[*savedsavedTargetRegisteredCourse.CourseID].String()
		}
		return nil, apperr.New(sharederr.CodeAlreadyExists, msg)
	}

	codeToCourse := lo.SliceToMap(courses, func(course *timetabledomain.Course) (timetabledomain.Code, *timetabledomain.Course) {
		return course.Code, course
	})

	registeredCourses := make([]*timetabledomain.RegisteredCourse, 0, len(codes))

	for _, code := range codes {
		course, ok := codeToCourse[code]
		if !ok {
			return nil, apperr.New(sharederr.CodeNotFound, fmt.Sprintf("not found course with code %s", code))
		}
		registeredCourse, err := uc.f.NewRegisteredCourseFromCourse(userID, course)
		if err != nil {
			return nil, err
		}
		registeredCourses = append(registeredCourses, registeredCourse)
	}

	err = uc.r.Transaction(ctx, func(rtx timetableport.Repository) error {
		return rtx.CreateRegisteredCourses(ctx, registeredCourses...)
	})
	return registeredCourses, err
}

func (uc *impl) CreateRegisteredCourseManually(ctx context.Context, in timetablemodule.CreateRegisteredCourseManuallyIn) (*timetabledomain.RegisteredCourse, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	registeredCourse, err := uc.f.NewRegisteredCourseMannualy(
		userID,
		in.Year,
		in.Name,
		in.Instructors,
		in.Credit,
		in.Methods,
		in.Schedules,
	)
	if err != nil {
		return nil, err
	}

	return registeredCourse, uc.r.CreateRegisteredCourses(ctx, registeredCourse)
}

func (uc impl) GetRegisteredCourseByID(ctx context.Context, id idtype.RegisteredCourseID) (*timetabledomain.RegisteredCourse, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	registeredCourse, err := uc.r.FindRegisteredCourse(ctx, timetableport.FindRegisteredCourseConds{
		ID:     id,
		UserID: &userID,
	}, sharedport.LockNone)
	if err != nil {
		if errors.Is(err, sharedport.ErrNotFound) {
			return nil, apperr.New(sharederr.CodeNotFound, fmt.Sprintf("not found registered course whose id is %s", registeredCourse.ID))
		}
		return nil, err
	}

	return registeredCourse, uc.r.LoadCourseToRegisteredCourse(ctx, []*timetabledomain.RegisteredCourse{registeredCourse}, sharedport.LockNone)
}

func (uc impl) GetRegisteredCourses(ctx context.Context, year *shareddomain.AcademicYear) ([]*timetabledomain.RegisteredCourse, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	registeredCourses, err := uc.r.ListRegisteredCourses(ctx, timetableport.ListRegisteredCoursesConds{
		UserID: &userID,
		Year:   year,
	}, sharedport.LockNone)
	if err != nil {
		return nil, err
	}

	return registeredCourses, uc.r.LoadCourseToRegisteredCourse(ctx, registeredCourses, sharedport.LockNone)
}

func (uc impl) UpdateRegisteredCourse(ctx context.Context, in timetablemodule.UpdateRegisteredCourseIn) (registeredCourse *timetabledomain.RegisteredCourse, err error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	err = uc.r.Transaction(ctx, func(rtx timetableport.Repository) (err error) {
		registeredCourse, err = rtx.FindRegisteredCourse(ctx, timetableport.FindRegisteredCourseConds{
			ID:     in.ID,
			UserID: &userID,
		}, sharedport.LockExclusive)
		if err != nil {
			if errors.Is(err, sharedport.ErrNotFound) {
				return apperr.New(sharederr.CodeNotFound, fmt.Sprintf("not found registered course whose id is %s", in.ID))
			}
			return err
		}

		if err := uc.r.LoadCourseToRegisteredCourse(ctx, []*timetabledomain.RegisteredCourse{registeredCourse}, sharedport.LockNone); err != nil {
			return err
		}

		if in.TagIDs != nil {
			tags, err := rtx.ListTags(ctx, timetableport.ListTagsConds{
				UserID: &userID,
			}, sharedport.LockShared)
			if err != nil {
				return err
			}

			for _, tagID := range *in.TagIDs {
				if !lo.ContainsBy(tags, func(tag *timetabledomain.Tag) bool {
					return tagID == tag.ID
				}) {
					return apperr.New(sharederr.CodeInvalidArgument, fmt.Sprintf("the tag whose id is %s is not found", tagID))
				}
			}
		}

		registeredCourse.BeforeUpdateHook()

		registeredCourse.Update(timetabledomain.RegisteredCourseDataToUpdate{
			Name:        in.Name,
			Instructors: in.Instructors,
			Credit:      in.Credit,
			Methods:     in.Methods,
			Schedules:   in.Schedules,
			Memo:        in.Memo,
			Attendance:  in.Attendance,
			Absence:     in.Absence,
			Late:        in.Late,
			TagIDs:      in.TagIDs,
		})

		return rtx.UpdateRegisteredCourse(ctx, registeredCourse)
	})

	return
}

func (uc impl) DeleteRegisteredCourse(ctx context.Context, id idtype.RegisteredCourseID) error {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	rowsAffected, err := uc.r.DeleteRegisteredCourses(ctx, timetableport.DeleteRegisteredCoursesConds{
		ID:     &id,
		UserID: &userID,
	})
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperr.New(sharederr.CodeNotFound, fmt.Sprintf("not found registered course whose id is %s", id))
	}

	return nil
}
