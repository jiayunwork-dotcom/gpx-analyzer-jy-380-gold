// Package stats computes summary statistics for a GPS track.
package stats

import (
	"fmt"
	"time"

	"gpx-analyzer/internal/gpx"
	"gpx-analyzer/internal/geo"
)

// Point is an alias for the track point type defined in the gpx package.
type Point = gpx.Point

// Stats holds the computed summary for a track.
type Stats struct {
	Distance    float64 // total path length in km
	DurationMin float64 // elapsed time between first and last fix, in minutes
	AvgSpeed    float64 // km/h over the whole track
	MaxSpeed    float64 // km/h of the fastest segment
	EleGain     float64 // sum of positive consecutive elevation differences, in meters
}

const timeLayout = "2006-01-02T15:04:05"

// Analyze computes track statistics. It requires at least 2 points.
func Analyze(points []Point) (Stats, error) {
	if len(points) < 2 {
		return Stats{}, fmt.Errorf("need at least 2 points")
	}

	dist, err := geo.Distance(points)
	if err != nil {
		return Stats{}, err
	}
	var s Stats
	s.Distance = dist

	// Elevation gain: sum of positive consecutive elevation differences.
	for i := 1; i < len(points); i++ {
		diff := points[i].Ele - points[i-1].Ele
		if diff > 0 {
			s.EleGain += diff
		}
	}

	// Duration and speeds from timestamps, when both ends parse.
	t0, err0 := time.Parse(timeLayout, points[0].Time)
	t1, err1 := time.Parse(timeLayout, points[len(points)-1].Time)
	if err0 == nil && err1 == nil {
		dur := t1.Sub(t0)
		s.DurationMin = dur.Minutes()
		hours := dur.Hours()
		if hours > 0 {
			s.AvgSpeed = dist / hours
		}
	}

	// Max segment speed: fastest consecutive pair with parseable times.
	for i := 1; i < len(points); i++ {
		ta, ea := time.Parse(timeLayout, points[i-1].Time)
		tb, eb := time.Parse(timeLayout, points[i].Time)
		if ea != nil || eb != nil {
			continue
		}
		segHours := tb.Sub(ta).Hours()
		if segHours <= 0 {
			continue
		}
		segDist := geo.Haversine(points[i-1], points[i])
		sp := segDist / segHours
		if sp > s.MaxSpeed {
			s.MaxSpeed = sp
		}
	}

	return s, nil
}
