# Changelog

All notable changes to the VEVOR Weather Station Bridge (Go) will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.4](https://github.com/lenucksi/VevorWeatherbridge/compare/vevor-weatherbridge-go-v0.1.3...vevor-weatherbridge-go-v0.1.4) (2025-11-29)


### Features

* **security:** re-add AppArmor profile with Supervisor API support ([d1c7dd5](https://github.com/lenucksi/VevorWeatherbridge/commit/d1c7dd58ddd86c53c19e93688f0f694dddeed353))


### Bug Fixes

* **go:** normalize timestamps with single-digit time components ([d1c7dd5](https://github.com/lenucksi/VevorWeatherbridge/commit/d1c7dd58ddd86c53c19e93688f0f694dddeed353))

## [0.1.3](https://github.com/lenucksi/VevorWeatherbridge/compare/vevor-weatherbridge-go-v0.1.2...vevor-weatherbridge-go-v0.1.3) (2025-11-24)


### Bug Fixes

* **security:** remove AppArmor profiles blocking Supervisor API ([6978c05](https://github.com/lenucksi/VevorWeatherbridge/commit/6978c0525cbb747bc7cb242ae33ccd9040716b89))

## [0.1.2](https://github.com/lenucksi/VevorWeatherbridge/compare/vevor-weatherbridge-go-v0.1.1...vevor-weatherbridge-go-v0.1.2) (2025-11-24)


### Bug Fixes

* **mqtt:** improve MQTT configuration detection and error messages ([ff2494a](https://github.com/lenucksi/VevorWeatherbridge/commit/ff2494a084f9759ea16eaa5146b5fef12fb3fc64))
* **security:** update AppArmor profiles for HAOS compatibility ([8860140](https://github.com/lenucksi/VevorWeatherbridge/commit/8860140cb733575e9269c9cfeb0f050b60fbcc6e))

## [0.1.1](https://github.com/lenucksi/VevorWeatherbridge/compare/vevor-weatherbridge-go-v0.1.0...vevor-weatherbridge-go-v0.1.1) (2025-11-23)


### Features

* add Go implementation of VEVOR Weather Station Bridge ([569ea2f](https://github.com/lenucksi/VevorWeatherbridge/commit/569ea2fafa92d5e78547388147fe8a1aabe2676e))
* **go:** add AppArmor security profile for Go addon ([6c9494c](https://github.com/lenucksi/VevorWeatherbridge/commit/6c9494c8b1ce436f008a59d31ad97690e7f8f0b3))


### Bug Fixes

* **go:** resolve golangci-lint errcheck warnings ([d1bd6be](https://github.com/lenucksi/VevorWeatherbridge/commit/d1bd6be799d6a2bdfc39ab096283dbf0fa877f47))
* **linter:** codelinting changes Docker,Py,GHA ([dd6a8d9](https://github.com/lenucksi/VevorWeatherbridge/commit/dd6a8d95e16a102d8aa0e62794de73c7f445cbfb))
* **markdown:** markdown linting fixes ([d800575](https://github.com/lenucksi/VevorWeatherbridge/commit/d8005750dfc4dac90ad400db5e9c54129fd5044b))

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
