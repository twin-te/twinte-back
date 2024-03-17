package announcementusecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/apperr"
	announcementdomain "github.com/twin-te/twinte-back/module/announcement/domain"
	announcementerr "github.com/twin-te/twinte-back/module/announcement/err"
	announcementport "github.com/twin-te/twinte-back/module/announcement/port"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
)

func (uc *impl) GetReadFlags(ctx context.Context, ids []idtype.AnnouncementID) (map[idtype.AnnouncementID]bool, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	announcements, err := uc.r.ListAnnouncements(ctx, announcementport.ListAnnouncementsConds{
		PublishedAtBefore: lo.ToPtr(time.Now()),
	}, sharedport.LockNone)
	if err != nil {
		return nil, err
	}

	for _, id := range ids {
		if !lo.ContainsBy(announcements, func(announcement *announcementdomain.Announcement) bool {
			return id == announcement.ID
		}) {
			return nil, apperr.New(announcementerr.CodeAnnouncementNotFound, fmt.Sprintf("not found announcement whose id is %s", id))
		}
	}

	alreadyReads, err := uc.r.ListAlreadyReads(ctx, announcementport.ListAlreadyReadsConds{
		UserID:          &userID,
		AnnouncementIDs: &ids,
	}, sharedport.LockNone)
	if err != nil {
		return nil, err
	}

	idToReadFlag := lo.SliceToMap(ids, func(id idtype.AnnouncementID) (idtype.AnnouncementID, bool) {
		return id, false
	})

	for _, alreadyRead := range alreadyReads {
		idToReadFlag[alreadyRead.AnnouncementID] = true
	}

	return idToReadFlag, nil
}

func (uc *impl) UpdateReadFlag(ctx context.Context, id idtype.AnnouncementID, readFlag bool) error {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	_, err = uc.r.FindAnnouncement(ctx, announcementport.FindAnnouncementConds{
		ID: id, PublishedAtBefore: lo.ToPtr(time.Now()),
	}, sharedport.LockNone)
	if err != nil {
		if errors.Is(err, sharedport.ErrNotFound) {
			return apperr.New(announcementerr.CodeAnnouncementNotFound, fmt.Sprintf("not found announcement whose id is %s", id))
		}
		return err
	}

	if readFlag {
		return uc.r.Transaction(ctx, func(rtx announcementport.Repository) error {
			_, err := uc.r.FindAlreadyRead(ctx, announcementport.FindAlreadyReadConds{
				UserID:         userID,
				AnnouncementID: id,
			}, sharedport.LockNone)
			if err != nil && !errors.Is(err, sharedport.ErrNotFound) {
				return err
			}
			if err == nil {
				return nil
			}

			alreadyRead, err := uc.f.NewAlreadyRead(userID, id)
			if err != nil {
				return err
			}

			return uc.r.CreateAlreadyReads(ctx, alreadyRead)
		})
	}

	_, err = uc.r.DeleteAlreadyReads(ctx, announcementport.DeleteAlreadyReadsConds{
		UserID:         &userID,
		AnnouncementID: &id,
	})
	return err
}
