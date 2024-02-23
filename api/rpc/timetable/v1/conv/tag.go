package timetablev1conv

import (
	sharedconv "github.com/twin-te/twinte-back/api/rpc/shared/conv"
	timetablev1 "github.com/twin-te/twinte-back/api/rpcgen/timetable/v1"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
)

func ToPBTag(tag *timetabledomain.Tag) *timetablev1.Tag {
	pbTag := &timetablev1.Tag{
		Id:       sharedconv.ToPBUUID(tag.ID),
		UserId:   sharedconv.ToPBUUID(tag.UserID),
		Name:     tag.Name.String(),
		Position: int32(tag.Position),
	}
	return pbTag
}
