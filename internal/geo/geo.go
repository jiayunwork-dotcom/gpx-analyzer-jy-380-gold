// Package geo provides great-circle distance calculations for track points.
package geo

import (
	"fmt"
	"math"

	"gpx-analyzer/internal/gpx"
)

// Point is an alias for the track point type defined in the gpx package.
type Point = gpx.Point

// EarthRadiusKm is the mean Earth radius in kilometers.
const EarthRadiusKm = 6371.0

// Haversine returns the great-circle distance in kilometers between a and b.
func Haversine(a, b Point) float64 {
	lat1 := a.Lat * math.Pi / 180
	lat2 := b.Lat * math.Pi / 180
	dLat := (b.Lat - a.Lat) * math.Pi / 180
	dLon := (b.Lon - a.Lon) * math.Pi / 180

	h := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(h), math.Sqrt(1-h))
	return EarthRadiusKm * c
}

// Distance returns the total path length in kilometers by summing the
// haversine distance between consecutive points. It requires at least 2 points.
func Distance(points []Point) (float64, error) {
	if len(points) < 2 {
		return 0, fmt.Errorf("need at least 2 points")
	}
	var total float64
	for i := 1; i < len(points); i++ {
		total += Haversine(points[i-1], points[i])
	}
	return total, nil
}
