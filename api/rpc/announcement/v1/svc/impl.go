package announcementv1svc

import (
	"context"

	"github.com/bufbuild/connect-go"

	announcementv1conv "github.com/twin-te/twinte-back/api/rpc/announcement/v1/conv"
	sharedconv "github.com/twin-te/twinte-back/api/rpc/shared/conv"
	announcementv1 "github.com/twin-te/twinte-back/api/rpcgen/announcement/v1"
	"github.com/twin-te/twinte-back/api/rpcgen/announcement/v1/announcementv1connect"
	"github.com/twin-te/twinte-back/base"
	announcementmodule "github.com/twin-te/twinte-back/module/announcement"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

var _ announcementv1connect.AnnouncementServiceHandler = (*impl)(nil)

type impl struct {
	uc announcementmodule.UseCase
}

func (svc *impl) GetAnnouncements(ctx context.Context, req *connect.Request[announcementv1.GetAnnouncementsRequest]) (res *connect.Response[announcementv1.GetAnnouncementsResponse], err error) {
	announcements, err := svc.uc.GetAnnouncements(ctx)
	if err != nil {
		return
	}

	pbAnnouncements, err := base.MapWithErr(announcements, announcementv1conv.ToPBAnnouncement)
	if err != nil {
		return
	}

	res = connect.NewResponse(&announcementv1.GetAnnouncementsResponse{
		Announcements: pbAnnouncements,
	})

	return
}

func (svc *impl) GetReadFlags(ctx context.Context, req *connect.Request[announcementv1.GetReadFlagsRequest]) (res *connect.Response[announcementv1.GetReadFlagsResponse], err error) {
	ids, err := base.MapWithArgAndErr(req.Msg.Ids, idtype.ParseAnnouncementID, sharedconv.FromPBUUID[idtype.AnnouncementID])
	if err != nil {
		return
	}

	idToReadFlag, err := svc.uc.GetReadFlags(ctx, ids)
	if err != nil {
		return
	}

	readFlags := base.Map(ids, func(id idtype.AnnouncementID) bool {
		return idToReadFlag[id]
	})

	res = connect.NewResponse(&announcementv1.GetReadFlagsResponse{
		ReadFlags: readFlags,
	})

	return
}

func (svc *impl) UpdateReadFlag(ctx context.Context, req *connect.Request[announcementv1.UpdateReadFlagRequest]) (res *connect.Response[announcementv1.UpdateReadFlagResponse], err error) {
	id, err := sharedconv.FromPBUUID(req.Msg.Id, idtype.ParseAnnouncementID)
	if err != nil {
		return
	}

	if err = svc.uc.UpdateReadFlag(ctx, id, req.Msg.ReadFlag); err != nil {
		return
	}

	res = connect.NewResponse(&announcementv1.UpdateReadFlagResponse{})

	return
}

func New(uc announcementmodule.UseCase) *impl {
	return &impl{uc: uc}
}
