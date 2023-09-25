package timetableentity

import "github.com/twin-te/twinte-back/idtype"

type Tag struct {
	ID       idtype.TagID
	UserID   idtype.UserID
	Name     string
	Position int
}
