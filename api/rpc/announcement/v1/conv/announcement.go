package announcementv1conv

import (
	sharedconv "github.com/twin-te/twinte-back/api/rpc/shared/conv"
	announcementv1 "github.com/twin-te/twinte-back/api/rpcgen/announcement/v1"
	"github.com/twin-te/twinte-back/base"
	announcementdomain "github.com/twin-te/twinte-back/module/announcement/domain"
)

func ToPBAnnouncement(announcement *announcementdomain.Announcement) (*announcementv1.Announcement, error) {
	pbAnnouncementTag, err := base.MapWithErr(announcement.Tags, ToPBAnnouncementTag)
	if err != nil {
		return nil, err
	}

	pbAnnouncement := &announcementv1.Announcement{
		Id:          sharedconv.ToPBUUID(announcement.ID),
		Tags:        pbAnnouncementTag,
		Title:       announcement.Title.String(),
		Content:     announcement.Content.String(),
		PublishedAt: sharedconv.ToPBRFC3339DateTime(announcement.PublishedAt),
	}

	return pbAnnouncement, nil
}
