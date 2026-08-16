// Package gpx parses GPS track points from CSV into Point values.
package gpx

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// Point is a single GPS track fix.
type Point struct {
	Lat, Lon, Ele float64
	Time          string
}

// ParseTrack reads CSV track points from r. The first row is treated as a
// header and skipped when its first cell equals "lat". Each data row must have
// at least two fields (lat, lon); lat and lon must parse as float64.
func ParseTrack(r io.Reader) ([]Point, error) {
	cr := csv.NewReader(r)
	cr.FieldsPerRecord = -1 // allow rows with varying field counts
	cr.TrimLeadingSpace = true

	rows, err := cr.ReadAll()
	if err != nil {
		return nil, err
	}

	points := make([]Point, 0, len(rows))
	for i, row := range rows {
		if i == 0 && len(row) > 0 && row[0] == "lat" {
			continue
		}
		if len(row) < 2 {
			return nil, fmt.Errorf("short row")
		}
		lat, err := strconv.ParseFloat(row[0], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid coordinate")
		}
		lon, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid coordinate")
		}
		p := Point{Lat: lat, Lon: lon}
		if len(row) >= 3 && row[2] != "" {
			ele, err := strconv.ParseFloat(row[2], 64)
			if err != nil {
				return nil, fmt.Errorf("invalid coordinate")
			}
			p.Ele = ele
		}
		if len(row) >= 4 {
			p.Time = row[3]
		}
		points = append(points, p)
	}
	return points, nil
}
