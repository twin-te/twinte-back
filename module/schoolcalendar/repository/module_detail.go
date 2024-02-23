package schoolcalendarrepository

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/base"
	schoolcalendardomain "github.com/twin-te/twinte-back/module/schoolcalendar/domain"
	schoolcalendarport "github.com/twin-te/twinte-back/module/schoolcalendar/port"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
)

func (r *impl) ListModuleDetails(ctx context.Context, conds schoolcalendarport.ListModuleDetailsConds, lock sharedport.Lock) ([]*schoolcalendardomain.ModuleDetail, error) {
	moduleDetails := r.moduleDetails

	if conds.Year != nil {
		moduleDetails = lo.Filter(moduleDetails, func(moduleDetail *schoolcalendardomain.ModuleDetail, _ int) bool {
			return moduleDetail.Year == *conds.Year
		})
	}

	moduleDetails = base.Map(moduleDetails, func(moduleDetail *schoolcalendardomain.ModuleDetail) *schoolcalendardomain.ModuleDetail {
		return moduleDetail.Clone()
	})

	return moduleDetails, nil
}

func (r *impl) CreateModuleDetails(ctx context.Context, moduleDetails ...*schoolcalendardomain.ModuleDetail) error {
	ids := base.Map(moduleDetails, func(event *schoolcalendardomain.ModuleDetail) idtype.ModuleDetailID {
		return event.ID
	})

	savedIDs := base.Map(r.moduleDetails, func(event *schoolcalendardomain.ModuleDetail) idtype.ModuleDetailID {
		return event.ID
	})

	intersect := lo.Intersect(ids, savedIDs)
	if len(intersect) != 0 {
		return fmt.Errorf("duplicate ids: %v", intersect)
	}

	r.moduleDetails = append(r.moduleDetails, moduleDetails...)

	return nil
}