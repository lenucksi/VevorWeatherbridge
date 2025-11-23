# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- **Windrose Visualization Support**: Enhanced Wind Direction sensor for windrose display compatibility
  - Added `degrees_to_cardinal()` function for 16-point compass direction conversion
  - Wind Direction sensor now includes `cardinal` attribute (e.g., "N", "NNE", "NE")
  - Added compass icon (`mdi:compass-outline`) to Wind Direction sensor
  - Added `suggested_display_precision: 0` to Wind Direction for cleaner display
  - Comprehensive windrose documentation in DOCS.md with installation and configuration guide
  - Example Lovelace configurations (`lovelace-windrose-example.yaml`) with 5 ready-to-use layouts:
    - Basic windrose with default settings
    - Advanced configuration with custom speed ranges and colors
    - Beaufort scale configuration
    - Compact configuration for small spaces
    - Real-time focus configuration with current conditions

### Changed

- Wind Direction sensor attributes now include both `measured_on` timestamp and `cardinal` direction
- Sensor attribute publishing logic now supports sensor-specific attributes via `attributes` dict key
- DOCS.md expanded with detailed windrose card setup instructions and compatibility information

## [0.1.7] - 2025-11-10

### Added

- **MQTT Discovery Support**: Sensors now automatically appear in Home Assistant
  - Added required `origin` field with addon name, version, and support URL
  - Added `availability_topic` for online/offline status tracking
  - Configured MQTT Last Will and Testament (LWT) for proper offline detection
  - Availability status published as "online" on connect, "offline" on disconnect

### Fixed

- **CRITICAL**: Device-based MQTT discovery now works properly
  - Home Assistant will automatically create all weather sensors
  - Sensors grouped under single device in HA device registry
  - `device_class` field only included when applicable (not for UV Index, Wind Direction)
  - All sensors now appear in HA UI without manual configuration

### Changed

- Improved MQTT discovery payload structure per HA 2025 requirements
- Better device availability tracking with retained messages

## [0.1.6] - 2025-11-10

### Changed

- **BREAKING**: Migrated from manual jq/curl to bashio library for all addon operations
  - run.sh now uses `#!/usr/bin/with-contenv bashio` shebang
  - Configuration reading via `bashio::config` functions
  - MQTT service discovery via `bashio::services.available` and `bashio::services`
  - Logging via `bashio::log.*` functions (info, fatal, etc.)
- Removed jq dependency from Dockerfile (bashio provides all needed functionality)
- Simplified MQTT auto-detection logic using bashio's built-in service query
- Improved logging with proper bashio log levels

### Fixed

- **CRITICAL**: SUPERVISOR_TOKEN access now handled correctly by bashio
  - No more "SUPERVISOR_TOKEN not available" errors
  - Bashio automatically manages Supervisor API authentication
  - Proper fallback when MQTT service not available

## [0.1.5] - 2025-11-10

### Added

- Comprehensive test suite with 29 passing tests
- test_run_sh.py for configuration and MQTT auto-detection tests
- VS Code devcontainer setup for local Home Assistant testing
- VS Code tasks for running tests and rebuilding addon
- Complete testing documentation (dev-docs/guides/TESTING.md)
- Improved post-edit hook that runs full test suite

### Fixed

- SUPERVISOR_TOKEN detection with proper fallback handling
- Debug logging for MQTT auto-detection process
- Test runner hook now properly runs all tests before allowing commits

### Changed

- Post-edit hook now runs both test files
- Better error messages in run.sh for MQTT configuration

## [0.1.4] - 2025-11-10

### Fixed

- **CRITICAL**: Fixed MQTT authentication using Supervisor Services API
  - Properly retrieve MQTT credentials from HA Supervisor instead of using empty auth
  - Added `services: - mqtt:need` to config.yaml to declare MQTT service usage
  - Use curl with SUPERVISOR_TOKEN to query /services/mqtt endpoint
  - Extract host, port, username, password from Supervisor API response

### Changed

- Removed broken ping-based MQTT discovery (was finding host but not credentials)
- Improved error messages when MQTT broker not available

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
