package main

import (
	"fmt"
	"time"
	"math"
)

// make sure all Sin/Cos are in Radians or converted
// 180/pi = radians to degrees, pi/180 = degrees to radians

func main() {
	long := 87.037403 // TODO: don't hardcode this value

	fmt.Println(SolarTime(long))
}

// calculate local solar time
// TODO: need to properly calculate the difference in time. Also make sure daylight savings is considered
func SolarTime(long float64) string {
	return time.Now().Add(time.ParseDuration(timeCorrectionFactor(long))).Format(time.ANSIC)
}

// Calculate the difference in time between standard time and solar time.
func timeCorrectionFactor(long float64) float64 {
	return 4 * (float64(localStandardTimeMeridian(time.Now())) - long) + equationOfTime(time.Now().YearDay())
}

// Get standard meridian for time zone in degrees.
// 75 Eastern ; 90 Central ; 105 Mountain ; 120 Pacific
func localStandardTimeMeridian(t time) int {
	// UTC difference; Eg: UTC-6 == 6 & UTC+7 == 7
	_, zone := t.Zone()
	// Zone() returns how far away from UTC a person is in seconds. Divide by 3600 to get hours.
	return 15 * intAbs(zone/3600)
}

// TODO: Need to verify and find out if the formula uses radians or degrees
func equationOfTime(days int) float64 {
	// days: specifically, days since the start of the year. Eg. Feb 2 should be 32 (0 is Jan 1)
	b := 0.9863 * float64(days - 81)
	return 229.2 * (0.000075 + 0.001868 * math.Cos(b) - 0.032077 * math.Sin(b) - 0.014615 * math.Cos(2 * b) - 0.04089 * math.Sin(2 * b))
}

// Calculate the absolute value of an integer. This prevents the need to convert to a float64 and back to use the built in abs function.
func intAbs(x int) int {
	if x < 0 {
		x = -x
	}

	return x
}
