#!/bin/sh

sshfs pi@campanile-node-1:/home/pi/time-server/data /home/pi/time-data/node1
sshfs pi@campanile-node-2:/home/pi/time-server/data /home/pi/time-data/node2
sshfs pi@campanile-node-3:/home/pi/time-server/data /home/pi/time-data/node3
