package authrepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/twin-te/twinte-back/apperr"
	dbmodel "github.com/twin-te/twinte-back/db/models"
	"github.com/twin-te/twinte-back/idtype"
	authentity "github.com/twin-te/twinte-back/module/auth/entity"
	authport "github.com/twin-te/twinte-back/module/auth/port"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *Impl) SaveSession(ctx context.Context, session *authentity.Session) error {
	dbSession := &dbmodel.Session{
		ID:        session.ID.String(),
		UserID:    session.UserID.String(),
		ExpiredAt: session.ExpiredAt,
	}
	return dbSession.Upsert(ctx, r.db, true, []string{dbmodel.SessionColumns.ID}, boil.Infer(), boil.Infer())
}

func (r *Impl) FindSession(ctx context.Context, conds authport.FindSessionConds) (*authentity.Session, error) {
	mods := make([]qm.QueryMod, 0)

	if conds.ID != nil {
		mods = append(mods, dbmodel.SessionWhere.ID.EQ(conds.ID.String()))
	}

	if conds.UserID != nil {
		mods = append(mods, dbmodel.SessionWhere.UserID.EQ(conds.UserID.String()))
	}

	dbSession, err := dbmodel.Sessions(mods...).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperr.ErrNotFound
		}
		return nil, err
	}

	return fromDBSession(dbSession)
}

func (r *Impl) DeleteSessions(ctx context.Context, conds authport.DeleteSessionsConds) error {
	mods := make([]qm.QueryMod, 0)

	if conds.IDs != nil {
		mods = append(mods, dbmodel.SessionWhere.ID.IN(conds.IDs.StringSlice()))
	}

	if conds.UserIDs != nil {
		mods = append(mods, dbmodel.SessionWhere.UserID.IN(conds.UserIDs.StringSlice()))
	}

	_, err := dbmodel.Sessions(mods...).DeleteAll(ctx, r.db)
	return err
}

func fromDBSession(dbSession *dbmodel.Session) (*authentity.Session, error) {
	id, err := idtype.NewSessionIDFromString(dbSession.ID)
	if err != nil {
		return nil, err
	}

	userID, err := idtype.NewUserIDFromString(dbSession.UserID)
	if err != nil {
		return nil, err
	}

	session := &authentity.Session{
		ID:        id,
		UserID:    userID,
		ExpiredAt: dbSession.ExpiredAt,
	}

	return session, nil
}
