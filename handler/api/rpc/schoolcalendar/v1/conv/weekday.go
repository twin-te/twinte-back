package schoolcalendarv1conv

import (
	"fmt"

	schoolcalendarv1 "github.com/twin-te/twinte-back/handler/api/rpcgen/schoolcalendar/v1"
	schoolcalendardomain "github.com/twin-te/twinte-back/module/schoolcalendar/domain"
)

func FromPBWeekday(pbWeekday schoolcalendarv1.Weekday) (schoolcalendardomain.Weekday, error) {
	switch pbWeekday {
	case schoolcalendarv1.Weekday_WEEKDAY_SUNDAY:
		return schoolcalendardomain.WeekdaySunday, nil
	case schoolcalendarv1.Weekday_WEEKDAY_MONDAY:
		return schoolcalendardomain.WeekdayMonday, nil
	case schoolcalendarv1.Weekday_WEEKDAY_TUESDAY:
		return schoolcalendardomain.WeekdayTuesday, nil
	case schoolcalendarv1.Weekday_WEEKDAY_WEDNESDAY:
		return schoolcalendardomain.WeekdayWednesday, nil
	case schoolcalendarv1.Weekday_WEEKDAY_THURSDAY:
		return schoolcalendardomain.WeekdayThursday, nil
	case schoolcalendarv1.Weekday_WEEKDAY_FRIDAY:
		return schoolcalendardomain.WeekdayFriday, nil
	case schoolcalendarv1.Weekday_WEEKDAY_SATURDAY:
		return schoolcalendardomain.WeekdaySaturday, nil
	}
	return 0, fmt.Errorf("invalid %#v", pbWeekday)
}

func ToPBWeekday(weekday schoolcalendardomain.Weekday) (schoolcalendarv1.Weekday, error) {
	switch weekday {
	case schoolcalendardomain.WeekdaySunday:
		return schoolcalendarv1.Weekday_WEEKDAY_SUNDAY, nil
	case schoolcalendardomain.WeekdayMonday:
		return schoolcalendarv1.Weekday_WEEKDAY_MONDAY, nil
	case schoolcalendardomain.WeekdayTuesday:
		return schoolcalendarv1.Weekday_WEEKDAY_TUESDAY, nil
	case schoolcalendardomain.WeekdayWednesday:
		return schoolcalendarv1.Weekday_WEEKDAY_WEDNESDAY, nil
	case schoolcalendardomain.WeekdayThursday:
		return schoolcalendarv1.Weekday_WEEKDAY_THURSDAY, nil
	case schoolcalendardomain.WeekdayFriday:
		return schoolcalendarv1.Weekday_WEEKDAY_FRIDAY, nil
	case schoolcalendardomain.WeekdaySaturday:
		return schoolcalendarv1.Weekday_WEEKDAY_SATURDAY, nil
	}
	return 0, fmt.Errorf("invalid %#v", weekday)
}
