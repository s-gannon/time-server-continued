package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	solar "timeserver/solarconversion"
)

// port to run the webserver on
const PORT = "8080"

// variables in the web clock template
type Data struct {
		ClockDateTime string
		Clock string
		SolarTime string
		DateDateTime string
		Date string
}

func main() {
	// load template file for the web clock
	siteTemplate := template.Must(template.ParseFiles("index.html"))

	// serve the clock template file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleClockSite(w, siteTemplate)
	})

	// serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// run the server and print status notifications
	fmt.Printf("Starting HTTP server on port %s...\n", PORT)
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
			fmt.Printf("Error starting the HTTP server on port %s. Something else may be running already...\n", PORT)
	}
}

// process time variables for clock template
func handleClockSite(w http.ResponseWriter, siteTemplate *template.Template) {
		data := getData()
		
		siteTemplate.Execute(w, data)
}

func getData() Data {
	current := time.Now().Unix()
	solarTime := convertToSolar(current)

	currentFormat := formatTime(current)
	solarFormat := formatTime(solarTime)

	date := time.Unix(current, 0).Format("January 1, 2006")

	data := Data {
		ClockDateTime: currentFormat,
		Clock: currentFormat,
		SolarTime: solarFormat,
		DateDateTime: date,
		Date: date,
	}

	return data
}

// Helper function to calculate the correction factor to solar time
func convertToSolar(unix int64) int64 {
	const zone = -6
	const long = 87.037403

	meridian := solar.LocalStandardTimeMeridian(zone)
	eqnOfTime := 60 * solar.EquationOfTime(time.Now().YearDay())

	return unix + solar.TimeCorrectionFactor(long, meridian, eqnOfTime)
}

// parse time to readable form
func formatTime(unix int64) string {
	t := time.Unix(unix, 0)
	return t.Format("03:04:05")
}
