#!/bin/sh

cd clock; go build -o webclock server.go; cd ..

loc="$(echo "$(cd "$(dirname "$1")"; pwd -P)/$(basename "$1")")"

ln -s $loc/clock/webclock /usr/bin/
ln -s $loc/clock/webclock@.service /etc/systemd/system/

systemctl enable --now webclock
