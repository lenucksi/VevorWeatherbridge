# VEVOR Weatherbridge v0.1.7

## Critical Fix: MQTT Discovery Now Works! 🎉

This release fixes the issue where sensor data was being published to MQTT but **not appearing in Home Assistant's UI**. All weather sensors now automatically appear under a single device!

## What's New

### Added
- **MQTT Discovery Support**: Sensors now automatically appear in Home Assistant
  - Added required `origin` field with addon name, version, and support URL
  - Added `availability_topic` for online/offline status tracking
  - Configured MQTT Last Will and Testament (LWT) for proper offline detection
  - Availability status published as "online" on connect, "offline" on disconnect

### Fixed
- **CRITICAL**: Device-based MQTT discovery now works properly
  - Home Assistant will automatically create all 11 weather sensors
  - Sensors grouped under single device in HA device registry
  - `device_class` field only included when applicable (not for UV Index, Wind Direction)
  - All sensors now appear in HA UI without manual configuration

### Changed
- Improved MQTT discovery payload structure per HA 2025 requirements
- Better device availability tracking with retained messages

## How It Works

1. **On MQTT connect**: Publishes "online" to availability topic
2. **On disconnect/crash**: LWT auto-publishes "offline"
3. **Each sensor update**: Publishes discovery config + state
4. **HA automatically**: Creates sensors under single device
5. **Sensors show**: Proper device classes, units, and icons

## Sensors Available

All 11 sensors will automatically appear in Home Assistant under **Settings → Devices → VEVOR Weather Station**:

- 🌡️ Temperature (°C/°F)
- 💧 Humidity (%)
- 🌪️ Barometric Pressure (hPa/inHG)
- 💨 Dew Point (°C/°F)
- 🌧️ Rainfall (mm/in)
- 📊 Daily Rainfall (mm/in)
- 🧭 Wind Direction (°)
- 💨 Wind Speed (km/h/mph)
- 🌬️ Wind Gust Speed (km/h/mph)
- ☀️ UV Index
- ⚡ Solar Radiation (W/m²)

## Installation

### Home Assistant Addon

1. Update the addon to v0.1.7 in Home Assistant
2. Restart the addon
3. Wait for the next weather station update
4. Check **Settings → Devices & Services → MQTT**
5. Your weather station device with all sensors should appear automatically!

### Docker Image

```bash
docker pull ghcr.io/lenucksi/vevor-weatherbridge-amd64:latest
```

Available architectures: `amd64`, `armv7`, `aarch64`, `armhf`, `i386`

## Full Changelog

See [CHANGELOG.md](https://github.com/lenucksi/VevorWeatherbridge/blob/main/vevor-weatherbridge/CHANGELOG.md) for complete version history.

## Contributors

- @lenucksi (Lenucksi)
- @C9H13NO3-dev (Original project creator)

---

**Full Commit**: https://github.com/lenucksi/VevorWeatherbridge/commit/cc325be
