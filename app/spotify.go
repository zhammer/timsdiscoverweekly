package app

import (
	"context"
	"fmt"

	"github.com/zmb3/spotify/v2"
)

type SpotifyClient struct {
	client *spotify.Client
}

func (s *SpotifyClient) GetAlbums(ctx context.Context, ids []spotify.ID) ([]*spotify.FullAlbum, error) {
	return s.client.GetAlbums(ctx, ids)
}

func (s *SpotifyClient) CreatePlaylist(ctx context.Context, input CreatePlaylistInput) (*spotify.FullPlaylist, error) {
	account, err := s.client.CurrentUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting current user: %w", err)
	}
	return s.client.CreatePlaylistForUser(ctx, account.ID, input.Name, "", true, false)
}

func (s *SpotifyClient) AddTracksToPlaylist(ctx context.Context, playlistID spotify.ID, trackIDs []spotify.ID) error {
	_, err := s.client.AddTracksToPlaylist(ctx, playlistID, trackIDs...)
	return err
}

func NewSpotifyClient(client *spotify.Client) *SpotifyClient {
	return &SpotifyClient{
		client: client,
	}
}
