package announcementusecase

import (
	"context"
	"time"

	"github.com/samber/lo"
	announcementdomain "github.com/twin-te/twinte-back/module/announcement/domain"
	announcementport "github.com/twin-te/twinte-back/module/announcement/port"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
)

func (uc *impl) GetAnnouncements(ctx context.Context) ([]*announcementdomain.Announcement, error) {
	return uc.r.ListAnnouncements(ctx, announcementport.ListAnnouncementsConds{
		PublishedAtBefore: lo.ToPtr(time.Now()),
	}, sharedport.LockNone)
}
