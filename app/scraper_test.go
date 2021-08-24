package app

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPluckListeningParties(t *testing.T) {
	csv := `"2021-08-08T21:00","Life Without Buildings","Any Other City","https://twitter.com/Tim_Burgess/status/1417529737472364551","https://timstwitterlisteningparty.com/pages/replay/feed_873.html","@SueTompkins2018:@JohnKennedy:@rdg_music:@Tim_Burgess","https://twitter.com/LlSTENlNG_PARTY/timelines/1426531470521380866","https://open.spotify.com/album/1c7eigkoEcDAKKhkajY3Br","https://i.scdn.co/image/ab67616d00001e02f557471eedfbd9a7ab55d75b","https://i.scdn.co/image/ab67616d00004851f557471eedfbd9a7ab55d75b","2000","917","https://i.scdn.co/image/ab67616d0000b273f557471eedfbd9a7ab55d75b"
"2021-12-20T20:00","No Listening Party","The Charlatans are playing at Aberdeen","https://www.thecharlatans.net/gigs","","","","https://www.thecharlatans.net/gigs","https://pbs.twimg.com/media/E4rVC5jXIAQgb-i?format=jpg&name=900x900","https://pbs.twimg.com/media/E4rVC5jXIAQgb-i?format=jpg&name=900x900","2021","-1","https://pbs.twimg.com/media/E4rVC5jXIAQgb-i?format=jpg&name=900x900"`
	body := strings.NewReader(csv)

	listeningParties, err := pluckListeningParties(context.Background(), body)
	assert.NoError(t, err)
	assert.ElementsMatch(t, listeningParties, []ListeningParty{
		{
			AlbumID:   "1c7eigkoEcDAKKhkajY3Br",
			ReplayURL: "https://timstwitterlisteningparty.com/pages/replay/feed_873.html",
		},
	})
}
