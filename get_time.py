import os
import time
import json
import serial
import datetime

# takes in data string and outputs either None if there's no signal or
# a 5-tuple containing GMT time, date, latitude, longitude, and unix
# stamp

def parse_gps(data):
    if data[0:6] == "$GPRMC":
        ser_data = data.split(",")
        if ser_data[2] == "V":
            print("No satellite data!")
            return None
        else:
            gmt_time = ser_data[1].split(".")[0]    #gets rid of stupid decimal
            date = ser_data[9]
            lat = ser_data[3]
            lon = ser_data[5]

            dt = dateime.datetime.strptime(gmt_time + date, "%H%M%S%d%m%y")
            unix_time = dt.timestamp()
            return (gmt_time, date, lat, lon, unix_time)

sleep_time = 10
port = "/dev/serial0"
ser = serial.Serial(port, baudrate = 9600, timeout = 2)

while True:
    data = ser.readline()
    result = parse_gps(data)

    if result != None:
        file = open("current_gps_data.txt", "w")
        json_str = json.dumps({"lat":f"{result[2]}", "lat":f"{result[3]}", "unix_time":f"{result[4]}"})

        file.write(json_str)
        file.close()
    time.sleep(sleep_time)
