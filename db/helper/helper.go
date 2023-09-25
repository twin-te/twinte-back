package dbhelper

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// CreateValuesQuery creates a sql for values.
// e.g.) CreateValuesQuery(3, 2) returns `($1,$2,$3),($4,$5,$6)`
func CreateValuesQuery(ncolumns, nrows int) (query string) {
	for i := 0; i < nrows; i++ {
		if i > 0 {
			query += ","
		}
		query += "("
		for j := 0; j < ncolumns; j++ {
			if j > 0 {
				query += ","
			}
			query += fmt.Sprintf("$%d", ncolumns*i+j+1)
		}
		query += ")"
	}
	return query
}

func ExecPreparedStmt(ctx context.Context, db boil.ContextExecutor, query string, args []any) error {
	castedDB, ok := db.(interface {
		PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	})
	if !ok {
		return fmt.Errorf("db does not support `PrepareContext`")
	}

	stmt, err := castedDB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}

	return nil
}
