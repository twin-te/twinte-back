package appenv

import (
	"os"
	"strconv"
)

var DB_URL string
var SESSION_LIFE_TIME_DAYS int
var COOKIE_SECURE bool
var COOKIE_SESSION_NAME string
var COOKIE_OAUTH2_STATE_NAME string
var OAUTH2_REDIRECT_URL string

func init() {
	DB_URL = os.Getenv("DB_URL")
	SESSION_LIFE_TIME_DAYS = must(strconv.Atoi(os.Getenv("DB_URL")))
	COOKIE_SECURE = must(strconv.ParseBool(os.Getenv("COOKIE_SECURE")))
	COOKIE_SESSION_NAME = os.Getenv("COOKIE_SESSION_NAME")
	COOKIE_OAUTH2_STATE_NAME = os.Getenv("COOKIE_OAUTH2_STATE_NAME")
	OAUTH2_REDIRECT_URL = os.Getenv("OAUTH2_REDIRECT_URL")
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
