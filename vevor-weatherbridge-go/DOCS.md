# VEVOR Weather Station Bridge (Go)

This is a Go implementation of the VEVOR Weather Station Bridge for Home Assistant.

## Overview

This add-on intercepts weather data from VEVOR weather stations (which use the Weather Underground protocol) and publishes it to Home Assistant via MQTT with auto-discovery support.

## Requirements

1. **MQTT Broker**: The Mosquitto broker add-on or an external MQTT broker
2. **DNS Redirect**: Configure your router/DNS to redirect `rtupdate.wunderground.com` to your Home Assistant IP address

## Installation

1. Add this repository to your Home Assistant add-on store
2. Install the "VEVOR Weather Station Bridge (Go)" add-on
3. Configure the add-on options (see Configuration below)
4. Start the add-on

## Configuration

| Option | Description | Default |
|--------|-------------|---------|
| `device_name` | Name shown in Home Assistant | "Weather Station" |
| `device_manufacturer` | Manufacturer name | "VEVOR" |
| `device_model` | Model name | "7-in-1 Weather Station" |
| `units` | Unit system: `metric` or `imperial` | "metric" |
| `mqtt_host` | MQTT broker host (leave empty for auto-detect) | "" |
| `mqtt_port` | MQTT broker port | 1883 |
| `mqtt_user` | MQTT username (leave empty for auto-detect) | "" |
| `mqtt_password` | MQTT password (leave empty for auto-detect) | "" |
| `mqtt_prefix` | MQTT discovery prefix | "homeassistant" |
| `timezone` | Timezone for timestamps | "Europe/Berlin" |
| `wu_forward` | Forward data to Weather Underground | false |
| `wu_username` | Weather Underground station ID | "" |
| `wu_password` | Weather Underground password | "" |
| `log_level` | Logging level: DEBUG, INFO, WARNING, ERROR | "INFO" |

## Sensors

The following sensors are automatically created in Home Assistant:

- **Temperature** - Current temperature
- **Humidity** - Relative humidity
- **Barometric Pressure** - Atmospheric pressure
- **Dew Point** - Dew point temperature
- **Wind Speed** - Current wind speed
- **Wind Gust Speed** - Maximum gust speed
- **Wind Direction** - Wind direction in degrees (with cardinal direction attribute)
- **Rainfall** - Hourly rainfall
- **Daily Rainfall** - Daily accumulated rainfall
- **UV Index** - UV radiation index
- **Solar Radiation** - Solar irradiance

## DNS Setup

Your weather station sends data to `rtupdate.wunderground.com`. You need to redirect this to your Home Assistant IP.

### Using Pi-hole / AdGuard Home

Add a DNS rewrite/local DNS entry:

- Domain: `rtupdate.wunderground.com`
- IP: Your Home Assistant IP address

### Using Router DNS

Most routers support custom DNS entries. Add:

- Hostname: `rtupdate.wunderground.com`
- IP: Your Home Assistant IP

## Weather Underground Forwarding

If you still want your data to appear on Weather Underground while using this add-on:

1. Set `wu_forward: true`
2. Enter your Weather Underground Station ID and Password
3. The add-on will forward data after processing

The add-on uses Google DNS (8.8.8.8) to resolve the real Weather Underground IP, bypassing your local DNS redirect.

## Differences from Python Version

This Go implementation offers:

- **Smaller image size**: ~15 MB vs ~150 MB for Python
- **Lower memory usage**: ~5 MB vs ~50 MB
- **Faster startup**: Nearly instant vs 2-3 seconds
- **Single binary**: No Python runtime or dependencies

Both versions produce identical MQTT messages and are fully compatible.

## Troubleshooting

### No data appearing

1. Check the add-on logs for incoming requests
2. Verify DNS redirect is working: `nslookup rtupdate.wunderground.com`
3. Ensure your weather station is configured to send data to Weather Underground

### MQTT connection issues

1. Verify Mosquitto broker is running
2. Check MQTT credentials if using authentication
3. Try setting explicit `mqtt_host` if auto-detect fails

### Health Check

The add-on exposes a health endpoint at `/health` which returns:

- `200 OK` - MQTT connected and operational
- `503 Service Unavailable` - MQTT disconnected

## Support

Report issues at: <https://github.com/lenucksi/VevorWeatherbridge>
