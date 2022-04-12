package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
  solar "timeserver/solarconversion"
  "time"
)

func main() {
        // parse variable flags
        unix := flag.Int64("t", 0, "Local time as a unix timestamp.")
        long := flag.Float64("l", 87.037403, "Longitude to use in the calculation.") // default longitude is the location of the solar furnace at Valparaiso University
        zone := flag.Int("z", -6, "The current time zone: UTC +- how many hours. Eg. UTC-6 == -6.")

        flag.Usage = func() {
                fmt.Println("Usage: solar [OPTIONS]")
                fmt.Println("Returns: Unix timestamp of the current solar time.")
                flag.PrintDefaults()
        }

        flag.Parse()

        if *unix <= 0 {
                fmt.Println("Specifying a valid unix time is required.")
                os.Exit(1) // return an erring exit code
        } else {
                meridian := solar.LocalStandardTimeMeridian(*zone)
                eqnOfTime := 60 * solar.EquationOfTime(time.Now().YearDay()) // convert from minutes to seconds

                fmt.Println(strconv.FormatInt(
                                *unix + solar.TimeCorrectionFactor(*long, meridian, eqnOfTime),
                                10)) // magic number 10 is to format in base 10
                os.Exit(0) // return a successful exit code
        }
}
