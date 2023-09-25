package timetableentity

import (
	"time"

	"github.com/twin-te/twinte-back/idtype"
)

type EventType int

const (
	EventTypeHoliday EventType = iota + 1
	EventTypePublicHoliday
	EventTypeExam
	EventTypeSubstituteDay
	EventTypeOther
)

type Event struct {
	ID          idtype.EventID
	Type        EventType
	Date        time.Time
	Description string
	ChangeTo    *NormalDay
}

type ModuleDetail struct {
	ID     idtype.ModuleDetailID
	Year   int
	Module Module
	Start  time.Time
	End    time.Time
}
