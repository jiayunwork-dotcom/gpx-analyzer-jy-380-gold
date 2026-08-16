package geo

import (
	"math"
	"testing"
)

func TestHaversineKnown(t *testing.T) {
	a := Point{Lat: 0, Lon: 0}
	b := Point{Lat: 1, Lon: 0}
	got := Haversine(a, b)
	want := 111.19
	if math.Abs(got-want) > 0.5 {
		t.Errorf("Haversine(0,0 -> 1,0) = %.3f km, want ~%.2f km (tol 0.5)", got, want)
	}
}

func TestDistanceInsufficient(t *testing.T) {
	_, err := Distance([]Point{{Lat: 1, Lon: 2}})
	if err == nil {
		t.Fatal("expected error for insufficient points, got nil")
	}
	if err.Error() != "need at least 2 points" {
		t.Errorf("expected error %q, got %v", "need at least 2 points", err)
	}
}

func TestDistanceSum(t *testing.T) {
	pts := []Point{
		{Lat: 0, Lon: 0},
		{Lat: 1, Lon: 0},
		{Lat: 2, Lon: 0},
	}
	got, err := Distance(pts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 2 * 111.19
	if math.Abs(got-want) > 1.0 {
		t.Errorf("Distance = %.3f km, want ~%.2f km", got, want)
	}
}
