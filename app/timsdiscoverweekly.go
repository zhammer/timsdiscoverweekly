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

type Archiver interface {
	Archive(ctx context.Context, playlist spotify.FullPlaylist, tracks []spotify.SimpleTrack, parties []ListeningParty) error
}

type TimsDiscoverWeekly struct {
	scraper  Scraper
	spotify  SpotifyGateway
	archiver Archiver

	playlistLength int
}

func NewTimsDiscoverWeekly(spotify SpotifyGateway) *TimsDiscoverWeekly {
	return &TimsDiscoverWeekly{
		scraper:        &ScraperClient{},
		archiver:       &ArchiverClient{},
		spotify:        spotify,
		playlistLength: 30,
	}
}

func (t *TimsDiscoverWeekly) CreatePlaylist(ctx context.Context) error {
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

	var tracks []spotify.SimpleTrack
	for _, album := range albums {
		tracks = append(tracks, randomTrack(*album))
	}
	var trackIDs []spotify.ID
	for _, track := range tracks {
		trackIDs = append(trackIDs, track.ID)
	}

	playlistInput := fromTime(time.Now().UTC())
	playlist, err := t.spotify.CreatePlaylist(ctx, playlistInput)
	if err != nil {
		return err
	}
	if err := t.spotify.AddTracksToPlaylist(ctx, playlist.ID, trackIDs); err != nil {
		return err
	}

	if err := t.archiver.Archive(ctx, *playlist, tracks, listeningParties); err != nil {
		return err
	}

	return nil
}

func shuffleListeningParties(listeningParties []ListeningParty) []ListeningParty {
	// copy array
	var out []ListeningParty
	out = append(out, listeningParties...)

	rand.Shuffle(len(out), func(i, j int) { out[i], out[j] = out[j], out[i] })

	return out
}

func randomTrack(album spotify.FullAlbum) spotify.SimpleTrack {
	return album.Tracks.Tracks[rand.Intn(len(album.Tracks.Tracks))]
}
