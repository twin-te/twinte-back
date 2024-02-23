package shareddomain

import (
	"fmt"
)

// AcademicYear starts in April and ends in March.
// Zero value is invalid.
type AcademicYear int

func (year AcademicYear) Int() int {
	return int(year)
}

func (year AcademicYear) IsZero() bool {
	return year.Int() == 0
}

func ParseAcademicYear(i int) (AcademicYear, error) {
	if i == 0 {
		return 0, fmt.Errorf("failed to parse AcademicYear %#v", i)
	}
	return AcademicYear(i), nil
}
