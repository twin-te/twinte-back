package authrepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/apperr"
	dbhelper "github.com/twin-te/twinte-back/db/helper"
	dbmodel "github.com/twin-te/twinte-back/db/models"
	"github.com/twin-te/twinte-back/idtype"
	authentity "github.com/twin-te/twinte-back/module/auth/entity"
	authport "github.com/twin-te/twinte-back/module/auth/port"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *Impl) SaveUser(ctx context.Context, user *authentity.User) error {
	return r.innerTransaction(ctx, func(tx boil.ContextTransactor) error {
		dbUser := &dbmodel.User{
			ID: user.ID.String(),
		}

		if err := dbUser.Upsert(ctx, tx, true, []string{dbmodel.UserColumns.ID}, boil.Infer(), boil.Infer()); err != nil {
			return err
		}

		return r.saveUserAuthentications(ctx, tx, user)
	}, false)
}

func (r *Impl) FindUser(ctx context.Context, conds authport.FindUserConds) (*authentity.User, error) {
	mods := make([]qm.QueryMod, 0)
	authenticationMods := make([]qm.QueryMod, 0)

	if conds.ID != nil {
		mods = append(mods, dbmodel.UserWhere.ID.EQ(conds.ID.String()))
	}

	if conds.UserAuthentication != nil {
		dbProvider, err := toDBProvider(conds.UserAuthentication.Provider)
		if err != nil {
			return nil, err
		}

		authenticationMods = append(
			authenticationMods,
			dbmodel.UserAuthenticationWhere.Provider.EQ(dbProvider),
			dbmodel.UserAuthenticationWhere.SocialID.EQ(conds.UserAuthentication.SocialID.String()),
		)
	}

	mods = append(
		mods,
		dbmodel.UserWhere.DeletedAt.EQ(null.NewTime(time.Time{}, false)),
		qm.Load(dbmodel.UserRels.UserAuthentications, authenticationMods...),
	)

	dbUser, err := dbmodel.Users(mods...).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperr.ErrNotFound
		}
		return nil, err
	}

	return fromDBUser(dbUser)
}

func (r *Impl) DeleteUser(ctx context.Context, id idtype.UserID) error {
	return r.innerTransaction(ctx, func(db boil.ContextTransactor) error {
		if _, err := dbmodel.UserAuthentications(
			dbmodel.UserAuthenticationWhere.UserID.EQ(id.String()),
		).DeleteAll(ctx, db); err != nil {
			return err
		}

		dbUser := &dbmodel.User{ID: id.String(), DeletedAt: null.TimeFrom(time.Now())}
		_, err := dbUser.Update(ctx, db, boil.Whitelist(dbmodel.UserColumns.DeletedAt))
		return err
	}, false)
}

func (r *Impl) saveUserAuthentications(ctx context.Context, db boil.ContextExecutor, user *authentity.User) error {
	dbUserAuthentications, err := toDBUserAuthentications(user.ID, user.Authentications)
	if err != nil {
		return err
	}

	dbProviders := lo.Map(dbUserAuthentications, func(dbUserAuthentication *dbmodel.UserAuthentication, index int) string {
		return dbUserAuthentication.Provider
	})

	_, err = dbmodel.UserAuthentications(
		dbmodel.UserAuthenticationWhere.UserID.EQ(user.ID.String()),
		dbmodel.UserAuthenticationWhere.Provider.NIN(dbProviders),
	).DeleteAll(ctx, db)
	if err != nil {
		return err
	}

	args := make([]any, 0, 3*len(dbUserAuthentications))
	for _, dbUserAuthentication := range dbUserAuthentications {
		args = append(
			args,
			dbUserAuthentication.ID,
			dbUserAuthentication.Provider,
			dbUserAuthentication.SocialID,
		)
	}

	query := fmt.Sprintf(`insert into user_authentications (user_id,provider,social_id) values %s
on conflict (user_id,provider)
do update set (social_id) = (excluded.social_id)`, dbhelper.CreateValuesQuery(3, len(dbUserAuthentications)))

	return dbhelper.ExecPreparedStmt(ctx, db, query, args)
}

func fromDBProvider(dbProvider string) (authentity.Provider, error) {
	switch dbProvider {
	case dbmodel.UserAuthenticationsProviderEnumGoogle:
		return authentity.ProviderGoogle, nil
	case dbmodel.UserAuthenticationsProviderEnumApple:
		return authentity.ProviderApple, nil
	case dbmodel.UserAuthenticationsProviderEnumTwitter:
		return authentity.ProviderTwitter, nil
	}
	return 0, fmt.Errorf("invalid provider %s", dbProvider)
}

func toDBProvider(provider authentity.Provider) (string, error) {
	switch provider {
	case authentity.ProviderGoogle:
		return dbmodel.UserAuthenticationsProviderEnumGoogle, nil
	case authentity.ProviderApple:
		return dbmodel.UserAuthenticationsProviderEnumApple, nil
	case authentity.ProviderTwitter:
		return dbmodel.UserAuthenticationsProviderEnumTwitter, nil
	}
	return "", fmt.Errorf("invalid provider %d", provider)
}

func fromDBUserAuthentication(dbUserAuthentication *dbmodel.UserAuthentication) (authentity.UserAuthentication, error) {
	provider, err := fromDBProvider(dbUserAuthentication.Provider)
	if err != nil {
		return authentity.UserAuthentication{}, err
	}

	userAuthentication := authentity.UserAuthentication{
		Provider: provider,
		SocialID: authentity.NewSocialIDFromString(dbUserAuthentication.SocialID),
	}

	return userAuthentication, nil
}

func fromDBUserAuthentications(dbUserAuthentications dbmodel.UserAuthenticationSlice) ([]authentity.UserAuthentication, error) {
	userAuthentications := make([]authentity.UserAuthentication, 0, len(dbUserAuthentications))

	for _, dbUserAuthentication := range dbUserAuthentications {
		userAuthentication, err := fromDBUserAuthentication(dbUserAuthentication)
		if err != nil {
			return nil, err
		}
		userAuthentications = append(userAuthentications, userAuthentication)
	}

	return userAuthentications, nil
}

func toDBUserAuthentication(userID idtype.UserID, userAuthentication authentity.UserAuthentication) (*dbmodel.UserAuthentication, error) {
	dbProvider, err := toDBProvider(userAuthentication.Provider)
	if err != nil {
		return nil, err
	}

	dbUserAuthentication := &dbmodel.UserAuthentication{
		UserID:   userID.String(),
		Provider: dbProvider,
		SocialID: userAuthentication.SocialID.String(),
	}

	return dbUserAuthentication, nil
}

func toDBUserAuthentications(userID idtype.UserID, userAuthentications []authentity.UserAuthentication) (dbmodel.UserAuthenticationSlice, error) {
	dbUserAuthentications := make(dbmodel.UserAuthenticationSlice, 0, len(userAuthentications))
	for _, userAuthentication := range userAuthentications {
		dbUserAuthentication, err := toDBUserAuthentication(userID, userAuthentication)
		if err != nil {
			return nil, err
		}
		dbUserAuthentications = append(dbUserAuthentications, dbUserAuthentication)
	}
	return dbUserAuthentications, nil
}

func fromDBUser(dbUser *dbmodel.User) (*authentity.User, error) {
	id, err := idtype.NewUserIDFromString(dbUser.ID)
	if err != nil {
		return nil, err
	}

	userAuthentications, err := fromDBUserAuthentications(dbUser.R.GetUserAuthentications())
	if err != nil {
		return nil, err
	}

	user := &authentity.User{
		ID:              id,
		Authentications: userAuthentications,
	}

	return user, nil
}
