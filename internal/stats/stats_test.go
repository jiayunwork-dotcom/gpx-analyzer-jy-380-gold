package stats

import (
	"math"
	"testing"
)

func TestAnalyzeEleGain(t *testing.T) {
	pts := []Point{
		{Lat: 0, Lon: 0, Ele: 100},
		{Lat: 0.01, Lon: 0, Ele: 150},
		{Lat: 0.02, Lon: 0, Ele: 120},
		{Lat: 0.03, Lon: 0, Ele: 200},
	}
	s, err := Analyze(pts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Positive diffs: (150-100)=50, (200-120)=80 => 130
	if s.EleGain != 130 {
		t.Errorf("EleGain = %v, want 130", s.EleGain)
	}
}

func TestAnalyzeMaxSpeed(t *testing.T) {
	pts := []Point{
		{Lat: 0, Lon: 0, Time: "2024-01-01T10:00:00"},
		{Lat: 1, Lon: 0, Time: "2024-01-01T11:00:00"},    // ~111.19 km in 1h
		{Lat: 1.001, Lon: 0, Time: "2024-01-01T11:01:00"}, // ~0.111 km in 1m
	}
	s, err := Analyze(pts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(s.MaxSpeed-111.19) > 0.5 {
		t.Errorf("MaxSpeed = %.3f km/h, want ~111.19 (tol 0.5)", s.MaxSpeed)
	}
}

func TestAnalyzeDuration(t *testing.T) {
	pts := []Point{
		{Lat: 0, Lon: 0, Time: "2024-01-01T10:00:00"},
		{Lat: 0.1, Lon: 0, Time: "2024-01-01T11:30:00"},
	}
	s, err := Analyze(pts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.DurationMin != 90 {
		t.Errorf("DurationMin = %v, want 90", s.DurationMin)
	}
}

func TestAnalyzeInsufficient(t *testing.T) {
	_, err := Analyze([]Point{{Lat: 1, Lon: 2}})
	if err == nil {
		t.Fatal("expected error for insufficient points, got nil")
	}
}

func TestAnalyzeNoTime(t *testing.T) {
	pts := []Point{
		{Lat: 0, Lon: 0, Ele: 100},
		{Lat: 0.01, Lon: 0, Ele: 130},
	}
	s, err := Analyze(pts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.DurationMin != 0 || s.AvgSpeed != 0 || s.MaxSpeed != 0 {
		t.Errorf("expected zero duration/speeds with no time, got dur=%v avg=%v max=%v",
			s.DurationMin, s.AvgSpeed, s.MaxSpeed)
	}
	if s.EleGain != 30 {
		t.Errorf("EleGain = %v, want 30", s.EleGain)
	}
}
