package timetableentity

import "fmt"

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

type NormalDay int

const (
	NormalDaySun NormalDay = iota + 1
	NormalDayMon
	NormalDayTue
	NormalDayWed
	NormalDayThu
	NormalDayFri
	NormalDaySat
)

type SpecialDay int

const (
	SpecialDayIntensive SpecialDay = iota + 1
	SpecialDayAppointment
	SpecialDayAnyTime
)

type Period int

func (p Period) Int() int {
	return int(p)
}

func NewPeriodFromInt(i int) (Period, error) {
	if 1 <= i && i <= 8 {
		return Period(i), nil
	}
	return 0, fmt.Errorf("invalid period %d", i)
}

type NormalSchedule struct {
	Module Module
	Day    NormalDay
	Period Period
	Rooms  string
}

func (ns NormalSchedule) Normal() (NormalSchedule, bool) {
	return ns, true
}

func (ns NormalSchedule) Special() (SpecialSchedule, bool) {
	return SpecialSchedule{}, false
}

type SpecialSchedule struct {
	Module Module
	Day    SpecialDay
	Rooms  string
}

func (ss SpecialSchedule) Normal() (NormalSchedule, bool) {
	return NormalSchedule{}, false
}

func (ss SpecialSchedule) Special() (SpecialSchedule, bool) {
	return ss, true
}

type Schedule interface {
	Normal() (NormalSchedule, bool)
	Special() (SpecialSchedule, bool)
}
