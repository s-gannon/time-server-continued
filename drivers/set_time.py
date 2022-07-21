#!/usr/bin/env python3

import datetime
import json
import os

while True:
	node_fails = [False, False, False]
	num_fails = 0
	try:
		fnode1 = open("/home/pi/time-data/node1/current_gps_data.json", "r")
	        data1 = json.load(fnode1)
	except:
		node_fails[0] = True
		num_fails += 1
	try:
		fnode2 = open("/home/pi/time-data/node2/current_gps_data.json", "r")
		data2 = json.load(fnode2)
	except:
                node_fails[1] = True
                num_fails += 1
	try:
		fnode3 = open("/home/pi/time-data/node3/current_gps_data.json", "r")
		data3 = json.load(fnode3)
	except:
                node_fails[2] = True
                num_fails += 1
	if num_fails == 3:
		continue
	else:
		sum = 0
		for node_num in range(3):
			if not node_fails[node_num]:
				if node_num == 1:
					sum += int(data1["unix_time"])
				elif node_num == 2:
					sum += int(data2["unix_time"])
				elif node_num == 3:
					sum += int(data3["unix_time"])
		try:
			sum /= (3 - num_fails)
		except:
			print("ERR: Divide by Zero occured")
			return
		date = datetime.utcfromtimestamp(int(sum)).strftime("%Y-%m-%d %H:%M:%S")
		os.system(f"sudo date -s {date}")
