package announcementmodule

import (
	"context"

	announcementdomain "github.com/twin-te/twinte-back/module/announcement/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

// UseCase represents application specific business rules.
//
// The error codes for authentication and authorization failures are not stated explicitly.
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
	//   - shared.NotFound ( if any of the announcements specified by the given ids is not found )
	GetReadFlags(ctx context.Context, ids []idtype.AnnouncementID) (map[idtype.AnnouncementID]bool, error)

	// UpdateReadFlag updates the read flag of the announcement specified by the given id.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - shared.NotFound
	UpdateReadFlag(ctx context.Context, id idtype.AnnouncementID, readFlag bool) error
}
