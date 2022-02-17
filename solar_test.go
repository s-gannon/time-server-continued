package main

import (
	"testing"
)

// unit test for calculating solar time
func TestSolarTime(t *testing.T) {
	//if SolarTime() != {
	//	t.Error()
	//}
}

// unit test for finding the time correction factor between standard & solar time
func TestTimeCorrectionFactor(t *testing.T) {
	//if timeCorrectionFactor() != {
	//	t.Error()
	//}
}

// unit test for obtaining standard time zone meridian
func TestTimeMeridian(t *testing.T) {
	if localStandardTimeMeridian(time.LoadLocation("America/Chicago")) != 6 {
		t.Error()
	}
}

// unit test for equation of time
// result shown if calculating in degrees, answer should be -1.079397746 if formula accounts for radians
func TestEquationOfTime(t *testing.T) {
	if equationOfTime(32) != 15.49075338 {
		t.Error()
	}
}

// unit test for calculating absolute value
func TestIntAbs(t *testing.T) {
	if intAbs(-10) != 10 {
		t.Error()
	}
}
