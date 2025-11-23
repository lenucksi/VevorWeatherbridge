# Changelog

All notable changes to the VEVOR Weather Station Bridge (Go) will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2024-XX-XX

### Added

- Initial Go implementation of the VEVOR Weather Station Bridge
- HTTP endpoint at `/weatherstation/updateweatherstation.php`
- MQTT publishing with Home Assistant MQTT Discovery
- All 11 weather sensors (temperature, humidity, pressure, wind, rain, UV, solar)
- Unit conversions (Fahrenheit→Celsius, inHg→hPa, mph→km/h, inches→mm)
- 16-point compass rose for wind direction
- Weather Underground forwarding with DNS bypass
- Configurable logging levels
- Health check endpoint at `/health`
- Full feature parity with Python version

### Technical Details

- Pure Go implementation with single external dependency (paho.mqtt.golang)
- Multi-stage Docker build with minimal runtime image
- ~15 MB image size vs ~150 MB for Python version
- Support for all 5 Home Assistant architectures (amd64, armv7, aarch64, armhf, i386)
