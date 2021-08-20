package app

import (
	"context"

	"github.com/zmb3/spotify/v2"
)

type SpotifyClient struct {
	client *spotify.Client
	userID string
}

func (s *SpotifyClient) GetAlbums(ctx context.Context, ids []spotify.ID) ([]*spotify.FullAlbum, error) {
	return s.client.GetAlbums(ctx, ids)
}

func (s *SpotifyClient) CreatePlaylist(ctx context.Context, input CreatePlaylistInput) (*spotify.FullPlaylist, error) {
	return s.client.CreatePlaylistForUser(ctx, s.userID, input.Name, "", true, false)
}

func (s *SpotifyClient) AddTracksToPlaylist(ctx context.Context, playlistID spotify.ID, trackIDs []spotify.ID) error {
	_, err := s.client.AddTracksToPlaylist(ctx, playlistID, trackIDs...)
	return err
}

func NewSpotifyClient(client *spotify.Client, userID string) *SpotifyClient {
	return &SpotifyClient{
		client: client,
	}
}
