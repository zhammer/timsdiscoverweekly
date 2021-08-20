package app

import (
	"context"
	"math/rand"
	"time"

	"github.com/zmb3/spotify/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Scraper interface {
	GetListeningParties(context.Context) ([]ListeningParty, error)
}

type SpotifyGateway interface {
	GetAlbums(ctx context.Context, ids []spotify.ID) ([]*spotify.FullAlbum, error)
	CreatePlaylist(context.Context, CreatePlaylistInput) (*spotify.FullPlaylist, error)
	AddTracksToPlaylist(ctx context.Context, playlistID spotify.ID, trackIDs []spotify.ID) error
}

type TimsDiscoverWeekly struct {
	scraper Scraper
	spotify SpotifyGateway

	playlistLength int
}

func NewTimsDiscoverWeekly(scraper Scraper, spotify SpotifyGateway) *TimsDiscoverWeekly {
	return &TimsDiscoverWeekly{
		scraper:        scraper,
		spotify:        spotify,
		playlistLength: 30,
	}
}

// Scrape the timstwitterlisteningparty listening party csv,
// for any new albums, get the # of tracks, then save to our
// local collection.
func (t *TimsDiscoverWeekly) ScrapeAlbums(ctx context.Context) error {
	listeningParties, err := t.scraper.GetListeningParties(ctx)
	if err != nil {
		return err
	}

	selectedParties := shuffleListeningParties(listeningParties)[:t.playlistLength]

	var albumIDs []spotify.ID
	for _, lp := range selectedParties {
		albumIDs = append(albumIDs, spotify.ID(lp.AlbumID))
	}

	albums, err := t.spotify.GetAlbums(ctx, albumIDs)
	if err != nil {
		return err
	}

	var trackIDs []spotify.ID
	for _, album := range albums {
		track := randomTrack(*album)
		trackIDs = append(trackIDs, track.ID)
	}

	playlist, err := t.spotify.CreatePlaylist(ctx, fromTime(time.Now().UTC()))
	if err != nil {
		return err
	}
	if err := t.spotify.AddTracksToPlaylist(ctx, playlist.ID, trackIDs); err != nil {
		return err
	}

	return nil
}

func shuffleListeningParties(listeningParties []ListeningParty) []ListeningParty {
	// copy array
	var out []ListeningParty
	for _, lp := range listeningParties {
		out = append(out, lp)
	}

	rand.Shuffle(len(out), func(i, j int) { out[i], out[j] = out[j], out[i] })

	return out
}

func randomTrack(album spotify.FullAlbum) spotify.SimpleTrack {
	return album.Tracks.Tracks[rand.Intn(len(album.Tracks.Tracks))]
}
