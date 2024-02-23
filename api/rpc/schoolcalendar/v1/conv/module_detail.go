package schoolcalendarv1conv

import (
	sharedconv "github.com/twin-te/twinte-back/api/rpc/shared/conv"
	schoolcalendarv1 "github.com/twin-te/twinte-back/api/rpcgen/schoolcalendar/v1"
	schoolcalendardomain "github.com/twin-te/twinte-back/module/schoolcalendar/domain"
)

func ToPBModuleDetail(moduleDetail *schoolcalendardomain.ModuleDetail) (*schoolcalendarv1.ModuleDetail, error) {
	pbModule, err := ToPBModule(moduleDetail.Module)
	if err != nil {
		return nil, err
	}

	pbModuleDetail := &schoolcalendarv1.ModuleDetail{
		Id:     int32(moduleDetail.ID.Int()),
		Year:   sharedconv.ToPBAcademicYear(moduleDetail.Year),
		Module: pbModule,
		Start:  sharedconv.ToPBRFC3339FullDate(moduleDetail.Start),
		End:    sharedconv.ToPBRFC3339FullDate(moduleDetail.End),
	}

	return pbModuleDetail, nil
}
