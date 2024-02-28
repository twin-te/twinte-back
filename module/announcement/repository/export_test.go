package announcementrepository

import (
	_ "embed"

	announcementdomain "github.com/twin-te/twinte-back/module/announcement/domain"
)

//go:embed data/announcement_test.json
var rawTestAnnouncements []byte

func LoadTestAnnouncements() ([]*announcementdomain.Announcement, error) {
	return loadAnnouncements(rawTestAnnouncements)
}
