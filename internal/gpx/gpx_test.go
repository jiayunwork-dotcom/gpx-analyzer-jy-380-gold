package gpx

import (
	"strings"
	"testing"
)

func TestParseTrackValid(t *testing.T) {
	in := "lat,lon,ele,time\n" +
		"48.137,11.575,520,2024-01-01T10:00:00\n" +
		"48.140,11.580,528,2024-01-01T10:05:00\n" +
		"48.145,11.590,510,2024-01-01T10:10:00\n"
	pts, err := ParseTrack(strings.NewReader(in))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(pts) < 2 {
		t.Fatalf("expected at least 2 points, got %d", len(pts))
	}
	if pts[0].Lat != 48.137 || pts[0].Lon != 11.575 {
		t.Errorf("unexpected first point: %+v", pts[0])
	}
	if pts[0].Ele != 520 {
		t.Errorf("unexpected first ele: got %v want 520", pts[0].Ele)
	}
	if pts[0].Time != "2024-01-01T10:00:00" {
		t.Errorf("unexpected first time: got %q", pts[0].Time)
	}
	if pts[1].Ele != 528 {
		t.Errorf("unexpected second ele: got %v want 528", pts[1].Ele)
	}
}

func TestParseTrackBadFloat(t *testing.T) {
	in := "lat,lon,ele,time\n" +
		"notanumber,11.575,520,2024-01-01T10:00:00\n"
	_, err := ParseTrack(strings.NewReader(in))
	if err == nil {
		t.Fatal("expected error for bad float coordinate, got nil")
	}
	if err.Error() != "invalid coordinate" {
		t.Errorf("expected error %q, got %v", "invalid coordinate", err)
	}
}

func TestParseTrackShortRow(t *testing.T) {
	in := "lat,lon,ele,time\n48.137\n"
	_, err := ParseTrack(strings.NewReader(in))
	if err == nil {
		t.Fatal("expected error for short row, got nil")
	}
	if err.Error() != "short row" {
		t.Errorf("expected error %q, got %v", "short row", err)
	}
}
