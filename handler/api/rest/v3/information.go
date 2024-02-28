package restv3

import (
	"context"

	"github.com/twin-te/twinte-back/appctx"
	"github.com/twin-te/twinte-back/base"
	"github.com/twin-te/twinte-back/handler/api/rest/v3/openapi"
	announcementdomain "github.com/twin-te/twinte-back/module/announcement/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

func toApiInformation(announcement *announcementdomain.Announcement, idToReadFlag map[idtype.AnnouncementID]bool) openapi.Information {
	ret := openapi.Information{
		Content:     announcement.Content.String(),
		Id:          toApiUUID(announcement.ID),
		PublishedAt: announcement.PublishedAt,
		Title:       announcement.Title.String(),
	}

	if readFlag, ok := idToReadFlag[announcement.ID]; ok {
		ret.Read = &readFlag
	}

	return ret
}

// お知らせを取得
// (GET /information)
func (h *impl) GetInformation(ctx context.Context, request openapi.GetInformationRequestObject) (res openapi.GetInformationResponseObject, err error) {
	announcements, err := h.announcementUseCase.GetAnnouncements(ctx)
	if err != nil {
		return
	}

	ids := base.Map(announcements, func(announcement *announcementdomain.Announcement) idtype.AnnouncementID {
		return announcement.ID
	})

	var idToReadFlag map[idtype.AnnouncementID]bool

	actor, ok := appctx.GetActor(ctx)
	if ok && actor.AuthNUser() != nil {
		idToReadFlag, err = h.announcementUseCase.GetReadFlags(ctx, ids)
		if err != nil {
			return nil, err
		}
	}

	apiInformationList := base.MapWithArg(announcements, idToReadFlag, toApiInformation)

	res = openapi.GetInformation200JSONResponse(apiInformationList)

	return
}

// お知らせの既読
// (PUT /information/{id})
func (h *impl) PutInformationId(ctx context.Context, request openapi.PutInformationIdRequestObject) (res openapi.PutInformationIdResponseObject, err error) {
	id, err := idtype.ParseAnnouncementID(request.Id.String())
	if err != nil {
		return
	}

	err = h.announcementUseCase.UpdateReadFlag(ctx, id, request.Body.Read)

	return
}
