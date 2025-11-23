# Go Migration Start Prompt

Use this prompt to start a new Claude Code session for implementing the Go version of VevorWeatherbridge.

---

## Context

You are working on the VevorWeatherbridge project, a Home Assistant add-on that intercepts VEVOR weather station data and forwards it to Home Assistant via MQTT Discovery.

Read `dev-docs/go-rust-migration-analysis.md` for the complete analysis and effort estimates.

## Goal

Implement a Go version of the weatherstation bridge that:
1. Can coexist with the Python version in the same repository
2. Will appear as a **separate add-on** when the repo is added to Home Assistant
3. Maintains 100% feature parity with the Python version

## Repository Structure Target

```
VevorWeatherbridge/
├── repository.yaml                    # Lists both add-ons
├── vevor-weatherbridge/               # Python version (existing)
│   ├── config.yaml                    # slug: vevor-weatherbridge
│   ├── Dockerfile
│   ├── weatherstation.py
│   └── ...
├── vevor-weatherbridge-go/            # Go version (NEW)
│   ├── config.yaml                    # slug: vevor-weatherbridge-go
│   ├── Dockerfile                     # FROM scratch, single binary
│   ├── build.json
│   ├── run.sh
│   ├── main.go
│   ├── config.go
│   ├── convert.go
│   ├── compass.go
│   ├── mqtt.go
│   ├── weather.go                     # WU forwarding
│   ├── go.mod
│   ├── go.sum
│   ├── DOCS.md
│   ├── CHANGELOG.md
│   ├── icon.png                       # Copy from Python version
│   └── logo.png
└── ...
```

## Implementation Requirements

### 1. Feature Parity

The Go version MUST implement:
- HTTP endpoint at `/weatherstation/updateweatherstation.php`
- All unit conversions (F→C, inHg→hPa, mph→km/h, in→mm)
- 16-point compass rose conversion (`degrees_to_cardinal`)
- MQTT publishing with Home Assistant Discovery format
- Device grouping (all sensors under one device)
- Availability topics (Last Will and Testament)
- Weather Underground forwarding (optional, via DNS bypass)
- Configurable logging levels
- Environment variable configuration (same as Python version)

### 2. Add-on Configuration

Create `vevor-weatherbridge-go/config.yaml` with:
- `name: "VEVOR Weather Station Bridge (Go)"`
- `slug: "vevor-weatherbridge-go"`
- `version: "0.1.0"`
- Same options schema as Python version
- `image: ghcr.io/lenucksi/vevor-weatherbridge-go-{arch}`

### 3. Dockerfile

Use multi-stage build:
```dockerfile
# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o weatherbridge

# Stage 2: Runtime (minimal)
FROM scratch
COPY --from=builder /build/weatherbridge /weatherbridge
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 80
ENTRYPOINT ["/weatherbridge"]
```

### 4. Build Configuration

Create `vevor-weatherbridge-go/build.json`:
```json
{
  "build_from": {
    "amd64": "golang:1.22-alpine",
    "armv7": "golang:1.22-alpine",
    "aarch64": "golang:1.22-alpine",
    "armhf": "golang:1.22-alpine",
    "i386": "golang:1.22-alpine"
  },
  "args": {}
}
```

### 5. Update repository.yaml

The existing `repository.yaml` should automatically discover both add-ons since they're in separate directories.

### 6. MQTT Discovery Format

Match the Python implementation exactly:
```go
// Config topic: {prefix}/sensor/{device_id}_{sensor_id}/config
// State topic: {prefix}/sensor/{device_id}_{sensor_id}/state
// Attributes topic: {prefix}/sensor/{device_id}_{sensor_id}/attributes

type DiscoveryPayload struct {
    Name                string            `json:"name"`
    StateTopic          string            `json:"state_topic"`
    UniqueID            string            `json:"unique_id"`
    DeviceClass         string            `json:"device_class,omitempty"`
    UnitOfMeasurement   string            `json:"unit_of_measurement,omitempty"`
    StateClass          string            `json:"state_class,omitempty"`
    Icon                string            `json:"icon,omitempty"`
    SuggestedPrecision  int               `json:"suggested_display_precision,omitempty"`
    Device              DeviceInfo        `json:"device"`
    AvailabilityTopic   string            `json:"availability_topic"`
    JSONAttributesTopic string            `json:"json_attributes_topic,omitempty"`
    Origin              OriginInfo        `json:"origin,omitempty"`
}
```

### 7. Testing

Create tests in the Go package:
- `convert_test.go` - Unit conversion tests
- `compass_test.go` - Cardinal direction tests
- `mqtt_test.go` - MQTT payload structure tests

## Reference Files

Read these files to understand the current implementation:
- `vevor-weatherbridge/weatherstation.py` - Main Python implementation
- `vevor-weatherbridge/config.yaml` - Add-on configuration schema
- `vevor-weatherbridge/run.sh` - Bashio option parsing (for env var names)

## Success Criteria

1. `go build` succeeds with no errors
2. `go test ./...` passes
3. Docker build succeeds for all architectures
4. When repo added to HA, both "VEVOR Weather Station Bridge" and "VEVOR Weather Station Bridge (Go)" appear
5. Go version produces identical MQTT messages as Python version
6. Docker image size < 15 MB (vs ~150 MB for Python)

## Do NOT

- Modify the Python version
- Change the repository.yaml format
- Use CGO (must be pure Go for cross-compilation)
- Add any external web frameworks (use net/http stdlib)

---

**Start by reading the reference files, then implement in this order:**
1. `go.mod` and basic project structure
2. `config.go` - Environment variable parsing
3. `convert.go` - Unit conversion functions with tests
4. `compass.go` - Cardinal direction conversion with tests
5. `mqtt.go` - MQTT client and discovery payloads
6. `main.go` - HTTP server and integration
7. `weather.go` - Weather Underground forwarding
8. Add-on configuration files (config.yaml, build.json, Dockerfile)
9. Documentation (DOCS.md, CHANGELOG.md)
