package schoolcalendarusecase

import (
	"context"

	schoolcalendardomain "github.com/twin-te/twinte-back/module/schoolcalendar/domain"
	schoolcalendarport "github.com/twin-te/twinte-back/module/schoolcalendar/port"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
)

func (uc *impl) GetModuleDetails(ctx context.Context, year shareddomain.AcademicYear) ([]*schoolcalendardomain.ModuleDetail, error) {
	return uc.r.ListModuleDetails(ctx, schoolcalendarport.ListModuleDetailsConds{
		Year: &year,
	}, sharedport.LockNone)
}
