# Stage 1: Build the Go binary
ARG BUILD_FROM
FROM golang:1.25-alpine AS builder

WORKDIR /build

# Install CA certificates for HTTPS requests
# hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates tzdata

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY *.go ./

# Build the binary
ARG TARGETARCH
ARG TARGETVARIANT
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} GOARM=${TARGETVARIANT#v} \
    go build -ldflags="-s -w" -o weatherbridge .

# Stage 2: Runtime image using Home Assistant base
# hadolint ignore=DL3006
FROM ${BUILD_FROM}

# Install CA certificates
# hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates tzdata

# Copy the binary from builder
COPY --from=builder /build/weatherbridge /weatherbridge

# Copy run script
COPY run.sh /

# Make sure run.sh is executable
RUN chmod a+x /run.sh

# Expose HTTP port
EXPOSE 80

# Start via run.sh for bashio integration
CMD [ "/run.sh" ]
