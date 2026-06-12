# Changelog

All notable changes to the VEVOR Weather Station Bridge (Go) will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.9](https://github.com/lenucksi/VevorWeatherbridge/compare/vevor-weatherbridge-go-v0.1.8...vevor-weatherbridge-go-v0.1.9) (2026-06-12)


### Features

* add Go implementation of VEVOR Weather Station Bridge ([a10e413](https://github.com/lenucksi/VevorWeatherbridge/commit/a10e413451c711f5ebbfe6312d769112dbb9420b))
* add native Go fuzzing tests for input parsing and conversions ([0e54a11](https://github.com/lenucksi/VevorWeatherbridge/commit/0e54a1166638a080affa5db9d47ec58a32540872))
* add state_class to all MQTT Discovery sensors for Long-Term Statistics ([54ded25](https://github.com/lenucksi/VevorWeatherbridge/commit/54ded2560bdcf7c8b8334cebf45aae847bedf213))
* **go:** add AppArmor security profile for Go addon ([aaa0e6f](https://github.com/lenucksi/VevorWeatherbridge/commit/aaa0e6f863e3938ecc9001ab0b6bfd3735c78c9a))
* initialize Backlog.md project management ([bc0518b](https://github.com/lenucksi/VevorWeatherbridge/commit/bc0518b187f71042cbdf6ae6af262d8ee9f44be6))
* migrate from pre-commit (Python) to prek (Rust) ([c612a8f](https://github.com/lenucksi/VevorWeatherbridge/commit/c612a8f62266f8db917b5086d503b4c80e6c7e0e))
* **security:** re-add AppArmor profile with Supervisor API support ([822df18](https://github.com/lenucksi/VevorWeatherbridge/commit/822df182cb2c5575665f20adf860b0c4674bc957))
* upgrade HA base image from Alpine 3.19 to 3.21 ([7b6ce2e](https://github.com/lenucksi/VevorWeatherbridge/commit/7b6ce2e11b3ea41cdcba33f0700bc15a8fb09fbf))


### Bug Fixes

* add default BUILD_FROM for Quay.io builds (amd64) ([e5f4323](https://github.com/lenucksi/VevorWeatherbridge/commit/e5f4323bd74e48777994c5609d6c154d19624c04))
* add error handling for MQTT availability status publishing ([1ea8993](https://github.com/lenucksi/VevorWeatherbridge/commit/1ea8993e443cc429272402d7dc01a3711651f5d4))
* add LICENSE to Go module for go-licenses save ([d8c3e86](https://github.com/lenucksi/VevorWeatherbridge/commit/d8c3e865e384b7c30a783d89640b2029aa8841e1))
* drop deprecated HA architectures armv7, armhf, i386 ([309b067](https://github.com/lenucksi/VevorWeatherbridge/commit/309b06799d4e6490d39ad281b15092732f079ec8))
* Generate proper FOSS NOTICE with Go + Alpine attribution ([ce37166](https://github.com/lenucksi/VevorWeatherbridge/commit/ce37166de498662aeabdf59acf0dabe7c1a6ae87))
* **go:** normalize timestamps with single-digit time components ([822df18](https://github.com/lenucksi/VevorWeatherbridge/commit/822df182cb2c5575665f20adf860b0c4674bc957))
* **go:** resolve golangci-lint errcheck warnings ([cd9380c](https://github.com/lenucksi/VevorWeatherbridge/commit/cd9380caa2e23b61d913cba62eea93591ab870a6))
* handle non-zero-padded timestamps from weather station ([3b335fe](https://github.com/lenucksi/VevorWeatherbridge/commit/3b335fe9dd9847218d51d1637809cb9558e53b52))
* inject correct version into Go binary at build time via ldflags ([15d0e42](https://github.com/lenucksi/VevorWeatherbridge/commit/15d0e42dcbc3f5745405f630ff5baee201fc7824))
* **linter:** codelinting changes Docker,Py,GHA ([f748bbf](https://github.com/lenucksi/VevorWeatherbridge/commit/f748bbf30d52e57134702dd58fe231f69fd41e19))
* **markdown:** markdown linting fixes ([25daeb6](https://github.com/lenucksi/VevorWeatherbridge/commit/25daeb6db7be2b0526962c01a2bbecfa47fd1bc5))
* **mqtt:** improve MQTT configuration detection and error messages ([c3de8f6](https://github.com/lenucksi/VevorWeatherbridge/commit/c3de8f6f8c4acf75f81a68c65a068ce55bd57804))
* prevent go-licenses recursive path error ([3545237](https://github.com/lenucksi/VevorWeatherbridge/commit/354523712e40ef12f8e4e039e789ac028f9d3e27))
* resolve failing GH Actions, CVE flood, and SBOM warnings ([d71ff04](https://github.com/lenucksi/VevorWeatherbridge/commit/d71ff0426ee03d693c3a7a8d5d0f981666ab660b))
* **security:** remove AppArmor profiles blocking Supervisor API ([fda2642](https://github.com/lenucksi/VevorWeatherbridge/commit/fda26425548b68606ee1579c22ad4cc7f3e23a57))
* **security:** update AppArmor profiles for HAOS compatibility ([d2ebebb](https://github.com/lenucksi/VevorWeatherbridge/commit/d2ebebbce43b5106932b3f74171b932c901f25c3))
* update Go builder to 1.26.3, fix Clair report script ([4d0e4c5](https://github.com/lenucksi/VevorWeatherbridge/commit/4d0e4c59860588a49f3754d3a9efa2a38a958046))
* update Go deps (x/net v0.53.0), add Clair API report to CI ([8dcc2eb](https://github.com/lenucksi/VevorWeatherbridge/commit/8dcc2eb9e1302ef1abf68534975ac0378562fba5))
* use Basic auth for Quay API (robot password, not OAuth token) ([1d104c9](https://github.com/lenucksi/VevorWeatherbridge/commit/1d104c99b60b7af6254f9b65f706a28eefd325c7))

## [0.1.8](https://github.com/lenucksi/VevorWeatherbridge/compare/vevor-weatherbridge-go-v0.1.7...vevor-weatherbridge-go-v0.1.8) (2026-02-24)


### Bug Fixes

* inject correct version into Go binary at build time via ldflags ([9e1aaaf](https://github.com/lenucksi/VevorWeatherbridge/commit/9e1aaaf253e64f20a7db03a3fdfab88ec5169f2b))

## [0.1.7](https://github.com/lenucksi/VevorWeatherbridge/compare/vevor-weatherbridge-go-v0.1.6...vevor-weatherbridge-go-v0.1.7) (2026-02-24)


### Features

* add state_class to all MQTT Discovery sensors for Long-Term Statistics ([726bbc9](https://github.com/lenucksi/VevorWeatherbridge/commit/726bbc9761002fa9eb3a90c250f36bfc698cd13e))

## [0.1.6](https://github.com/lenucksi/VevorWeatherbridge/compare/vevor-weatherbridge-go-v0.1.5...vevor-weatherbridge-go-v0.1.6) (2026-02-07)


### Features

* add native Go fuzzing tests for input parsing and conversions ([20b4383](https://github.com/lenucksi/VevorWeatherbridge/commit/20b43835ccf7e0f94fdedfb48f180f87ede6766c))

## [0.1.5](https://github.com/lenucksi/VevorWeatherbridge/compare/vevor-weatherbridge-go-v0.1.4...vevor-weatherbridge-go-v0.1.5) (2025-12-01)


### Bug Fixes

* add error handling for MQTT availability status publishing ([f57acbe](https://github.com/lenucksi/VevorWeatherbridge/commit/f57acbebe659af9cc01a3e4b4c7f3b1444c3e5b3))
* handle non-zero-padded timestamps from weather station ([6765335](https://github.com/lenucksi/VevorWeatherbridge/commit/67653354aa65c0f1103904825a671469a63d1355))

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
