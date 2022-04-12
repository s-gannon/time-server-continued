package solarconversion

import "math"

// Calculate the difference in time (in seconds) between standard time and solar time.
// Note: Equation of time should be in seconds.
func TimeCorrectionFactor(long float64, meridian int, eqnOfTime float64) int64 {
	// round then cast to int64 to prevent incorrect truncation
	return int64(math.Round(4 * (float64(meridian) - long) + eqnOfTime))
}

// Get standard meridian for time zone in degrees.
// Zone: UTC difference; Eg: UTC-6 == -6 & UTC+7 == 7
// Returns: 75 Eastern ; 90 Central ; 105 Mountain ; 120 Pacific
func LocalStandardTimeMeridian(zone int) int {
	// Zone is how far away from UTC a person is.
	return 15 * intAbs(zone)
}

// Calculate the "equation of time" in minutes.
func EquationOfTime(days int) float64 {
	// days: specifically, days since the start of the year. Eg. February 3rd is 34.
	b := float64(days - 1) * 0.9863013699 // 0.9863013699 approx. == 360/365
	return 229.2 * (0.000075 + 0.001868 * Cosd(b) - 0.032077 * Sind(b) - 0.014615 * Cosd(2 * b) - 0.04089 * Sind(2 * b))
}

/* Helper math functions */

// Calculate the absolute value of an integer. This prevents the need to convert to a float64 and back to use the built in abs function.
func intAbs(x int) int {
	if x < 0 {
		x = -x
	}

	return x
}

// convert "x" degrees into radians
func degreesToRadians(x float64) float64 {
	// the long constant is 180 / pi (to the 50th decimal place) calculated in Chez Scheme
	// math.Pi is not used to increase decimal precision
	return x / 57.29577951308232
}

// calculate the cosine of x where x is in degrees
func Cosd(x float64) float64 {
	return math.Cos(degreesToRadians(x))
}

// calculate the sine of x where x is in degrees
func Sind(x float64) float64 {
	return math.Sin(degreesToRadians(x))
}
