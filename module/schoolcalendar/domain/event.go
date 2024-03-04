package schoolcalendardomain

import (
	"errors"
	"fmt"

	"cloud.google.com/go/civil"
	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/base"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=Weekday -trimprefix=Weekday -output=weekday_string.gen.go
type Weekday int

const (
	WeekdaySunday Weekday = iota + 1
	WeekdayMonday
	WeekdayTuesday
	WeekdayWednesday
	WeekdayThursday
	WeekdayFriday
	WeekdaySaturday
)

var AllWeekdays = []Weekday{
	WeekdaySunday,
	WeekdayMonday,
	WeekdayTuesday,
	WeekdayWednesday,
	WeekdayThursday,
	WeekdayFriday,
	WeekdaySaturday,
}

func ParseWeekday(s string) (Weekday, error) {
	ret, ok := base.FindByString(AllWeekdays, s)
	if ok {
		return ret, nil
	}
	return 0, fmt.Errorf("failed to parse Weekday %#v", s)
}

//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType -trimprefix=EventType -output=event_type_string.gen.go
type EventType int

func (et EventType) IsZero() bool {
	return et == 0
}

func (et EventType) IsSubstituteDay() bool {
	return et == EventTypeSubstituteDay
}

const (
	EventTypeHoliday EventType = iota + 1
	EventTypePublicHoliday
	EventTypeExam
	EventTypeSubstituteDay
	EventTypeOther
)

var AllEventTypes = []EventType{
	EventTypeHoliday,
	EventTypePublicHoliday,
	EventTypeExam,
	EventTypeSubstituteDay,
	EventTypeOther,
}

func ParseEventType(s string) (EventType, error) {
	ret, ok := base.FindByString(AllEventTypes, s)
	if ok {
		return ret, nil
	}
	return 0, fmt.Errorf("failed to parse EventType %#v", s)
}

type Event struct {
	ID          idtype.EventID
	Type        EventType
	Date        civil.Date
	Description string

	// It is not nil, only if Type is EventTypeSubstituteDay.
	ChangeTo *Weekday

	EntityBeforeUpdated *Event
}

func (e *Event) Clone() *Event {
	ret := lo.ToPtr(*e)

	if e.ChangeTo != nil {
		*ret.ChangeTo = *e.ChangeTo
	}

	return ret
}

func ConstructEvent(fn func(e *Event) (err error)) (*Event, error) {
	e := new(Event)
	if err := fn(e); err != nil {
		return nil, err
	}

	if e.Type.IsSubstituteDay() && e.ChangeTo == nil {
		return nil, errors.New("field 'ChangeTo' must not be nil for substitute event")
	}

	if e.ID.IsZero() || e.Type.IsZero() || e.Date.IsZero() {
		return nil, fmt.Errorf("failed to construct %+v", e)
	}

	return e, nil
}

func ParseEvent(id int, eventType string, date string, description string, changeTo *string) (event *Event, err error) {
	return ConstructEvent(func(e *Event) (err error) {
		e.ID, err = idtype.ParseEventID(id)
		if err != nil {
			return
		}

		e.Type, err = ParseEventType(eventType)
		if err != nil {
			return
		}

		e.Date, err = civil.ParseDate(date)
		if err != nil {
			return
		}

		if e.Type.IsSubstituteDay() {
			if changeTo == nil {
				return errors.New("field 'ChangeTo' must not be nil for substitute event")
			}

			e.ChangeTo, err = base.ToPtrWithErr(ParseWeekday(*changeTo))
			if err != nil {
				return
			}
		}

		return
	})
}
