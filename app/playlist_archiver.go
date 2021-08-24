package app

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/zmb3/spotify/v2"
)

type ArchiverClient struct{}

func (a *ArchiverClient) Archive(ctx context.Context, playlist spotify.FullPlaylist, tracks []spotify.SimpleTrack, parties []ListeningParty) error {
	file, err := os.Create(path.Join("playlists", playlist.Name+".md"))
	if err != nil {
		return err
	}

	markdown := renderMarkdown(playlist, tracks, parties)
	_, err = file.WriteString(markdown)
	return err
}

func renderMarkdown(playlist spotify.FullPlaylist, tracks []spotify.SimpleTrack, parties []ListeningParty) string {
	out := ""
	out += fmt.Sprintf("# [%s](https://open.spotify.com/user/%s/playlist/%s)\n", playlist.Name, playlist.Owner.ID, playlist.ID)
	out += "\n"
	out += "| Title | Artist | Listening Party |\n"
	out += "| --- | --- | --- |\n"
	for i, track := range tracks {
		party := parties[i]
		out += fmt.Sprintf("| %s | %s | [link](%s) |\n", track.Name, renderArtists(track.Artists), party.ReplayURL)
	}

	return out
}

func renderArtists(artists []spotify.SimpleArtist) string {
	var artistNames []string
	for _, artist := range artists {
		artistNames = append(artistNames, artist.Name)
	}
	return strings.Join(artistNames, ", ")
}
