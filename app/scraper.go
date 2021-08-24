package app

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/rs/zerolog"
)

type ScraperClient struct {
}

var spotifyLinkRegex = regexp.MustCompile(`https://open\.spotify\.com/album/(.+)`)

func (s *ScraperClient) GetListeningParties(ctx context.Context) ([]ListeningParty, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/matbroughty/timstwitterlisteningparty/master/data/time-slot-data.csv")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received unexpected status code (%d) when downloading listening parties file", resp.StatusCode)
	}

	return pluckListeningParties(ctx, resp.Body)
}

func pluckListeningParties(ctx context.Context, body io.Reader) ([]ListeningParty, error) {
	records, err := csv.NewReader(body).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error parsing listening party csv: %w", err)
	}

	var listeningParties []ListeningParty
	for _, record := range records {
		spotifyLink := record[7]
		groups := spotifyLinkRegex.FindStringSubmatch(spotifyLink)
		if len(groups) != 2 {
			zerolog.Ctx(ctx).Debug().Msgf("Encountered spotify album link in unexpected format: '%s'", spotifyLink)
			continue
		}
		listeningParties = append(listeningParties, ListeningParty{
			AlbumID:   groups[1],
			ReplayURL: record[4],
		})
	}
	return listeningParties, nil
}
