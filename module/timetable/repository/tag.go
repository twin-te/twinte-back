package timetablerepository

import (
	"context"
	"fmt"

	dbhelper "github.com/twin-te/twinte-back/db/helper"
	dbmodel "github.com/twin-te/twinte-back/db/models"
	"github.com/twin-te/twinte-back/idtype"
	timetableentity "github.com/twin-te/twinte-back/module/timetable/entity"
	timetableport "github.com/twin-te/twinte-back/module/timetable/port"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (r *Impl) SaveTag(ctx context.Context, tag *timetableentity.Tag) error {
	dbTag := toDBTag(tag)
	return dbTag.Upsert(ctx, r.db, true, []string{dbmodel.TagColumns.ID}, boil.Infer(), boil.Infer())
}

func (r *Impl) UpdateTag(ctx context.Context, tags []*timetableentity.Tag) error {
	return nil
}

func (r *Impl) SaveTags(ctx context.Context, tags []*timetableentity.Tag) error {
	dbTags := toDBTags(tags)

	args := make([]any, 0, 4*len(dbTags))
	for _, dbTag := range dbTags {
		args = append(
			args,
			dbTag.ID,
			dbTag.UserID,
			dbTag.Name,
			dbTag.Position,
		)
	}

	query := fmt.Sprintf(`insert into tags (id,user_id,name,position) values %s
on conflict (id,user_id)
do update set (name,position) = (excluded.name,excluded.position)`, dbhelper.CreateValuesQuery(4, len(dbTags)))

	return dbhelper.ExecPreparedStmt(ctx, r.db, query, args)
}

func (r *Impl) FindTag(ctx context.Context, conds timetableport.FindTagConds) (*timetableentity.Tag, error) {
	return nil, nil
}

func (r *Impl) ListTags(ctx context.Context, conds timetableport.ListTagsConds) ([]*timetableentity.Tag, error) {
	return nil, nil
}

func (r *Impl) DeleteTag(ctx context.Context, id idtype.TagID) error {
	return nil
}

func toDBTag(tag *timetableentity.Tag) *dbmodel.Tag {
	return &dbmodel.Tag{
		ID:       tag.ID.String(),
		UserID:   tag.UserID.String(),
		Name:     tag.Name,
		Position: tag.Position,
	}
}

func toDBTags(tags []*timetableentity.Tag) dbmodel.TagSlice {
	dbTags := make(dbmodel.TagSlice, 0, len(tags))
	for _, tag := range tags {
		dbTags = append(dbTags, toDBTag(tag))
	}
	return dbTags
}
