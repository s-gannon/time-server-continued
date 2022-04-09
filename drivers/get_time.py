#!/usr/bin/env python3

import time
import json
import serial
import datetime

# takes in data string and outputs either None if there's no signal or
# a 5-tuple containing GMT time, date, latitude, longitude, and unix
# stamp

def parse_gps(data):
    if data[0:6] == "$GPRMC" or data[0:6] == "$GNRMC":  #handles different standards
        ser_data = data.split(",")
        if ser_data[2] == "V":
            print("No satellite data!")
            return None
        else:
            gmt_time = ser_data[1].split(".")[0]    #gets rid of stupid decimal
            date = ser_data[9]
            lat = ser_data[3]
            lat_dir = (1 if ser_data[4] == "N" else -1)
            lon = ser_data[5]
            lon_dir = (1 if ser_data[6] == "E" else -1)

            dt = datetime.datetime.strptime(gmt_time + date, "%H%M%S%d%m%y")
            unix_time = dt.timestamp()
            print(f"{gmt_time} {date} {(float(lat)/100) * lat_dir} {(float(lon)/100) * lon_dir} {unix_time}")
            return (gmt_time, date, (float(lat)/100) * lat_dir, (float(lon)/100) * lon_dir, unix_time)

sleep_time = 1
ser = serial.Serial("/dev/serial0", 9600, timeout = sleep_time)
ser.flushInput()
while True:
    data = ser.readline().decode("ascii")
    result = parse_gps(data)

    if result != None:
        file = open("/home/pi/time-server/data/current_gps_data.json", "w")
        json_str = json.dumps({"lat":f"{result[2]}", "lon":f"{result[3]}", "unix_time":f"{result[4]}"})

        file.write(json_str)
        file.write("\n")
        file.close()
    time.sleep(sleep_time)
ser.close()
