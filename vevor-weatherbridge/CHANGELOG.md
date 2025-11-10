# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.3] - 2025-11-10

### Fixed
- **CRITICAL**: Fixed MQTT callback signatures for API v2 compatibility
  - Updated `on_connect` to accept 5 parameters (client, userdata, flags, reason_code, properties)
  - Updated `on_disconnect` to accept 5 parameters (client, userdata, flags, reason_code, properties)
  - Updated `on_publish` to accept 5 parameters (client, userdata, mid, reason_codes, properties)
  - Resolved TypeError that prevented MQTT connection from working

### Added
- Comprehensive test suite with pytest (test_weatherstation.py)
- Tests for MQTT callbacks, unit conversions, endpoint functionality
- Tests for metric conversion, timezone handling, error scenarios
- pytest and pytest-mock added to requirements.txt
- Post-edit hook for automatic testing after code changes

## [0.1.2] - 2025-11-10

### Fixed
- MQTT connection state tracking - data now publishes correctly to MQTT broker
- MQTT callbacks now properly track connection status and handle reconnection
- Enhanced error logging with detailed MQTT publish status information

### Added
- Configurable log level via addon options (DEBUG, INFO, WARNING, ERROR)
- MQTT connection state callbacks (on_connect, on_disconnect, on_publish)
- Detailed debug logging for MQTT publish operations
- Automatic MQTT reconnection attempts when disconnected

### Changed
- Upgraded to MQTT Client API v2 to fix deprecation warnings
- Improved logging with connection status and publish confirmations
- Enhanced weather data processing logs with sensor counts

## [0.1.1] - 2025-11-10

### Fixed
- MQTT broker auto-detection now works without bashio dependency
- Replaced bashio calls with standard bash and jq for better compatibility
- Added jq package to Docker image for JSON config parsing
- Implemented multi-method MQTT broker detection (secrets, ping, fallback)

### Changed
- Startup script now uses plain bash instead of bashio helpers
- Improved logging format for better clarity in HA logs

## [0.1.0] - 2025-11-10

### Added
- Initial release as Home Assistant Add-on
- Automatic MQTT discovery for all weather sensors
- Support for both metric and imperial units
- Device grouping in Home Assistant (all sensors under one device)
- Automatic detection of Home Assistant's internal MQTT broker
- Support for external MQTT broker configuration
- Optional Weather Underground data forwarding
- Multi-architecture support (amd64, armv7, aarch64, armhf, i386)
- Comprehensive user documentation (DOCS.md)
- Configurable device name, manufacturer, and model
- Timezone support for local time display
- Enhanced logging for add-on environment

### Features
- Temperature sensor with device class
- Humidity sensor with device class
- Barometric pressure sensor with device class
- Dew point sensor with device class
- Rainfall sensor with device class
- Daily rainfall sensor with device class
- Wind direction sensor
- Wind speed sensor with device class
- Wind gust speed sensor with device class
- UV index sensor
- Solar radiation sensor with device class

### Technical
- Built on Home Assistant base Python Alpine images
- Uses bashio for add-on configuration management
- Flask HTTP server for weather station endpoint
- Paho MQTT client for Home Assistant communication
- DNS resolution for Weather Underground forwarding
- Graceful error handling and connection resilience

## [Unreleased]

### Planned
- Historical data storage
- Graphical weather dashboard
- Configurable sensor update intervals
- Support for additional weather station models
- Advanced alerting capabilities
- Data export functionality
