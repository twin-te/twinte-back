package schoolcalendarv1svc

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/twin-te/twinte-back/base"
	schoolcalendarv1conv "github.com/twin-te/twinte-back/handler/api/rpc/schoolcalendar/v1/conv"
	sharedconv "github.com/twin-te/twinte-back/handler/api/rpc/shared/conv"
	schoolcalendarv1 "github.com/twin-te/twinte-back/handler/api/rpcgen/schoolcalendar/v1"
	"github.com/twin-te/twinte-back/handler/api/rpcgen/schoolcalendar/v1/schoolcalendarv1connect"
	schoolcalendarmodule "github.com/twin-te/twinte-back/module/schoolcalendar"
)

var _ schoolcalendarv1connect.SchoolCalendarServiceHandler = (*impl)(nil)

type impl struct {
	uc schoolcalendarmodule.UseCase
}

func (svc *impl) GetEvents(ctx context.Context, req *connect.Request[schoolcalendarv1.GetEventsRequest]) (res *connect.Response[schoolcalendarv1.GetEventsResponse], err error) {
	year, err := sharedconv.FromPBAcadimicYear(req.Msg.Year)
	if err != nil {
		return
	}

	events, err := svc.uc.GetEvents(ctx, year)
	if err != nil {
		return
	}

	pbEvents, err := base.MapWithErr(events, schoolcalendarv1conv.ToPBEvent)
	if err != nil {
		return
	}

	res = connect.NewResponse(&schoolcalendarv1.GetEventsResponse{
		Events: pbEvents,
	})

	return
}

func (svc *impl) GetModuleDetails(ctx context.Context, req *connect.Request[schoolcalendarv1.GetModuleDetailsRequest]) (res *connect.Response[schoolcalendarv1.GetModuleDetailsResponse], err error) {
	year, err := sharedconv.FromPBAcadimicYear(req.Msg.Year)
	if err != nil {
		return
	}

	moduleDetails, err := svc.uc.GetModuleDetails(ctx, year)
	if err != nil {
		return
	}

	pbModuleDetails, err := base.MapWithErr(moduleDetails, schoolcalendarv1conv.ToPBModuleDetail)
	if err != nil {
		return
	}

	res = connect.NewResponse(&schoolcalendarv1.GetModuleDetailsResponse{
		ModuleDetails: pbModuleDetails,
	})

	return
}

func New(uc schoolcalendarmodule.UseCase) *impl {
	return &impl{uc: uc}
}
