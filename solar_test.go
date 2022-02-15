package main

import (
	"testing"
)

func TestTimeMeridian(t *testing.T) {
	if localStandardTimeMeridian(time.LoadLocation("America/Chicago")) != 6 {
		t.Error()
	}
}

func TestEquationOfTime(t *testing.T) {
	if equationOfTime(7) == {
	}
}
