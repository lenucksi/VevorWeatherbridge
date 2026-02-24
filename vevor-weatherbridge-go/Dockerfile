# SPDX-License-Identifier: GPL-3.0-or-later
# Copyright (C) 2025 Lenucksi
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

# Stage 1: Build the Go binary
ARG BUILD_FROM
FROM golang:1.26-alpine@sha256:d4c4845f5d60c6a974c6000ce58ae079328d03ab7f721a0734277e69905473e5 AS builder

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
ARG VERSION=dev
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} GOARM=${TARGETVARIANT#v} \
    go build -ldflags="-s -w -X main.Version=${VERSION}" -o weatherbridge .

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
