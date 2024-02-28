package timetabledomain

import (
	"fmt"

	"github.com/twin-te/twinte-back/base"
	"golang.org/x/exp/constraints"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=Module -trimprefix=Module -output=module_string.gen.go
type Module int

const (
	ModuleSpringA Module = iota + 1
	ModuleSpringB
	ModuleSpringC
	ModuleFallA
	ModuleFallB
	ModuleFallC
	ModuleSummerVacation
	ModuleSpringVacation
)

var AllModules = []Module{
	ModuleSpringA,
	ModuleSpringB,
	ModuleSpringC,
	ModuleFallA,
	ModuleFallB,
	ModuleFallC,
	ModuleSummerVacation,
	ModuleSpringVacation,
}

func ParseModule(s string) (Module, error) {
	ret, ok := base.FindByString(AllModules, s)
	if ok {
		return ret, nil
	}
	return 0, fmt.Errorf("failed to parse Module %#v", s)
}

//go:generate go run golang.org/x/tools/cmd/stringer -type=Day -trimprefix=Day -output=day_string.gen.go
type Day int

func (d Day) IsNormal() bool {
	return DaySun <= d && d <= DaySat
}

func (d Day) IsSpecial() bool {
	return DayIntensive <= d && d <= DayAnyTime
}

const (
	DaySun Day = iota + 1
	DayMon
	DayTue
	DayWed
	DayThu
	DayFri
	DaySat
	DayIntensive
	DayAppointment
	DayAnyTime
)

var NormalDays = []Day{
	DaySun,
	DayMon,
	DayTue,
	DayWed,
	DayThu,
	DayFri,
	DaySat,
}

var SpecialDays = []Day{
	DayIntensive,
	DayAppointment,
	DayAnyTime,
}

var AllDays = []Day{
	DaySun,
	DayMon,
	DayTue,
	DayWed,
	DayThu,
	DayFri,
	DaySat,
	DayIntensive,
	DayAppointment,
	DayAnyTime,
}

func ParseDay(s string) (Day, error) {
	ret, ok := base.FindByString(AllDays, s)
	if ok {
		return ret, nil
	}
	return 0, fmt.Errorf("failed to parse Day %#v", s)
}

// Period is between 1 and 8.
type Period int

func (p Period) Int() int {
	return int(p)
}

func (p Period) IsZero() bool {
	return p == 0
}

func ParsePeriod[T constraints.Signed](i T) (Period, error) {
	if 1 <= i && i <= 8 {
		return Period(i), nil
	}
	return 0, fmt.Errorf("failed to parse Period %#v", i)
}

// Schedule shows when the course is offered.
//
// There are two types.
//   - normal schedule
//   - special schedule
//
// If this struct represents normal schedule,
// it has module, day, period and rooms fields.
//
// If this struct represents special schedule,
// it has module, day and rooms fields.
type Schedule struct {
	Module Module
	Day    Day
	Period Period
	Rooms  string
}

func (s Schedule) IsNormal() bool {
	return s.Day.IsNormal() && s.Period != 0
}

func (s Schedule) IsSpecial() bool {
	return s.Day.IsSpecial()
}

func ConstructSchedule(fn func() (schedule Schedule, err error)) (schedule Schedule, err error) {
	schedule, err = fn()
	if err != nil {
		return
	}

	if schedule.Day.IsNormal() && !schedule.Period.IsZero() {
		return
	}

	if schedule.Day.IsSpecial() && schedule.Period.IsZero() {
		return
	}

	return schedule, fmt.Errorf("failed to construct %#v", schedule)
}

func ParseSchedule(module string, day string, period int, rooms string) (s Schedule, err error) {
	s.Module, err = ParseModule(module)
	if err != nil {
		return
	}

	s.Day, err = ParseDay(day)
	if err != nil {
		return
	}

	if s.Day.IsNormal() {
		s.Period, err = ParsePeriod(period)
		if err != nil {
			return
		}
	}

	s.Rooms = rooms

	return
}
