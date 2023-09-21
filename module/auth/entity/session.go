package authentity

import (
	"time"

	"github.com/twin-te/twinte-back/appenv"
	"github.com/twin-te/twinte-back/idtype"
)

var SessionLifeTime time.Duration

func init() {
	SessionLifeTime = time.Duration(appenv.SESSION_LIFE_TIME_DAYS) * 24 * time.Hour
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
