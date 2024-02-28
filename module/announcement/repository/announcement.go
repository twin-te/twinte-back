package announcementrepository

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/base"
	announcementdomain "github.com/twin-te/twinte-back/module/announcement/domain"
	announcementport "github.com/twin-te/twinte-back/module/announcement/port"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
)

func (r *impl) FindAnnouncement(ctx context.Context, conds announcementport.FindAnnouncementConds, lock sharedport.Lock) (*announcementdomain.Announcement, error) {
	announcement, ok := lo.Find(r.announcements, func(announcement *announcementdomain.Announcement) bool {
		return conds.ID == announcement.ID
	})
	if !ok {
		return nil, sharedport.ErrNotFound
	}

	if conds.PublishedAtBefore != nil && !announcement.PublishedAt.Before(*conds.PublishedAtBefore) {
		return nil, sharedport.ErrNotFound
	}

	return announcement.Clone(), nil
}

func (r *impl) ListAnnouncements(ctx context.Context, conds announcementport.ListAnnouncementsConds, lock sharedport.Lock) ([]*announcementdomain.Announcement, error) {
	announcements := r.announcements

	if conds.PublishedAtBefore != nil {
		announcements = lo.Filter(announcements, func(announcement *announcementdomain.Announcement, _ int) bool {
			return announcement.PublishedAt.Before(*conds.PublishedAtBefore)
		})
	}

	announcements = base.Map(announcements, func(announcement *announcementdomain.Announcement) *announcementdomain.Announcement {
		return announcement.Clone()
	})

	return announcements, nil
}

func (r *impl) CreateAnnouncements(ctx context.Context, announcements ...*announcementdomain.Announcement) error {
	ids := base.Map(announcements, func(announcement *announcementdomain.Announcement) idtype.AnnouncementID {
		return announcement.ID
	})

	savedIDs := base.Map(r.announcements, func(announcement *announcementdomain.Announcement) idtype.AnnouncementID {
		return announcement.ID
	})

	intersect := lo.Intersect(ids, savedIDs)
	if len(intersect) != 0 {
		return fmt.Errorf("duplicate ids: %v", intersect)
	}

	r.announcements = append(r.announcements, announcements...)

	return nil
}
