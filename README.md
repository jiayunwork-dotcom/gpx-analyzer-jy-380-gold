# gpx-analyzer

A self-contained, offline Go tool that parses GPS track points from CSV and
computes distance, duration, speed, and elevation gain statistics.

## CSV format

Columns: `lat,lon,ele,time`

- `lat`, `lon`: required decimal degrees (float64)
- `ele`: optional elevation in meters
- `time`: optional timestamp, format `2006-01-02T15:04:05`

Example `example/track.csv`:

```
lat,lon,ele,time
48.137,11.575,520,2024-01-01T10:00:00
48.140,11.580,528,2024-01-01T10:05:00
```

## Usage

```
go run . -input track.csv
```

Exit codes:

- `2` — missing or invalid arguments (usage printed to stderr)
- `1` — bad input data (error printed to stderr)
- `0` — success

## Packages

- `internal/gpx` — CSV parsing into `Point` values
- `internal/geo` — haversine great-circle distance
- `internal/stats` — track statistics (distance, duration, speeds, elevation gain)

## Build

```
go build ./...
go test ./...
```

## Docker

```
docker build -t gpx-analyzer .
docker run --rm -v "$PWD/example:/data" gpx-analyzer sh -c "go run . -input /data/track.csv"
```
