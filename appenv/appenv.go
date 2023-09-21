package appenv

import (
	"os"
	"strconv"
)

// db
var DB_URL string

// cookie
var COOKIE_SECURE bool
var COOKIE_SESSION_NAME string
var COOKIE_OAUTH2_STATE_NAME string
var COOKIE_OAUTH2_STATE_MAX_AGE int

// auth
var AUTH_REDIRECT_URL string

// oauth2 google
var OAUTH2_GOOGLE_CLIENT_ID string
var OAUTH2_GOOGLE_CLIENT_SECRET string
var OAUTH2_GOOGLE_CALLBACK_URL string

// others
var SESSION_LIFE_TIME_DAYS int

func init() {
	// db
	DB_URL = os.Getenv("DB_URL")

	// cookie
	COOKIE_SECURE = must(strconv.ParseBool(os.Getenv("COOKIE_SECURE")))
	COOKIE_SESSION_NAME = os.Getenv("COOKIE_SESSION_NAME")
	COOKIE_OAUTH2_STATE_NAME = os.Getenv("COOKIE_OAUTH2_STATE_NAME")
	COOKIE_OAUTH2_STATE_MAX_AGE = must(strconv.Atoi(os.Getenv("COOKIE_OAUTH2_STATE_MAX_AGE")))

	// auth
	AUTH_REDIRECT_URL = os.Getenv("AUTH_REDIRECT_URL")

	// oauth2 google
	OAUTH2_GOOGLE_CLIENT_ID = os.Getenv("OAUTH2_GOOGLE_CLIENT_ID")
	OAUTH2_GOOGLE_CLIENT_SECRET = os.Getenv("OAUTH2_GOOGLE_CLIENT_SECRET")
	OAUTH2_GOOGLE_CALLBACK_URL = os.Getenv("OAUTH2_GOOGLE_CALLBACK_URL")

	// others
	SESSION_LIFE_TIME_DAYS = must(strconv.Atoi(os.Getenv("SESSION_LIFE_TIME_DAYS")))
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
