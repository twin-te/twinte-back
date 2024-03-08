package restv3

import (
	"context"
	"fmt"
	"sort"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/base"
	"github.com/twin-te/twinte-back/handler/api/rest/v3/openapi"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	timetablemodule "github.com/twin-te/twinte-back/module/timetable"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
)

func toApiTag(tag *timetabledomain.Tag) openapi.Tag {
	return openapi.Tag{
		Id:       openapi_types.UUID(tag.ID),
		Name:     tag.Name.String(),
		Position: lo.ToPtr(int(tag.Position)),
		UserId:   openapi_types.UUID(tag.UserID),
	}
}

// 全てのタグを取得する
// (GET /tags)
func (h *impl) GetTags(ctx context.Context, request openapi.GetTagsRequestObject) (res openapi.GetTagsResponseObject, err error) {
	tags, err := h.timetableUseCase.GetTags(ctx)
	if err != nil {
		return
	}

	apiTags := base.Map(tags, toApiTag)

	res = openapi.GetTags200JSONResponse(apiTags)

	return
}

// タグを並べ替える
// (PATCH /tags)
func (h *impl) PatchTags(ctx context.Context, request openapi.PatchTagsRequestObject) (res openapi.PatchTagsResponseObject, err error) {
	if request.Body == nil {
		return nil, fmt.Errorf("invalid request")
	}

	idToPosition := make(map[idtype.TagID]shareddomain.NonNegativeInt)

	for _, tagPositionOnly := range []openapi.TagPositionOnly(*request.Body) {
		id, err := idtype.ParseTagID(tagPositionOnly.Id.String())
		if err != nil {
			return nil, err
		}

		position, err := timetabledomain.ParsePosition(tagPositionOnly.Position)
		if err != nil {
			return nil, err
		}

		idToPosition[id] = position
	}

	ids := lo.Keys(idToPosition)
	sort.Slice(ids, func(i, j int) bool {
		return idToPosition[ids[i]] < idToPosition[ids[j]]
	})

	err = h.timetableUseCase.RearrangeTags(ctx, ids)
	if err != nil {
		return
	}

	res = openapi.PatchTags200JSONResponse([]openapi.TagPositionOnly(*request.Body))

	return
}

// タグを作成する
// (POST /tags)
func (h *impl) PostTags(ctx context.Context, request openapi.PostTagsRequestObject) (res openapi.PostTagsResponseObject, err error) {
	name, err := timetabledomain.ParseName(request.Body.Name)
	if err != nil {
		return
	}

	tag, err := h.timetableUseCase.CreateTag(ctx, name)
	if err != nil {
		return
	}

	apiTag := toApiTag(tag)

	res = openapi.PostTags200JSONResponse(apiTag)

	return
}

// タグを削除する
// (DELETE /tags/{id})
func (h *impl) DeleteTagsId(ctx context.Context, request openapi.DeleteTagsIdRequestObject) (res openapi.DeleteTagsIdResponseObject, err error) {
	id, err := idtype.ParseTagID(request.Id.String())
	if err != nil {
		return
	}

	err = h.timetableUseCase.DeleteTag(ctx, id)
	if err != nil {
		return
	}

	res = openapi.DeleteTagsId204Response{}

	return
}

// タグを更新する
// (PUT /tags/{id})
func (h *impl) PutTagsId(ctx context.Context, request openapi.PutTagsIdRequestObject) (res openapi.PutTagsIdResponseObject, err error) {
	id, err := idtype.ParseTagID(request.Id.String())
	if err != nil {
		return
	}

	name, err := timetabledomain.ParseName(request.Body.Name)
	if err != nil {
		return
	}

	tag, err := h.timetableUseCase.UpdateTag(ctx, timetablemodule.UpdateTagIn{
		ID:   id,
		Name: &name,
	})
	if err != nil {
		return
	}

	apiTag := toApiTag(tag)

	res = openapi.PutTagsId200JSONResponse(apiTag)

	return
}
