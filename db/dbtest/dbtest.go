package dbtest

import (
	"github.com/twin-te/twinte-back/appenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewTestDB() (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  appenv.TEST_DB_URL,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
		AllowGlobalUpdate:        true,
	})
}
