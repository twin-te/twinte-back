package authentity

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/twin-te/twinte-back/idtype"
)

var SessionLifeTime time.Duration

func init() {
	days, err := strconv.ParseInt(os.Getenv("SESSION_LIFE_TIME_DAYS"), 10, 64)
	if err != nil {
		panic(fmt.Errorf("failed to parse environment variable SESSION_LIFE_TIME_DAYS %s", os.Getenv("SESSION_LIFE_TIME_DAYS")))
	}
	SessionLifeTime = time.Duration(days) * 24 * time.Hour
}

type Session struct {
	ID        idtype.SessionID
	UserID    idtype.UserID
	ExpiredAt time.Time
}

func NewSession(userID idtype.UserID) *Session {
	return &Session{
		ID:        idtype.NewSessionID(),
		UserID:    userID,
		ExpiredAt: time.Now().Add(SessionLifeTime),
	}
}
