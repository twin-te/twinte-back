package timetableusecase

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/idtype"
	timetablemodule "github.com/twin-te/twinte-back/module/timetable"
	timetableentity "github.com/twin-te/twinte-back/module/timetable/entity"
	timetableport "github.com/twin-te/twinte-back/module/timetable/port"
)

func (uc *Impl) CreateRegisteredCoursesByCodes(ctx context.Context, year int, codes []timetableentity.Code) error {
	user, err := uc.a.AuthorizeAuthenticatedUser(ctx)
	if err != nil {
		return err
	}

	courses, err := uc.r.ListCourses(ctx, timetableport.ListCoursesConds{
		Year:  &year,
		Codes: &codes,
	})
	if err != nil {
		return err
	}

	codeToCourseMap := lo.SliceToMap(courses, func(course *timetableentity.Course) (timetableentity.Code, *timetableentity.Course) {
		return course.Code, course
	})

	registeredCourses := make([]*timetableentity.RegisteredCourse, 0, len(codes))

	for _, code := range codes {
		course, ok := codeToCourseMap[code]
		if !ok {
			return fmt.Errorf("not found course with code %s", code)
		}
		registeredCourses = append(registeredCourses, timetableentity.NewRegisteredCourseFromCourse(user.ID, course))
	}

	return uc.r.SaveRegisteredCourses(ctx, registeredCourses)
}

func (uc *Impl) CreateRegisteredCourseManually(ctx context.Context, in timetablemodule.CreateCourseManuallyIn) error {
	user, err := uc.a.AuthorizeAuthenticatedUser(ctx)
	if err != nil {
		return err
	}

	registeredCourse := &timetableentity.RegisteredCourse{
		ID:          idtype.NewRegisteredCourseID(),
		UserID:      user.ID,
		Year:        in.Year,
		Name:        &in.Name,
		Instructors: &in.Instructors,
		Cregit:      &in.Cregit,
		Methods:     &in.Methods,
		Schedules:   &in.Schedules,
	}

	return uc.r.SaveRegisteredCourse(ctx, registeredCourse)
}

func (uc Impl) GetRegisteredCourses(ctx context.Context, year *int) ([]*timetableentity.RegisteredCourse, error) {
	user, err := uc.a.AuthorizeAuthenticatedUser(ctx)
	if err != nil {
		return nil, err
	}

	return uc.r.ListRegisteredCourses(ctx, timetableport.ListRegisteredCoursesConds{
		UserID: &user.ID,
		Year:   year,
	})
}

func (uc Impl) DeleteRegisteredCourse(ctx context.Context, id idtype.RegisteredCourseID) error {
	_, err := uc.a.AuthorizeAuthenticatedUser(ctx)
	if err != nil {
		return err
	}

	return uc.r.DeleteRegisteredCourse(ctx, id)
}
