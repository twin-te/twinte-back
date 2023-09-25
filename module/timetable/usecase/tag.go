package timetableusecase

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/idtype"
	timetableentity "github.com/twin-te/twinte-back/module/timetable/entity"
	timetableport "github.com/twin-te/twinte-back/module/timetable/port"
)

func (uc Impl) CreateTag(ctx context.Context, name string) error {
	user, err := uc.a.AuthorizeAuthenticatedUser(ctx)
	if err != nil {
		return err
	}

	return uc.r.Transaction(ctx, func(rtx timetableport.Repository) error {
		tags, err := rtx.ListTags(ctx, timetableport.ListTagsConds{
			UserID: &user.ID,
		})
		if err != nil {
			return err
		}

		lastPosition := lo.Reduce(tags, func(lastPosition int, tag *timetableentity.Tag, index int) int {
			if lastPosition < tag.Position {
				lastPosition = tag.Position
			}
			return lastPosition
		}, 0)

		newTag := &timetableentity.Tag{
			ID:       idtype.NewTagID(),
			UserID:   user.ID,
			Name:     name,
			Position: lastPosition + 1,
		}

		return rtx.SaveTag(ctx, newTag)
	}, false)
}

func (uc Impl) GetTags(ctx context.Context) ([]*timetableentity.Tag, error) {
	user, err := uc.a.AuthorizeAuthenticatedUser(ctx)
	if err != nil {
		return nil, err
	}

	return uc.r.ListTags(ctx, timetableport.ListTagsConds{
		UserID: &user.ID,
	})
}

func (uc Impl) UpdateTagName(ctx context.Context, id idtype.TagID, name string) error {
	user, err := uc.a.AuthorizeAuthenticatedUser(ctx)
	if err != nil {
		return err
	}

	return uc.r.Transaction(ctx, func(rtx timetableport.Repository) error {
		tag, err := rtx.FindTag(ctx, timetableport.FindTagConds{
			ID:     &id,
			UserID: &user.ID,
		})

		if err != nil {
			return err
		}

		tag.Name = name
		return rtx.SaveTag(ctx, tag)
	}, false)
}

func (uc Impl) DeleteTag(ctx context.Context, id idtype.TagID) error {
	_, err := uc.a.AuthorizeAuthenticatedUser(ctx)
	if err != nil {
		return err
	}

	return uc.r.DeleteTag(ctx, id)
}

func (uc Impl) RearrangeTags(ctx context.Context, tagIDs idtype.TagIDs) error {
	user, err := uc.a.AuthorizeAuthenticatedUser(ctx)
	if err != nil {
		return err
	}

	return uc.r.Transaction(ctx, func(rtx timetableport.Repository) error {
		tags, err := rtx.ListTags(ctx, timetableport.ListTagsConds{
			UserID: &user.ID,
		})
		if err != nil {
			return err
		}

		if len(tagIDs) != len(tags) {
			return fmt.Errorf("invalid tag ids")
		}

		idToTagMap := lo.SliceToMap(tags, func(tag *timetableentity.Tag) (idtype.TagID, *timetableentity.Tag) {
			return tag.ID, tag
		})

		for position, id := range tagIDs {
			tag, ok := idToTagMap[id]
			if !ok {
				return fmt.Errorf("invalid tag id %s", id.String())
			}

			tag.Position = position
		}

		return rtx.SaveTags(ctx, tags)
	}, false)
}
