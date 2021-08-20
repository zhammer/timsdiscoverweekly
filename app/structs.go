package app

import (
	"fmt"
	"time"
)

type CreatePlaylistInput struct {
	Name string `json:"name"`
}

func fromTime(time time.Time) CreatePlaylistInput {
	date := time.Format("2006-01-02")
	name := fmt.Sprintf("[%s] Tim's Discover Weekly", date)
	return CreatePlaylistInput{name}
}

// listening parties scraped from the timstwitterlisteningparty repo
type ListeningParty struct {
	AlbumID string
}
