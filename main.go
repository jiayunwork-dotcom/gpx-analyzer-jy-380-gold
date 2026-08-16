package main

import (
	"flag"
	"fmt"
	"os"

	"gpx-analyzer/internal/gpx"
	"gpx-analyzer/internal/geo"
	"gpx-analyzer/internal/stats"
)

const usageMsg = "usage: go run . -input track.csv"

func main() {
	input := flag.String("input", "", "path to GPS track CSV file")
	flag.Parse()

	if *input == "" {
		fmt.Fprintln(os.Stderr, usageMsg)
		os.Exit(2)
	}

	f, err := os.Open(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	points, err := gpx.ParseTrack(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	st, err := stats.Analyze(points)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	dist, err := geo.Distance(points)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("points:        %d\n", len(points))
	fmt.Printf("distance:      %.2f km\n", dist)
	fmt.Printf("duration:      %.2f min\n", st.DurationMin)
	fmt.Printf("avg speed:     %.2f km/h\n", st.AvgSpeed)
	fmt.Printf("max speed:     %.2f km/h\n", st.MaxSpeed)
	fmt.Printf("elevation gain: %.2f m\n", st.EleGain)
}
