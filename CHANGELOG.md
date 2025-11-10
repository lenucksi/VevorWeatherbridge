# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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
