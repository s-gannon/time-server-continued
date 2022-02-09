from gps import *

#returns 2-tuple with latitude and longitude (both floats)
def get_longitude_latitude():
    next_gpsd = gpsd.next()
    if next_gpsd["class"] == "TPV":
        lat_lon = (getattr(next_gpsd, "lat", "Unknown"),
                   getattr(next_gpsd, "lon", "Unknown"))
        return lat_lon

#takes in 2-tuple with latitude and longitude and returns the time at that location
def get_time(lat_lon):
