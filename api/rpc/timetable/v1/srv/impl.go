package timetablev1svc

import (
	"context"

	"github.com/bufbuild/connect-go"
	sharedconv "github.com/twin-te/twinte-back/api/rpc/shared/conv"
	timetablev1conv "github.com/twin-te/twinte-back/api/rpc/timetable/v1/conv"
	timetablev1 "github.com/twin-te/twinte-back/api/rpcgen/timetable/v1"
	"github.com/twin-te/twinte-back/api/rpcgen/timetable/v1/timetablev1connect"
	"github.com/twin-te/twinte-back/base"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	timetablemodule "github.com/twin-te/twinte-back/module/timetable"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
)

var _ timetablev1connect.TimetableServiceHandler = (*impl)(nil)

type impl struct {
	uc timetablemodule.UseCase
}

func (svc *impl) GetCourses(ctx context.Context, req *connect.Request[timetablev1.GetCoursesRequest]) (res *connect.Response[timetablev1.GetCoursesResponse], err error) {
	year, err := sharedconv.FromPBAcadimicYear(req.Msg.Year)
	if err != nil {
		return
	}

	codes, err := base.MapWithErr(req.Msg.Codes, timetabledomain.ParseCode)
	if err != nil {
		return
	}

	courses, err := svc.uc.GetCoursesByCodes(ctx, year, codes)
	if err != nil {
		return
	}

	pbCourses, err := base.MapWithErr(courses, timetablev1conv.ToPBCourse)
	if err != nil {
		return
	}

	res = connect.NewResponse(&timetablev1.GetCoursesResponse{
		Courses: pbCourses,
	})

	return
}

func (svc *impl) CreateRegisteredCoursesByCodes(ctx context.Context, req *connect.Request[timetablev1.CreateRegisteredCoursesByCodesRequest]) (res *connect.Response[timetablev1.CreateRegisteredCoursesByCodesResponse], err error) {
	year, err := sharedconv.FromPBAcadimicYear(req.Msg.Year)
	if err != nil {
		return
	}

	codes, err := base.MapWithErr(req.Msg.Codes, timetabledomain.ParseCode)
	if err != nil {
		return
	}

	registeredCourses, err := svc.uc.CreateRegisteredCoursesByCodes(ctx, year, codes)
	if err != nil {
		return
	}

	pbRegisteredCourses, err := base.MapWithErr(registeredCourses, timetablev1conv.ToPBRegisteredCourse)
	if err != nil {
		return
	}

	res = connect.NewResponse(&timetablev1.CreateRegisteredCoursesByCodesResponse{
		RegisteredCourses: pbRegisteredCourses,
	})

	return
}

func (svc *impl) CreateRegisteredCourseManually(ctx context.Context, req *connect.Request[timetablev1.CreateRegisteredCourseManuallyRequest]) (res *connect.Response[timetablev1.CreateRegisteredCourseManuallyResponse], err error) {
	year, err := sharedconv.FromPBAcadimicYear(req.Msg.Year)
	if err != nil {
		return
	}

	name, err := timetabledomain.ParseName(req.Msg.Name)
	if err != nil {
		return
	}

	credit, err := timetabledomain.ParseCredit(req.Msg.Credit)
	if err != nil {
		return
	}

	methods, err := base.MapWithErr(req.Msg.Methods, timetablev1conv.FromPBCourseMethod)
	if err != nil {
		return
	}

	schedules, err := base.MapWithErr(req.Msg.Schedules, timetablev1conv.FromPBSchedule)
	if err != nil {
		return
	}

	in := timetablemodule.CreateRegisteredCourseManuallyIn{
		Year:        year,
		Name:        name,
		Instructors: req.Msg.Instructors,
		Credit:      credit,
		Methods:     methods,
		Schedules:   schedules,
	}

	registeredCourse, err := svc.uc.CreateRegisteredCourseManually(ctx, in)
	if err != nil {
		return
	}

	pbRegisteredCourse, err := timetablev1conv.ToPBRegisteredCourse(registeredCourse)
	if err != nil {
		return
	}

	res = connect.NewResponse(&timetablev1.CreateRegisteredCourseManuallyResponse{
		RegisteredCourse: pbRegisteredCourse,
	})

	return
}

func (svc *impl) GetRegisteredCourses(ctx context.Context, req *connect.Request[timetablev1.GetRegisteredCoursesRequest]) (res *connect.Response[timetablev1.GetRegisteredCoursesResponse], err error) {
	var year *shareddomain.AcademicYear

	if req.Msg.Year != nil {
		year, err = base.ToPtrWithErr(sharedconv.FromPBAcadimicYear(req.Msg.Year))
		if err != nil {
			return
		}
	}

	registeredCourses, err := svc.uc.GetRegisteredCourses(ctx, year)
	if err != nil {
		return
	}

	pbRegisteredCourses, err := base.MapWithErr(registeredCourses, timetablev1conv.ToPBRegisteredCourse)
	if err != nil {
		return
	}

	res = connect.NewResponse(&timetablev1.GetRegisteredCoursesResponse{
		RegisteredCourses: pbRegisteredCourses,
	})

	return
}

func (svc *impl) UpdateRegisteredCourse(ctx context.Context, req *connect.Request[timetablev1.UpdateRegisteredCourseRequest]) (res *connect.Response[timetablev1.UpdateRegisteredCourseResponse], err error) {
	in := timetablemodule.UpdateRegisteredCourseIn{
		Instructors: req.Msg.Instructors,
		Memo:        req.Msg.Memo,
	}

	in.ID, err = sharedconv.FromPBUUID(req.Msg.Id, idtype.ParseRegisteredCourseID)
	if err != nil {
		return
	}

	if req.Msg.Name != nil {
		in.Name, err = base.ToPtrWithErr(timetabledomain.ParseName(*req.Msg.Name))
		if err != nil {
			return
		}
	}

	if req.Msg.Credit != nil {
		in.Credit, err = base.ToPtrWithErr(timetabledomain.ParseCredit(*req.Msg.Credit))
		if err != nil {
			return
		}
	}

	if req.Msg.Methods != nil {
		in.Methods, err = base.ToPtrWithErr(base.MapWithErr(req.Msg.Methods.Values, timetablev1conv.FromPBCourseMethod))
		if err != nil {
			return
		}
	}

	if req.Msg.Schedules != nil {
		in.Schedules, err = base.ToPtrWithErr(base.MapWithErr(req.Msg.Schedules.Values, timetablev1conv.FromPBSchedule))
		if err != nil {
			return
		}
	}

	if req.Msg.Attendance != nil {
		in.Attendance, err = base.ToPtrWithErr(timetabledomain.ParseAttendance(int(*req.Msg.Attendance)))
		if err != nil {
			return
		}
	}

	if req.Msg.Late != nil {
		in.Late, err = base.ToPtrWithErr(timetabledomain.ParseLate(int(*req.Msg.Late)))
		if err != nil {
			return
		}
	}

	if req.Msg.Absence != nil {
		in.Absence, err = base.ToPtrWithErr(timetabledomain.ParseAbsence(int(*req.Msg.Absence)))
		if err != nil {
			return
		}
	}

	if req.Msg.TagIds != nil {
		in.TagIDs, err = base.ToPtrWithErr(base.MapWithArgAndErr(req.Msg.TagIds.Values, idtype.ParseTagID, sharedconv.FromPBUUID[idtype.TagID]))
		if err != nil {
			return
		}
	}

	registeredCourse, err := svc.uc.UpdateRegisteredCourse(ctx, in)
	if err != nil {
		return
	}

	pbRegisteredCourse, err := timetablev1conv.ToPBRegisteredCourse(registeredCourse)
	if err != nil {
		return
	}

	res = connect.NewResponse(&timetablev1.UpdateRegisteredCourseResponse{
		RegisteredCourse: pbRegisteredCourse,
	})

	return
}

func (svc *impl) DeleteRegisteredCourse(ctx context.Context, req *connect.Request[timetablev1.DeleteRegisteredCourseRequest]) (res *connect.Response[timetablev1.DeleteRegisteredCourseResponse], err error) {
	id, err := sharedconv.FromPBUUID(req.Msg.Id, idtype.ParseRegisteredCourseID)
	if err != nil {
		return
	}

	if err = svc.uc.DeleteRegisteredCourse(ctx, id); err != nil {
		return
	}

	res = connect.NewResponse(&timetablev1.DeleteRegisteredCourseResponse{})

	return
}

func (svc *impl) CreateTag(ctx context.Context, req *connect.Request[timetablev1.CreateTagRequest]) (res *connect.Response[timetablev1.CreateTagResponse], err error) {
	name, err := timetabledomain.ParseName(req.Msg.Name)
	if err != nil {
		return
	}

	tag, err := svc.uc.CreateTag(ctx, name)
	if err != nil {
		return
	}

	pbTag := timetablev1conv.ToPBTag(tag)

	res = connect.NewResponse(&timetablev1.CreateTagResponse{
		Tag: pbTag,
	})

	return
}

func (svc *impl) GetTags(ctx context.Context, req *connect.Request[timetablev1.GetTagsRequest]) (res *connect.Response[timetablev1.GetTagsResponse], err error) {
	tags, err := svc.uc.GetTags(ctx)
	if err != nil {
		return
	}

	pbTags := base.Map(tags, timetablev1conv.ToPBTag)

	res = connect.NewResponse(&timetablev1.GetTagsResponse{
		Tags: pbTags,
	})

	return
}

func (svc *impl) UpdateTag(ctx context.Context, req *connect.Request[timetablev1.UpdateTagRequest]) (res *connect.Response[timetablev1.UpdateTagResponse], err error) {
	in := timetablemodule.UpdateTagIn{}

	in.ID, err = sharedconv.FromPBUUID(req.Msg.Id, idtype.ParseTagID)
	if err != nil {
		return
	}

	if req.Msg.Name != nil {
		in.Name, err = base.ToPtrWithErr(timetabledomain.ParseName(*req.Msg.Name))
		if err != nil {
			return
		}
	}

	tag, err := svc.uc.UpdateTag(ctx, in)
	if err != nil {
		return
	}

	pbTag := timetablev1conv.ToPBTag(tag)

	res = connect.NewResponse(&timetablev1.UpdateTagResponse{
		Tag: pbTag,
	})

	return
}

func (svc *impl) DeleteTag(ctx context.Context, req *connect.Request[timetablev1.DeleteTagRequest]) (res *connect.Response[timetablev1.DeleteTagResponse], err error) {
	id, err := sharedconv.FromPBUUID(req.Msg.Id, idtype.ParseTagID)
	if err != nil {
		return
	}

	if err = svc.uc.DeleteTag(ctx, id); err != nil {
		return
	}

	res = connect.NewResponse(&timetablev1.DeleteTagResponse{})

	return
}

func (svc *impl) RearrangeTags(ctx context.Context, req *connect.Request[timetablev1.RearrangeTagsRequest]) (res *connect.Response[timetablev1.RearrangeTagsResponse], err error) {
	ids, err := base.MapWithArgAndErr(req.Msg.TagIds, idtype.ParseTagID, sharedconv.FromPBUUID[idtype.TagID])
	if err != nil {
		return
	}

	if err = svc.uc.RearrangeTags(ctx, ids); err != nil {
		return
	}

	res = connect.NewResponse(&timetablev1.RearrangeTagsResponse{})

	return
}

func New(uc timetablemodule.UseCase) *impl {
	return &impl{uc: uc}
}
