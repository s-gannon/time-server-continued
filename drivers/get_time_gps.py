#!/usr/bin/env python3

import time
import json
import serial
from datetime import datetime, tzinfo

# takes in data string and outputs either None if there's no signal or
# a 5-tuple containing GMT time, date, latitude, longitude, and unix
# stamp

def parse_gps(raw_data):
    if raw_data[0:6] == "$GPRMC" or raw_data[0:6] == "$GNRMC":  # handles different standards
        ser_data = raw_data.split(",")

        if ser_data[2] == "V":
            print("No satellite data!")
            return None
        else:
            gmt_time = ser_data[1].split(".")[0]    # gets rid of stupid decimal
            date = ser_data[9]

            lat = ser_data[3]
            lon = ser_data[5]

            # process the direction of the latitude and longitude
            lat_dir = (1 if ser_data[4] == "N" else -1)
            lon_dir = (1 if ser_data[6] == "E" else -1)

            # convert date from the sensor to unix time in GMT
            current_datetime = datetime.strptime(gmt_time + date, "%H%M%S%d%m%y")
			current_datetime = current_datetime.replace(tzinfo=None)
            unix_time = int(current_datetime.timestamp())

            # format lat and long for readable output
            lat_format = (float(lat)/100) * lat_dir
            lon_format = (float(lon)/100) * lon_dir

            print(f"{gmt_time} {date} {lat_format} {lon_format} {unix_time}")
            return int(gmt_time), date, lat_format, lon_format, unix_time
    else:
        # if the data received from the sensor is not of a known format, don't parse

        # don't actually print the below line otherwise it will be printed repeatedly even on working systems
        #print("Satellite data is in an unreadable format.")
        return None

SLEEP_TIME = 1
DEVICE = "/dev/serial0"

ser = serial.Serial(DEVICE, 9600, timeout = SLEEP_TIME)
ser.flushInput()

while True:
    try:
        data = ser.readline().decode("ascii")
        result = parse_gps(data)

        # if there is GPS data, write it to disk in JSON format
        if result is not None:
            with open("/home/pi/time-server/data/current_gps_data.json", "w") as file:
                json_str = json.dumps({"lat":f"{result[2]}",
                                       "lon":f"{result[3]}",
                                       "unix_time":f"{result[4]}"})

                file.write(json_str)
                file.write("\n")
    except serial.SerialTimeoutException:
        print("Serial timeout...")
        continue
    except KeyboardInterrupt:
        print("Keyboard Interrupt!")
        break

    time.sleep(SLEEP_TIME)

ser.close()
