package schoolcalendarrepository

import (
	_ "embed"

	schoolcalendardomain "github.com/twin-te/twinte-back/module/schoolcalendar/domain"
)

//go:embed data/event/test.json
var rawTestEvents []byte

//go:embed data/module_detail/test.json
var rawTestModuleDetails []byte

func LoadTestEvents() ([]*schoolcalendardomain.Event, error) {
	return loadEvents(rawTestEvents)
}

func LoadTestModuleDetails() ([]*schoolcalendardomain.ModuleDetail, error) {
	return loadModuleDetails(rawTestModuleDetails)
}
