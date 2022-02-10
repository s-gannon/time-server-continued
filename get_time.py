from gps import *
import datetime
import time
import os

gpsd = gps(mode=WATCH_ENABLE) #global gpsd instance

#returns 3-tuple with latitude, longitude (both floats), and time
def get_latitude_longitude_time():
    global gpsd
    next_gpsd = gpsd.next()
    return (gpsd.fix.latitude, gpsd.fix.longitude, gpsd.utc)

if __name__ == "__main__":
    try:
        while True:
            lat_lon_time = get_latitude_longitude_time()
            dt = datetime.datetime.strptime(f"{lat_lon_time[2]}","%Y%m%dT%H%M")
            unix_time = dt.timestamp()
            file = open("current_gps.txt", "w")

            file.write(f"{lat_lon_time[0]}\n{lat_lon_time[1]}\n{unix_time}")
            file.close()
            time.sleep(5)
    except KeyboardInterrupt:
        print("Node interrupted by KeyboardInterrupt!")
