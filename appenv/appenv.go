package appenv

import (
	"os"
	"strconv"
)

// db
var DB_URL string
var TEST_DB_URL string

// cookie
var COOKIE_SECURE bool
var COOKIE_SESSION_NAME string
var COOKIE_OAUTH2_STATE_NAME string
var COOKIE_OAUTH2_STATE_MAX_AGE int

// auth
var AUTH_REDIRECT_URL string
var SESSION_LIFE_TIME_DAYS int

// oauth2
var OAUTH2_GOOGLE_CLIENT_ID string
var OAUTH2_GOOGLE_CLIENT_SECRET string
var OAUTH2_GOOGLE_CALLBACK_URL string
var OAUTH2_TWITTER_CLIENT_ID string
var OAUTH2_TWITTER_CLIENT_SECRET string
var OAUTH2_TWITTER_CALLBACK_URL string

// timetable
var KDB_JSON_FILE_PATH string

func init() {
	// db
	DB_URL = os.Getenv("DB_URL")
	TEST_DB_URL = os.Getenv("TEST_DB_URL")

	// cookie
	COOKIE_SECURE = must(strconv.ParseBool(os.Getenv("COOKIE_SECURE")))
	COOKIE_SESSION_NAME = os.Getenv("COOKIE_SESSION_NAME")
	COOKIE_OAUTH2_STATE_NAME = os.Getenv("COOKIE_OAUTH2_STATE_NAME")
	COOKIE_OAUTH2_STATE_MAX_AGE = must(strconv.Atoi(os.Getenv("COOKIE_OAUTH2_STATE_MAX_AGE")))

	// auth
	AUTH_REDIRECT_URL = os.Getenv("AUTH_REDIRECT_URL")
	SESSION_LIFE_TIME_DAYS = must(strconv.Atoi(os.Getenv("SESSION_LIFE_TIME_DAYS")))

	// oauth2
	OAUTH2_GOOGLE_CLIENT_ID = os.Getenv("OAUTH2_GOOGLE_CLIENT_ID")
	OAUTH2_GOOGLE_CLIENT_SECRET = os.Getenv("OAUTH2_GOOGLE_CLIENT_SECRET")
	OAUTH2_GOOGLE_CALLBACK_URL = os.Getenv("OAUTH2_GOOGLE_CALLBACK_URL")
	OAUTH2_TWITTER_CLIENT_ID = os.Getenv("OAUTH2_TWITTER_CLIENT_ID")
	OAUTH2_TWITTER_CLIENT_SECRET = os.Getenv("OAUTH2_TWITTER_CLIENT_SECRET")
	OAUTH2_TWITTER_CALLBACK_URL = os.Getenv("OAUTH2_TWITTER_CALLBACK_URL")

	// timetable
	KDB_JSON_FILE_PATH = os.Getenv("KDB_JSON_FILE_PATH")
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
