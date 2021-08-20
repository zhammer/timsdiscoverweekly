package app

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/http"
	"regexp"

	"github.com/rs/zerolog"
)

type ScraperClient struct {
}

var spotifyLinkRegex = regexp.MustCompile(`"https://open\.spotify\.com/album/(.+)"`)

func GetListeningParties(ctx context.Context) ([]ListeningParty, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/matbroughty/timstwitterlisteningparty/master/data/time-slot-data.csv")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Received unexpected status code (%d) when downloading listening parties file", resp.StatusCode)
	}

	records, err := csv.NewReader(resp.Body).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Error parsing listening party csv: %w", err)
	}

	var listeningParties []ListeningParty
	// example record: "2021-08-08T21:00","Life Without Buildings","Any Other City","https://twitter.com/Tim_Burgess/status/1417529737472364551","https://timstwitterlisteningparty.com/pages/replay/feed_873.html","@SueTompkins2018:@JohnKennedy:@rdg_music:@Tim_Burgess","https://twitter.com/LlSTENlNG_PARTY/timelines/1426531470521380866","https://open.spotify.com/album/1c7eigkoEcDAKKhkajY3Br","https://i.scdn.co/image/ab67616d00001e02f557471eedfbd9a7ab55d75b","https://i.scdn.co/image/ab67616d00004851f557471eedfbd9a7ab55d75b","2000","917","https://i.scdn.co/image/ab67616d0000b273f557471eedfbd9a7ab55d75b"
	for _, record := range records {
		spotifyLink := record[7]
		groups := spotifyLinkRegex.FindStringSubmatch(spotifyLink)
		if len(groups) != 2 {
			zerolog.Ctx(ctx).Warn().Msgf("Encountered spotify album link in unexpected format: '%s'", spotifyLink)
			continue
		}
		listeningParties = append(listeningParties, ListeningParty{AlbumID: groups[2]})
	}

	return listeningParties, nil
}
