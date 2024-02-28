package schoolcalendarrepository

import (
	_ "embed"
	"encoding/json"

	"github.com/twin-te/twinte-back/base"
	schoolcalendardomain "github.com/twin-te/twinte-back/module/schoolcalendar/domain"
)

//go:embed data/event/prod.gen.json
var rawEvents []byte

//go:embed data/module_detail/prod.gen.json
var rawModuleDetails []byte

type jsonEvent struct {
	ID          int     `json:"id"`
	Type        string  `json:"type"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	ChangeTo    *string `json:"changeTo,omitempty"`
}

type jsonModuleDetail struct {
	ID     int    `json:"id"`
	Year   int    `json:"year"`
	Module string `json:"module"`
	Start  string `json:"start"`
	End    string `json:"end"`
}

func loadEvents(data []byte) (events []*schoolcalendardomain.Event, err error) {
	var jsonEvents []*jsonEvent
	if err = json.Unmarshal(data, &jsonEvents); err != nil {
		return
	}

	events, err = base.MapWithErr(jsonEvents, func(jsonEvent *jsonEvent) (*schoolcalendardomain.Event, error) {
		return schoolcalendardomain.ParseEvent(
			jsonEvent.ID,
			jsonEvent.Type,
			jsonEvent.Date,
			jsonEvent.Description,
			jsonEvent.ChangeTo,
		)
	})

	return
}

func loadModuleDetails(data []byte) (moduleDetails []*schoolcalendardomain.ModuleDetail, err error) {
	var jsonModuleDetails []*jsonModuleDetail
	if err = json.Unmarshal(data, &jsonModuleDetails); err != nil {
		return
	}

	moduleDetails, err = base.MapWithErr(jsonModuleDetails, func(jsonModuleDetail *jsonModuleDetail) (*schoolcalendardomain.ModuleDetail, error) {
		return schoolcalendardomain.ParseModuleDetail(
			jsonModuleDetail.ID,
			jsonModuleDetail.Year,
			jsonModuleDetail.Module,
			jsonModuleDetail.Start,
			jsonModuleDetail.End,
		)
	})

	return
}

func LoadEvents() ([]*schoolcalendardomain.Event, error) {
	return loadEvents(rawEvents)
}

func LoadModuleDetails() ([]*schoolcalendardomain.ModuleDetail, error) {
	return loadModuleDetails(rawModuleDetails)
}
