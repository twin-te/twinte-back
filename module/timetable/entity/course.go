package timetableentity

import (
	"time"

	"github.com/twin-te/twinte-back/idtype"
)

type Code string

func NewCodeFromString(s string) Code {
	return Code(s)
}

type Course struct {
	ID                idtype.CourseID
	Year              int
	Code              Code
	Name              string
	Instructors       string
	Cregit            float64
	Overview          string
	Remarks           string
	LastUpdatedAt     time.Time
	HasParseError     bool
	IsAnnual          bool
	RecommendedGrades []int
	Methods           []CourseMethod
	Schedules         []Schedule
}
