#!/bin/sh

# Helper file to view time outputs. This is not used to format time for the web clock.

parse_unix() {
		date -d "@$1" "+%H:%M:%S"
}

solar_time_unix=$(../solarcli/main -t $1)
solar_time_formatted=$(parse_unix $solar_time_unix)
clock_time_formatted=$(parse_unix $1)
todays_date=$(date -d "@$1" "+%m/%d/%Y")

echo $clock_time_formatted $solar_time_formatted $todays_date
