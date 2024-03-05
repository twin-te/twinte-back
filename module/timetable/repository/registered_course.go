package timetablerepository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/base"
	"github.com/twin-te/twinte-back/db/gen/model"
	dbhelper "github.com/twin-te/twinte-back/db/helper"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
	timetableport "github.com/twin-te/twinte-back/module/timetable/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *impl) FindRegisteredCourse(ctx context.Context, conds timetableport.FindRegisteredCourseConds, lock sharedport.Lock) (*timetabledomain.RegisteredCourse, error) {
	db := r.db.WithContext(ctx).Where("id = ?", conds.ID.String())

	if conds.UserID != nil {
		db = db.Where("user_id = ?", conds.UserID.String())
	}

	db = db.Clauses(clause.Locking{
		Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
		Table:    clause.Table{Name: clause.CurrentTable},
	})

	dbRegisteredCourse := new(model.RegisteredCourse)

	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Preload("Tags").Take(dbRegisteredCourse).Error
		return dbhelper.ConvertErrRecordNotFound(err)
	})
	if err != nil {
		return nil, err
	}

	return fromDBRegisteredCourse(dbRegisteredCourse)
}

func (r *impl) ListRegisteredCourses(ctx context.Context, conds timetableport.ListRegisteredCoursesConds, lock sharedport.Lock) ([]*timetabledomain.RegisteredCourse, error) {
	db := r.db.WithContext(ctx)

	if conds.UserID != nil {
		db = db.Where("user_id = ?", conds.UserID.String())
	}

	if conds.Year != nil {
		db = db.Where("year = ?", conds.Year.Int())
	}

	if conds.CourseIDs != nil {
		db = db.Where("course_id IN ?", base.MapByString(*conds.CourseIDs))
	}

	db = db.Clauses(clause.Locking{
		Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
		Table:    clause.Table{Name: clause.CurrentTable},
	})

	var dbRegisteredCourses []*model.RegisteredCourse

	err := db.Transaction(func(tx *gorm.DB) error {
		return tx.Preload("Tags").Find(&dbRegisteredCourses).Error
	})
	if err != nil {
		return nil, err
	}

	return base.MapWithErr(dbRegisteredCourses, fromDBRegisteredCourse)
}

func (r *impl) CreateRegisteredCourses(ctx context.Context, registeredCourses ...*timetabledomain.RegisteredCourse) error {
	dbRegisteredCourses, err := base.MapWithArgAndErr(registeredCourses, true, toDBRegisteredCourse)
	if err != nil {
		return err
	}
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(dbRegisteredCourses).Error
	}, nil)
}

func (r *impl) UpdateRegisteredCourse(ctx context.Context, registeredCourse *timetabledomain.RegisteredCourse) error {
	cols := make([]string, 0)

	if registeredCourse.UserID != registeredCourse.EntityBeforeUpdated.UserID {
		cols = append(cols, "user_id")
	}

	if registeredCourse.Year != registeredCourse.EntityBeforeUpdated.Year {
		cols = append(cols, "year")
	}

	if !base.EqualPtr(registeredCourse.CourseID, registeredCourse.EntityBeforeUpdated.CourseID) {
		cols = append(cols, "course_id")
	}

	if !base.EqualPtr(registeredCourse.Name, registeredCourse.EntityBeforeUpdated.Name) {
		cols = append(cols, "name")
	}

	if !base.EqualPtr(registeredCourse.Instructors, registeredCourse.EntityBeforeUpdated.Instructors) {
		cols = append(cols, "instractor")
	}

	if !base.EqualPtr(registeredCourse.Credit, registeredCourse.EntityBeforeUpdated.Credit) {
		cols = append(cols, "credit")
	}

	if !base.EqualSlicePtr(registeredCourse.Methods, registeredCourse.EntityBeforeUpdated.Methods) {
		cols = append(cols, "methods")
	}

	if !base.EqualSlicePtr(registeredCourse.Schedules, registeredCourse.EntityBeforeUpdated.Schedules) {
		cols = append(cols, "schedules")
	}

	if registeredCourse.Memo != registeredCourse.EntityBeforeUpdated.Memo {
		cols = append(cols, "memo")
	}

	if registeredCourse.Attendance != registeredCourse.EntityBeforeUpdated.Attendance {
		cols = append(cols, "attendance")
	}

	if registeredCourse.Absence != registeredCourse.EntityBeforeUpdated.Absence {
		cols = append(cols, "absence")
	}

	if registeredCourse.Late != registeredCourse.EntityBeforeUpdated.Late {
		cols = append(cols, "late")
	}

	dbRegisteredCourse, err := toDBRegisteredCourse(registeredCourse, false)
	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(cols) > 0 {
			if err := tx.Select(cols).Updates(dbRegisteredCourse).Error; err != nil {
				return err
			}
		}
		return r.updateRegisteredCourseTagIDs(tx, registeredCourse)
	}, nil)
}

func (r *impl) DeleteRegisteredCourses(ctx context.Context, conds timetableport.DeleteRegisteredCoursesConds) (rowsAffected int, err error) {
	db := r.db.WithContext(ctx)

	if conds.ID != nil {
		db = db.Where("id = ?", conds.ID.String())
	}

	if conds.UserID != nil {
		db = db.Where("user_id = ?", conds.UserID.String())
	}

	return int(db.Delete(&model.RegisteredCourse{}).RowsAffected), db.Error
}

func fromDBRegisteredCourse(dbRegisteredCourse *model.RegisteredCourse) (*timetabledomain.RegisteredCourse, error) {
	return timetabledomain.ConstructRegisteredCourse(func(registeredCourse *timetabledomain.RegisteredCourse) (err error) {
		registeredCourse.ID, err = idtype.ParseRegisteredCourseID(dbRegisteredCourse.ID)
		if err != nil {
			return err
		}

		registeredCourse.UserID, err = idtype.ParseUserID(dbRegisteredCourse.UserID)
		if err != nil {
			return err
		}

		registeredCourse.Year, err = shareddomain.ParseAcademicYear(int(dbRegisteredCourse.Year))
		if err != nil {
			return err
		}

		if dbRegisteredCourse.CourseID != nil {
			courseID, err := idtype.ParseCourseID(*dbRegisteredCourse.CourseID)
			if err != nil {
				return err
			}
			registeredCourse.CourseID = &courseID
		}

		if dbRegisteredCourse.Name != nil {
			name, err := timetabledomain.ParseName(*dbRegisteredCourse.Name)
			if err != nil {
				return err
			}
			registeredCourse.Name = &name
		}

		registeredCourse.Instructors = dbRegisteredCourse.Instractor

		if dbRegisteredCourse.Credit != nil {
			credit, err := timetabledomain.ParseCredit(fmt.Sprintf("%.1f", *dbRegisteredCourse.Credit))
			if err != nil {
				return err
			}
			registeredCourse.Credit = &credit
		}

		if dbRegisteredCourse.Methods != nil {
			methods, err := base.MapWithErr(*dbRegisteredCourse.Methods, timetabledomain.ParseCourseMethod)
			if err != nil {
				return err
			}
			registeredCourse.Methods = &methods
		}

		if dbRegisteredCourse.Schedules != nil {
			schedules, err := fromDBRegisteredCourseSchedules(*dbRegisteredCourse.Schedules)
			if err != nil {
				return err
			}
			registeredCourse.Schedules = &schedules
		}

		registeredCourse.Memo = dbRegisteredCourse.Memo

		registeredCourse.Attendance, err = timetabledomain.ParseAttendance(int(dbRegisteredCourse.Attendance))
		if err != nil {
			return
		}

		registeredCourse.Absence, err = timetabledomain.ParseAbsence(int(dbRegisteredCourse.Absence))
		if err != nil {
			return
		}

		registeredCourse.Late, err = timetabledomain.ParseLate(int(dbRegisteredCourse.Late))
		if err != nil {
			return
		}

		registeredCourse.TagIDs, err = base.MapWithErr(dbRegisteredCourse.Tags, fromDBRegisteredCourseTag)
		if err != nil {
			return err
		}

		return nil
	})
}

func toDBRegisteredCourse(registeredCourse *timetabledomain.RegisteredCourse, withAssociations bool) (*model.RegisteredCourse, error) {
	dbRegisteredCourse := &model.RegisteredCourse{
		ID:         registeredCourse.ID.String(),
		UserID:     registeredCourse.UserID.String(),
		Year:       int16(registeredCourse.Year),
		CourseID:   registeredCourse.CourseID.StringPtr(),
		Instractor: registeredCourse.Instructors,
		Memo:       registeredCourse.Memo,
		Attendance: int32(registeredCourse.Attendance.Int()),
		Absence:    int32(registeredCourse.Absence.Int()),
		Late:       int32(registeredCourse.Late.Int()),
	}

	if registeredCourse.Name != nil {
		dbRegisteredCourse.Name = lo.ToPtr(registeredCourse.Name.String())
	}

	if registeredCourse.Credit != nil {
		dbRegisteredCourse.Credit = lo.ToPtr(registeredCourse.Credit.Float())
	}

	if registeredCourse.Methods != nil {
		dbRegisteredCourse.Methods = lo.ToPtr(base.MapByString(*registeredCourse.Methods))
	}

	if registeredCourse.Schedules != nil {
		dbRegisteredCourseSchedules, err := toDBRegisteredCourseSchedulesJSON(*registeredCourse.Schedules)
		if err != nil {
			return nil, err
		}
		dbRegisteredCourse.Schedules = &dbRegisteredCourseSchedules
	}

	if withAssociations {
		dbRegisteredCourse.Tags = base.MapWithArg(registeredCourse.TagIDs, registeredCourse.ID, toDBRegisteredCourseTag)
	}

	return dbRegisteredCourse, nil
}

type dbRegisteredCourseSchedule struct {
	Module string `json:"module"`
	Day    string `json:"day"`
	Period int    `json:"period"`
	Room   string `json:"room"`
}

func fromDBRegisteredCourseSchedules(data []byte) ([]timetabledomain.Schedule, error) {
	var dbRegisteredCourseSchedules []dbRegisteredCourseSchedule

	if err := json.Unmarshal(data, &dbRegisteredCourseSchedules); err != nil {
		return nil, err
	}

	return base.MapWithErr(dbRegisteredCourseSchedules, func(dbRegisteredCourseSchedule dbRegisteredCourseSchedule) (timetabledomain.Schedule, error) {
		return timetabledomain.ParseSchedule(
			dbRegisteredCourseSchedule.Module,
			dbRegisteredCourseSchedule.Day,
			dbRegisteredCourseSchedule.Period,
			dbRegisteredCourseSchedule.Room,
		)
	})
}

func toDBRegisteredCourseSchedulesJSON(schedules []timetabledomain.Schedule) ([]byte, error) {
	dbRegisteredCourseSchedules := base.Map(schedules, func(schedule timetabledomain.Schedule) *dbRegisteredCourseSchedule {
		return &dbRegisteredCourseSchedule{
			Module: schedule.Module.String(),
			Day:    schedule.Day.String(),
			Period: schedule.Period.Int(),
			Room:   schedule.Rooms,
		}
	})
	return json.Marshal(dbRegisteredCourseSchedules)
}
