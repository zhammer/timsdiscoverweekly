package main

import (
	"context"
	"os"
	"timsdiscoverweekly/app"
	"timsdiscoverweekly/pkg/auth_http_client"

	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"
	"github.com/zmb3/spotify/v2"
)

type Config struct {
	SpotifyUserID      string `required:"true" split_words:"true"`
	SpotifyBearerToken string `required:"true" split_words:"true"`
}

func main() {
	logger := zerolog.New(zerolog.NewConsoleWriter())
	ctx := context.Background()
	ctx = logger.WithContext(ctx)

	c := cli.App{
		Name: "timsdiscoverweekly",
		Commands: []*cli.Command{
			{
				Name: "generate-playlist",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "spotify-user-id",
						EnvVars:  []string{"SPOTIFY_USER_ID"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "spotify-bearer-token",
						EnvVars:  []string{"SPOTIFY_BEARER_TOKEN"},
						Required: true,
					},
				},
				Action: generatePlaylist,
			},
		},
	}

	if err := c.RunContext(ctx, os.Args); err != nil {
		logger.Fatal().Msg(err.Error())
	}
}

func generatePlaylist(c *cli.Context) error {
	spotify := app.NewSpotifyClient(
		spotify.New(auth_http_client.New(c.String("spotify-bearer-token"))),
		c.String("spotify-user-id"),
	)
	scraper := &app.ScraperClient{}
	app := app.NewTimsDiscoverWeekly(scraper, spotify)

	return app.CreatePlaylist(c.Context)
}
