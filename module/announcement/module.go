package announcementmodule

import (
	"context"

	announcementdomain "github.com/twin-te/twinte-back/module/announcement/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

// UseCase represents application specific business rules.
//
// The following error codes are not stated explicitly in the each method, but may be returned.
//   - shared.InvalidArgument
//   - shared.Unauthenticated
//   - shared.Unauthorized
type UseCase interface {
	// GetAnnouncements returns all published announcements.
	//
	// [Authentication] not required
	GetAnnouncements(ctx context.Context) ([]*announcementdomain.Announcement, error)

	// GetReadFlags returns the read flags of the announcements specified by the given ids.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - announcement.AnnouncementNotFound
	GetReadFlags(ctx context.Context, ids []idtype.AnnouncementID) (map[idtype.AnnouncementID]bool, error)

	// UpdateReadFlag updates the read flag of the announcement specified by the given id.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - announcement.AnnouncementNotFound
	UpdateReadFlag(ctx context.Context, id idtype.AnnouncementID, readFlag bool) error
}
