package authrepository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	authport "github.com/twin-te/twinte-back/module/auth/port"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Impl struct {
	db boil.ContextExecutor
}

func (r *Impl) Transaction(ctx context.Context, fc func(rtx authport.Repository) error, readOnly bool) error {
	if db, ok := r.db.(boil.ContextBeginner); ok {
		tx, err := db.BeginTx(ctx, &sql.TxOptions{ReadOnly: readOnly})
		if err != nil {
			return err
		}
		if err := fc(&Impl{db: tx}); err != nil {
			return tx.Rollback()
		}
		return tx.Commit()
	}
	return fmt.Errorf("invalid db %+v", r.db)
}

func (r *Impl) innerTransaction(ctx context.Context, fc func(db boil.ContextTransactor) error, readOnly bool) error {
	switch db := r.db.(type) {
	case boil.ContextBeginner:
		tx, err := db.BeginTx(ctx, &sql.TxOptions{ReadOnly: readOnly})
		if err != nil {
			return err
		}
		if err := fc(tx); err != nil {
			return tx.Rollback()
		}
		return tx.Commit()
	case boil.ContextTransactor:
		return fc(db)
	}
	return fmt.Errorf("invalid db %+v", r.db)
}

func New(db boil.ContextExecutor) *Impl {
	return &Impl{db: db}
}
