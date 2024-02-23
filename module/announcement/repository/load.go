package announcementrepository

import (
	"encoding/json"
	"time"

	_ "embed"

	"github.com/twin-te/twinte-back/base"
	announcementdomain "github.com/twin-te/twinte-back/module/announcement/domain"
)

//go:embed data/announcement.json
var rawAnnouncements []byte

type jsonAnnouncement struct {
	ID          string    `json:"id"`
	Tags        []string  `json:"tags"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"publishedAt"`
}

func loadAnnouncements(data []byte) (announcements []*announcementdomain.Announcement, err error) {
	var jsonAnnouncements []*jsonAnnouncement
	if err = json.Unmarshal(data, &jsonAnnouncements); err != nil {
		return
	}

	announcements, err = base.MapWithErr(jsonAnnouncements, func(jsonAnnouncement *jsonAnnouncement) (*announcementdomain.Announcement, error) {
		return announcementdomain.ParseAnnouncement(
			jsonAnnouncement.ID,
			jsonAnnouncement.Tags,
			jsonAnnouncement.Title,
			jsonAnnouncement.Content,
			jsonAnnouncement.PublishedAt,
		)
	})

	return
}

func LoadAnnouncements() ([]*announcementdomain.Announcement, error) {
	return loadAnnouncements(rawAnnouncements)
}
