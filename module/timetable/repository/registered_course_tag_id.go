package timetablerepository

import (
	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/base"
	"github.com/twin-te/twinte-back/db/gen/model"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
	"gorm.io/gorm"
)

func (r *impl) updateRegisteredCourseTagIDs(db *gorm.DB, registeredCourse *timetabledomain.RegisteredCourse) error {
	toCreate, toDelete := lo.Difference(registeredCourse.TagIDs, registeredCourse.EntityBeforeUpdated.TagIDs)

	if len(toCreate) != 0 {
		dbTags := base.MapWithArg(toCreate, registeredCourse.ID, toDBRegisteredCourseTag)

		if err := db.Create(dbTags).Error; err != nil {
			return err
		}
	}

	if len(toDelete) != 0 {
		if err := db.Where("registered_course = ?", registeredCourse.ID.String()).
			Where("tag IN ?", base.MapByString(toDelete)).
			Delete(&model.RegisteredCourseTag{}).
			Error; err != nil {
			return err
		}
	}

	return nil
}

func fromDBRegisteredCourseTag(dbRegisteredCourseTag model.RegisteredCourseTag) (idtype.TagID, error) {
	return idtype.ParseTagID(dbRegisteredCourseTag.Tag)
}

func toDBRegisteredCourseTag(tagID idtype.TagID, registeredCourseID idtype.RegisteredCourseID) model.RegisteredCourseTag {
	return model.RegisteredCourseTag{
		Tag:              tagID.String(),
		RegisteredCourse: registeredCourseID.String(),
	}
}
