# VEVOR Weatherbridge v0.1.6

## Bashio Integration - Following HA Best Practices ✅

This release migrates from manual jq/curl scripts to the official **bashio library**, following Home Assistant addon best practices.

## Changes

### Migrated to Bashio

- **run.sh** now uses `#!/usr/bin/with-contenv bashio` shebang
- Configuration reading via `bashio::config` functions
- MQTT service discovery via `bashio::services.available` and `bashio::services`
- Logging via `bashio::log.*` functions (info, fatal, etc.)
- Removed jq dependency from Dockerfile (bashio provides all needed functionality)

### Fixed

- **CRITICAL**: SUPERVISOR_TOKEN access now handled correctly by bashio
  - No more "SUPERVISOR_TOKEN not available" errors
  - Bashio automatically manages Supervisor API authentication
  - Proper fallback when MQTT service not available

## Why Bashio?

- ✅ **Best Practice**: Home Assistant's official library for addon development
- ✅ **Pre-installed**: Included in all official HA base images
- ✅ **Automatic Auth**: Handles SUPERVISOR_TOKEN automatically
- ✅ **250+ Functions**: Comprehensive API for HA addon operations
- ✅ **Maintained**: Actively developed by HA Community Add-ons team

## Full Changelog

See [CHANGELOG.md](https://github.com/lenucksi/VevorWeatherbridge/blob/main/vevor-weatherbridge/CHANGELOG.md) for complete version history.

## Contributors

- @lenucksi (Lenucksi)

**Full Commit**: <https://github.com/lenucksi/VevorWeatherbridge/commit/6f517e7>
