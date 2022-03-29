package main

import (
	"math"
	"testing"
)

// unit test for finding the time correction factor between standard & solar time
func TestTimeCorrectionFactor(t *testing.T) {
	result := timeCorrectionFactor(89.4, 90, -300)
	if result != -298 {
		t.Error()
	}
}

// unit test for obtaining standard time zone meridian
func TestTimeMeridian(t *testing.T) {
	if localStandardTimeMeridian(-5) != 75 {
		t.Error()
	}
	if localStandardTimeMeridian(-6) != 90 {
		t.Error()
	}
	if localStandardTimeMeridian(-7) != 105 {
		t.Error()
	}
	if localStandardTimeMeridian(-8) != 120 {
		t.Error()
	}
}

// unit test for equation of time
// comparing constant was calculated using a graphing calculator
func TestEquationOfTime(t *testing.T) {
	if equationOfTime(34) != -13.488456930676538 {
		t.Error()
	}
}

// unit test for calculating absolute value
func TestIntAbs(t *testing.T) {
	if intAbs(-10) != 10 {
		t.Error()
	}
}

// test cosine using degrees
func TestCosd(t *testing.T) {
	if math.Round(Cosd(180)) != -1 {
		t.Error()
	}
}

// test sine using degrees
func TestSind(t *testing.T) {
	if math.Round(Sind(180)) != 0 {
		t.Error()
	}
}
