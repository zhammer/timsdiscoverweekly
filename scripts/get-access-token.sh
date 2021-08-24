#!/bin/bash

set -e

curl --silent -u $SPOTIFY_CLIENT_ID:$SPOTIFY_CLIENT_SECRET -d grant_type=refresh_token -d refresh_token=$SPOTIFY_REFRESH_TOKEN https://accounts.spotify.com/api/token \
    | python -c 'import json,sys;obj=json.load(sys.stdin); print(obj["access_token"])'
