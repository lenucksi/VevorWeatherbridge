# VEVOR Weather Station Bridge

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Flenucksi%2FVevorWeatherbridge.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Flenucksi%2FVevorWeatherbridge?ref=badge_shield)

[![CI](https://github.com/lenucksi/VevorWeatherbridge/actions/workflows/ci.yml/badge.svg)](https://github.com/lenucksi/VevorWeatherbridge/actions/workflows/ci.yml)

A **Home Assistant Add-on** that intercepts weather data from VEVOR 7-in-1 Wi-Fi weather stations (or any station using Weather Underground format) and forwards it to Home Assistant via **MQTT Discovery**.

## Important

No support or guarantees for function, safety or security of any sorts. 
Expect that this software will kill your dog and eat it. 

It is explicitly forbidden to use it for any purpose that would be, direct or indirectly, be connected to anything that would be related to safety or security of building, entity, machinery, human life, etc. You have been warned.


## Features

- **Home Assistant Add-on**: One-click installation from add-on repository
- **MQTT Auto-Discovery**: Sensors automatically appear in Home Assistant
- **Device Grouping**: All sensors grouped under one weather station device
- **Unit Conversion**: Metric or imperial units
- **16-Point Compass Rose**: Wind direction with cardinal directions (N, NNE, NE, etc.)
- **Weather Underground Forwarding**: Optional - keep uploading to WU while also using locally
- **Multi-Architecture**: amd64, armv7, aarch64, armhf, i386
- **Auto MQTT Detection**: Works automatically with HA's Mosquitto broker

## Quick Start

### Prerequisites

1. **Home Assistant OS** or **Home Assistant Supervised**
2. **MQTT Broker** - [Mosquitto broker add-on](https://github.com/home-assistant/addons/tree/master/mosquitto) recommended
3. **DNS Redirect** - Pi-hole, router DNS override, or similar
4. **Port 80 Access** - Via [Nginx Proxy Manager add-on](#nginx-proxy-manager-setup-recommended) or direct port mapping

### Installation

1. **Add the repository** to Home Assistant:
   - Go to **Settings** → **Add-ons** → **Add-on Store**
   - Click **⋮** (top right) → **Repositories**
   - Add: `https://github.com/lenucksi/VevorWeatherbridge`

2. **Install** "VEVOR Weather Station Bridge"

3. **Configure** the add-on (see [Configuration](#configuration))

4. **Set up port 80** access (see [Nginx Proxy Manager Setup](#nginx-proxy-manager-setup-recommended))

5. **Configure DNS redirect** (see [DNS Setup](#dns-setup))

6. **Start** the add-on

## Configuration

### Basic Options

```yaml
device_name: "Weather Station"
device_manufacturer: "VEVOR"
device_model: "7-in-1 Weather Station"
units: "metric"
timezone: "Europe/Berlin"
log_level: "INFO"
```

### MQTT Options

Leave empty for auto-detection (recommended with Mosquitto add-on):

```yaml
mqtt_host: ""
mqtt_port: 1883
mqtt_user: ""
mqtt_password: ""
mqtt_prefix: "homeassistant"
```

### Weather Underground Forwarding (Optional)

```yaml
wu_forward: true
wu_username: "YOUR_STATION_ID"
wu_password: "YOUR_STATION_KEY"
```

## Network Setup

### Nginx Proxy Manager Setup (Recommended)

The weather station sends data to port 80. Since Home Assistant uses port 80 for other services (emulated Hue), we recommend using [Nginx Proxy Manager](https://github.com/hassio-addons/addon-nginx-proxy-manager) to handle the routing.

#### Step 1: Install Nginx Proxy Manager Add-on

1. Add the community add-ons repository:
   - **Settings** → **Add-ons** → **Add-on Store** → **⋮** → **Repositories**
   - Add: `https://github.com/hassio-addons/repository`

2. Install **Nginx Proxy Manager** from the add-on store

3. Configure the add-on port:
   - Set the admin interface to port 81
   - Set HTTP port to 80
   - Set HTTPS port to 443

4. Start the add-on and access admin at `http://your-ha-ip:81`
   - Default login: `admin@example.com` / `changeme`

#### Step 2: Create Proxy Host for Weather Station

In Nginx Proxy Manager admin interface:

1. Go to **Proxy Hosts** → **Add Proxy Host**

2. Configure the **Details** tab:

   | Field | Value |
   |-------|-------|
   | Domain Names | `rtupdate.wunderground.com` |
   | Scheme | `http` |
   | Forward Hostname/IP | `172.30.32.1` (HA internal IP) or `localhost` |
   | Forward Port | `8099` (the port VevorWeatherbridge listens on) |

3. Enable **Block Common Exploits**

4. Click **Save**

#### Step 3: Update Home Assistant Configuration

Add to your `configuration.yaml`:

```yaml
http:
  use_x_forwarded_for: true
  trusted_proxies:
    - 172.16.0.0/12    # Docker network range
    - 127.0.0.1        # Localhost
```

Restart Home Assistant after this change.

#### Alternative: Direct Port Mapping

If you don't want to use Nginx Proxy Manager, you can map port 80 directly to the add-on. However, this may conflict with other services.

In the add-on configuration:

```yaml
# Network configuration
ports:
  80/tcp: 80  # Direct mapping instead of default 8099
```

> **Note:** This will prevent Home Assistant's emulated Hue from working on port 80.

### DNS Setup

The weather station connects to `rtupdate.wunderground.com`. You must redirect this domain to your Home Assistant IP.

#### Using Pi-hole (Recommended)

1. Go to **Local DNS** → **DNS Records**
2. Add record:
   - **Domain**: `rtupdate.wunderground.com`
   - **IP**: Your Home Assistant IP (e.g., `192.168.1.50`)

#### Using AdGuard Home

```yaml
# In AdGuard Home DNS rewrites
- domain: rtupdate.wunderground.com
  answer: 192.168.1.50
```

#### Using Router DNS

Most routers have a "Local DNS" or "DNS Override" feature:

1. Find DNS settings in your router admin
2. Add: `rtupdate.wunderground.com` → `192.168.1.50`

#### Verify DNS Setup

```bash
# Should return your Home Assistant IP
nslookup rtupdate.wunderground.com
```

### Weather Station Configuration

On your VEVOR weather station:

1. Open the weather station app (WS View or similar)
2. Enable **Weather Underground** uploads
3. Enter any Station ID and Key (these are captured but optional)
4. The DNS redirect will route traffic to your local add-on

## Sensors Created

All sensors appear under a single device in Home Assistant:

| Sensor | Device Class | Unit (Metric) | Unit (Imperial) |
|--------|--------------|---------------|-----------------|
| Temperature | `temperature` | °C | °F |
| Humidity | `humidity` | % | % |
| Dew Point | `temperature` | °C | °F |
| Barometric Pressure | `pressure` | hPa | inHg |
| Wind Speed | `wind_speed` | km/h | mph |
| Wind Gust Speed | `wind_speed` | km/h | mph |
| Wind Direction | - | ° | ° |
| Rainfall (hourly) | `precipitation` | mm | in |
| Daily Rainfall | `precipitation` | mm | in |
| UV Index | - | index | index |
| Solar Radiation | `irradiance` | W/m² | W/m² |

### Wind Direction Cardinal Attribute

The Wind Direction sensor includes a `cardinal` attribute with the 16-point compass direction:

```yaml
# Example sensor attributes
state: 225
cardinal: "SW"
measured_on: "2024-01-15T14:30:00+01:00"
```

## Dashboard Examples

### Windrose Card

For wind visualization, install the [Lovelace Windrose Card](https://github.com/aukedejong/lovelace-windrose-card) via HACS:

```yaml
type: custom:windrose-card
title: Wind Rose
data_period:
  hours_to_show: 24
wind_direction_entity: sensor.weather_station_wind_direction
windspeed_entities:
  - entity: sensor.weather_station_wind_speed
  - entity: sensor.weather_station_wind_gust_speed
refresh_interval: 300
wind_direction_count: 16
center_calm_percentage: true
```

See [vevor-weatherbridge/lovelace-windrose-example.yaml](vevor-weatherbridge/lovelace-windrose-example.yaml) for a complete example.

### Weather Card

```yaml
type: weather-forecast
entity: weather.home  # Your weather entity
show_current: true
show_forecast: false
```

### Sensor Cards

```yaml
type: entities
title: Weather Station
entities:
  - entity: sensor.weather_station_temperature
  - entity: sensor.weather_station_humidity
  - entity: sensor.weather_station_barometric_pressure
  - entity: sensor.weather_station_wind_speed
  - entity: sensor.weather_station_wind_direction
```

## Automation Examples

### High Wind Alert

```yaml
automation:
  - alias: "High Wind Alert"
    trigger:
      - platform: numeric_state
        entity_id: sensor.weather_station_wind_gust_speed
        above: 50
    action:
      - service: notify.mobile_app
        data:
          title: "Weather Alert"
          message: "Wind gusts exceeding 50 km/h!"
```

### Rain Detected

```yaml
automation:
  - alias: "Rain Detected"
    trigger:
      - platform: numeric_state
        entity_id: sensor.weather_station_rainfall
        above: 0
    action:
      - service: cover.close_cover
        target:
          entity_id: cover.patio_awning
```

## Troubleshooting

### No Data Received

1. **Check DNS redirect**:

   ```bash
   nslookup rtupdate.wunderground.com
   # Should return your HA IP
   ```

2. **Check Nginx Proxy Manager** (if using):
   - Verify proxy host is configured correctly
   - Check NPM logs for incoming requests

3. **Check add-on logs**:
   - Settings → Add-ons → VEVOR Weather Station Bridge → Log

4. **Verify port is accessible**:

   ```bash
   curl http://your-ha-ip:8099/weatherstation/updateweatherstation.php?test=1
   # Should return "success"
   ```

### MQTT Connection Failed

1. Ensure Mosquitto broker add-on is running
2. Check MQTT credentials if using external broker
3. Restart both Mosquitto and this add-on

### Sensors Not Appearing

1. Verify MQTT integration is configured
2. Check MQTT for discovery messages:
   - Developer Tools → MQTT → Listen to `homeassistant/sensor/#`
3. Wait 1-2 minutes for auto-discovery
4. Restart Home Assistant if needed

## Architecture

```text
┌─────────────────────┐
│   Weather Station   │
│    (VEVOR 7-in-1)   │
└──────────┬──────────┘
           │ HTTP POST to rtupdate.wunderground.com
           ▼
┌─────────────────────┐
│   DNS Redirect      │
│   (Pi-hole/Router)  │
└──────────┬──────────┘
           │ Redirected to Home Assistant IP
           ▼
┌─────────────────────────────────────────────────────┐
│                 Home Assistant                       │
│  ┌───────────────────────────────────────────────┐  │
│  │        Nginx Proxy Manager (port 80)          │  │
│  └───────────────────┬───────────────────────────┘  │
│                      │ Proxy to port 8099            │
│  ┌───────────────────▼───────────────────────────┐  │
│  │     VevorWeatherbridge Add-on (port 8099)     │  │
│  └───────────────────┬───────────────────────────┘  │
│                      │ MQTT Publish                  │
│  ┌───────────────────▼───────────────────────────┐  │
│  │          Mosquitto Broker Add-on              │  │
│  └───────────────────┬───────────────────────────┘  │
│                      │ MQTT Discovery                │
│  ┌───────────────────▼───────────────────────────┐  │
│  │          Home Assistant Core                  │  │
│  │          (Sensors & Entities)                 │  │
│  └───────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────┘
           │
           │ Optional: WU Forward
           ▼
┌─────────────────────┐
│  Weather Underground│
│  (rtupdate.wunder..│
│   ground.com)       │
└─────────────────────┘
```

## Related Projects & Links

### Home Assistant Add-ons

- [Mosquitto Broker](https://github.com/home-assistant/addons/tree/master/mosquitto) - MQTT broker for Home Assistant
- [Nginx Proxy Manager](https://github.com/hassio-addons/addon-nginx-proxy-manager) - Easy reverse proxy management

### Visualization

- [Lovelace Windrose Card](https://github.com/aukedejong/lovelace-windrose-card) - Wind direction visualization

### Documentation

- [Home Assistant MQTT Discovery](https://www.home-assistant.io/integrations/mqtt/#mqtt-discovery)
- [Nginx Proxy Manager Guide](https://nginxproxymanager.com/guide/)
- [Add-on Documentation](vevor-weatherbridge/DOCS.md)

## Development

### Project Structure


### Running Tests


### Code Quality


## License

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Flenucksi%2FVevorWeatherbridge.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Flenucksi%2FVevorWeatherbridge?ref=badge_large)

## Acknowledgements

- Original Weather Underground relay script inspiration by [@vlovmx](https://github.com/vlovmx)
- Python rewrite and containerization by C9H13NO3-dev
- Home Assistant Add-on framework by the Home Assistant team
- Community contributions and feedback

---

**Questions or issues?** [Open an issue](https://github.com/lenucksi/VevorWeatherbridge/issues)
