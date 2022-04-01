#!/bin/sh

# Build solar time binary if it is not present.
if [ -f ./solar_conversion/solar ]
then
	# If solar time binary does not exist, golang is required to compile it
	which go || { echo "Please install the go programming language '$ apt install golang' and retry."; exit 1 }

	go build -o ./solar_conversion/solar ./solar_conversion/solar.go
fi

# Install Time Cronjob

# Crontab line to be inserted
line="* * * * * time.sh $(date +%s)"

# create a temporary file to modify the crontab
cron_file=$(mktemp)

# write the current crontab to file
crontab -l > $cron_file

# if the cron_file does not already contain the entry, then add it
grep -v '$line' $cron_file && echo $line >> $cron_file
