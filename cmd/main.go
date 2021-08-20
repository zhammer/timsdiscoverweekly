package main

import (
	"context"
	"timsdiscoverweekly/app"
	"timsdiscoverweekly/pkg/auth_http_client"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
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

	if err := run(ctx); err != nil {
		logger.Fatal().Msg(err.Error())
	}
}

func run(ctx context.Context) error {
	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		return err
	}

	spotify := app.NewSpotifyClient(
		spotify.New(auth_http_client.New(cfg.SpotifyBearerToken)),
		cfg.SpotifyUserID,
	)
	scraper := &app.ScraperClient{}
	app := app.NewTimsDiscoverWeekly(scraper, spotify)

	return app.CreatePlaylist(ctx)
}
