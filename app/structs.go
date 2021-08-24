package app

import (
	"fmt"
	"time"
)

type CreatePlaylistInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func fromTime(time time.Time) CreatePlaylistInput {
	date := time.Format("2006-01-02")
	name := fmt.Sprintf("[%s] Tim's Discover Weekly", date)
	description := fmt.Sprintf("https://github.com/zhammer/timsdiscoverweekly/blob/main/playlists/%s.md", name)
	return CreatePlaylistInput{name, description}
}

// listening parties scraped from the timstwitterlisteningparty repo
type ListeningParty struct {
	AlbumID   string
	ReplayURL string
}
