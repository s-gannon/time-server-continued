package main

import (
	"fmt"
	"time"
	"math"
)

// following the functions detailed in this website: https://www.pveducation.org/pvcdrom/properties-of-sunlight/solar-time

// make sure all Sin/Cos are in Radians or converted

func main() {
	current := time.Now().Local()
	long := -87.037403 // TODO: don't hardcode this value

	fmt.Println(localSolarTime(long, current))
}

// calculate local solar time
func localSolarTime(long float64, time time.Time) float64 {
	return time.Now()/* current? */ + (timeCorrectionFactor(long) / 60)
}

func localStandardTimeMeridian() float64 {
	// zone: UTC +- what? Eg: UTC-6 == 6
	_, zone := time.Now().Zone()
	return 15 * math.Abs(zone)
}

// 180/pi = radians to degrees, pi/180 = degrees to radians

func equationOfTime() float64 {
	// days: specifically, days since the start of the year. Eg. Feb 2 should be 32 (0 is Jan 1)
	days := time.YearDay()
	b := (0.9863) * (days - 81)
	return 9.87 * math.Sin(2 * b) - 7.53 * math.Cos(b) - 1.5 * math.Sin(b)
}

func timeCorrectionFactor(long float64) float64 {
	return 4 * (long - localStandardTimeMeridian()) + equationOfTime()
}

// returns the angle of the sun in degrees
func hourAngle() int {
	return 15 * (localSolarTime() - 12)
}
