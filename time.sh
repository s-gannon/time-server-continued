#!/bin/sh

parse_unix() {
		date -d "@$1" "+%H:%M:%S"
}

solar_time=$(./solar -t $1)

echo $(parse_unix $1) $(parse_unix $solar_time) $(date -d "@$1" "+%m/%d/%Y") > ~/.cache/campanile_time
