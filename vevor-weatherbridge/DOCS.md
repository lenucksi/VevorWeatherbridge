# VEVOR Weather Station Bridge - Home Assistant Add-on

## About

This add-on intercepts data from VEVOR weather stations and forwards it to Home Assistant via MQTT. It acts as a local server that your weather station connects to instead of Weather Underground, allowing you to keep your weather data within your home network while still optionally forwarding to Weather Underground if desired.

## Features

- **Automatic MQTT Discovery**: All sensors are automatically discovered by Home Assistant
- **Device Grouping**: All sensors are grouped under a single weather station device
- **Unit Conversion**: Support for both metric and imperial units
- **Weather Underground Forwarding**: Optionally forward data to Weather Underground
- **Flexible MQTT**: Works with Home Assistant's internal MQTT broker or external MQTT servers
- **Multi-Architecture Support**: Runs on various platforms (amd64, armv7, aarch64, armhf, i386)

## Installation

1. Add this repository to your Home Assistant add-on store
2. Install the "VEVOR Weather Station Bridge" add-on
3. Configure the add-on (see Configuration section below)
4. Start the add-on
5. Configure DNS redirect (see Setup Instructions below)

## Configuration

### Basic Configuration

```yaml
device_name: "Weather Station"
device_manufacturer: "VEVOR"
device_model: "7-in-1 Weather Station"
units: "metric"
mqtt_prefix: "homeassistant"
timezone: "Europe/Berlin"
```

**Options:**

- `device_name` (string, required): Name of your weather station as it will appear in Home Assistant
- `device_manufacturer` (string, required): Manufacturer name (default: VEVOR)
- `device_model` (string, required): Model name (default: 7-in-1 Weather Station)
- `units` (string, required): Unit system - `metric` or `imperial`
- `mqtt_prefix` (string, required): MQTT topic prefix (usually `homeassistant`)
- `timezone` (string, required): Your timezone (e.g., `Europe/Berlin`, `America/New_York`)

### MQTT Configuration

If you have Home Assistant's Mosquitto broker add-on installed and configured, the add-on will automatically detect and use it. **No additional MQTT configuration needed!**

If you want to use an external MQTT broker:

```yaml
mqtt_host: "192.168.1.100"
mqtt_port: 1883
mqtt_user: "mqttuser"
mqtt_password: "mqttpassword"
```

**Options:**

- `mqtt_host` (string, optional): MQTT broker hostname/IP (leave empty to use HA's internal broker)
- `mqtt_port` (integer, optional): MQTT broker port (default: 1883)
- `mqtt_user` (string, optional): MQTT username
- `mqtt_password` (string, optional): MQTT password

### Weather Underground Forwarding (Optional)

If you want to continue uploading data to Weather Underground:

```yaml
wu_forward: true
wu_username: "YOUR_WU_STATION_ID"
wu_password: "YOUR_WU_PASSWORD"
```

**Options:**

- `wu_forward` (boolean, optional): Enable Weather Underground forwarding (default: false)
- `wu_username` (string, optional): Your Weather Underground station ID
- `wu_password` (string, optional): Your Weather Underground station password

## Setup Instructions

### Step 1: Install and Configure Add-on

1. Install the add-on from the add-on store
2. Configure the options according to your preferences
3. Start the add-on
4. Check the logs to ensure it's running correctly

### Step 2: Configure DNS Redirect

**Critical Step**: Your weather station needs to connect to this add-on instead of Weather Underground. You must configure your network to redirect Weather Underground traffic to your Home Assistant instance.

The weather station connects to: `rtupdate.wunderground.com`

You need to redirect this domain to your Home Assistant IP address.

#### Option A: Using Pi-hole (Recommended)

1. Log into your Pi-hole admin interface
2. Go to **Local DNS** → **DNS Records**
3. Add a new record:
   - **Domain**: `rtupdate.wunderground.com`
   - **IP Address**: Your Home Assistant IP (e.g., `192.168.1.50`)
4. Save the record

#### Option B: Using Router DNS Override

Many routers support local DNS overrides:

1. Log into your router admin interface
2. Look for **DNS Settings**, **Local DNS**, or **Host Name Mapping**
3. Add an entry:
   - **Hostname**: `rtupdate.wunderground.com`
   - **IP Address**: Your Home Assistant IP
4. Save settings and restart router if needed

#### Option C: Using `/etc/hosts` (For Testing Only)

For quick testing, you can modify the hosts file on a computer:

```text
# Linux/Mac: /etc/hosts
# Windows: C:\Windows\System32\drivers\etc\hosts

192.168.1.50  rtupdate.wunderground.com
```

**Note**: This only works for the specific device and won't redirect your weather station.

### Step 3: Configure Weather Station

Your weather station should already be configured to send data to Weather Underground. The DNS redirect will automatically route this traffic to your Home Assistant add-on.

1. Ensure your weather station is configured with Weather Underground settings
2. The station should automatically start sending data to this add-on after DNS redirect
3. Check the add-on logs to see incoming data

### Step 4: Verify in Home Assistant

1. Go to **Settings** → **Devices & Services**
2. Look for **MQTT** integration
3. You should see your weather station device with all sensors
4. All sensors will be automatically discovered and grouped

## Port Configuration

By default, the add-on uses port 8099 on your host to receive weather station data (which maps to port 80 inside the container). If you need to change this:

1. Go to add-on **Configuration** tab
2. Modify the port mapping under **Network**
3. Update your DNS redirect if you changed from port 80 internally

## Sensors Provided

The add-on creates the following sensors in Home Assistant:

- **Temperature** (°C or °F)
- **Humidity** (%)
- **Barometric Pressure** (hPa or inHg)
- **Dew Point** (°C or °F)
- **Rainfall** (mm or in)
- **Daily Rainfall** (mm or in)
- **Wind Direction** (°)
- **Wind Speed** (km/h or mph)
- **Wind Gust Speed** (km/h or mph)
- **UV Index**
- **Solar Radiation** (W/m²)

All sensors include a `measured_on` attribute with the local timestamp.

The **Wind Direction** sensor also includes a `cardinal` attribute with the compass direction (e.g., "N", "NNE", "NE").

## Windrose Visualization

The Wind Direction sensor is optimized for use with windrose cards that display wind patterns visually. A windrose shows the frequency and intensity of wind from different directions.

### Recommended: Lovelace Windrose Card

The community-standard solution is the **[Lovelace Windrose Card](https://github.com/aukedejong/lovelace-windrose-card)** - an actively maintained custom card that creates beautiful windrose visualizations.

**Features:**

- Historical wind direction frequency display
- Color-coded wind speed ranges (Beaufort scale or custom)
- Current wind direction indicator (red arrow)
- Support for both wind speed and wind gust sensors
- Highly customizable appearance
- Multiple unit options (m/s, km/h, mph, knots)

### Installation

1. **Install via HACS** (Home Assistant Community Store):
   - Open HACS → Frontend
   - Click the three dots menu → Custom repositories
   - Add URL: `https://github.com/aukedejong/lovelace-windrose-card`
   - Category: Lovelace
   - Click "Add"
   - Install "Windrose Card"

2. **Restart Home Assistant** to load the card

3. **Add to Dashboard**:
   - Edit your dashboard
   - Add card → Search for "Custom: Windrose Card"
   - Configure using the UI or YAML (see example below)

### Example Configuration

Add this to your Lovelace dashboard (see `lovelace-windrose-example.yaml` for full example):

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
cardinal_direction_letters: NESW
matching_strategy: direction-first
windspeed_bar_location: bottom
wind_direction_count: 16
center_calm_percentage: true
```

**Configuration Options:**

- `data_period.hours_to_show`: How many hours of historical data to display (default: 24)
- `wind_direction_entity`: Your wind direction sensor (in degrees)
- `windspeed_entities`: One or more wind speed sensors
- `refresh_interval`: Update frequency in seconds (default: 300)
- `wind_direction_count`: Number of compass segments (4-32, default: 16)
- `cardinal_direction_letters`: Compass labels (e.g., "NESW" or custom)
- `center_calm_percentage`: Show calm wind percentage in center
- `windspeed_bar_location`: Position bars at "bottom" or "right"

### Speed Ranges and Colors

You can customize speed ranges and colors to match your preferences:

```yaml
speed_range_beaufort: true  # Use Beaufort scale (default)
# OR define custom ranges:
speed_ranges:
  - from_value: 0
    to_value: 5
    color: '#d4d4d4'
  - from_value: 5
    to_value: 10
    color: '#a3d4d4'
  - from_value: 10
    to_value: 20
    color: '#4fa8d4'
```

### Compatible Sensors

This add-on publishes sensors that are fully compatible with the windrose card:

- ✅ **Wind Direction**: Degrees (0-359°) with compass icon
- ✅ **Wind Speed**: Metric (km/h) or Imperial (mph)
- ✅ **Wind Gust Speed**: For maximum wind display
- ✅ **Cardinal Direction**: Available as sensor attribute

See the included `lovelace-windrose-example.yaml` file for a complete ready-to-use configuration.

## Troubleshooting

### No data received

1. Check add-on logs for errors
2. Verify DNS redirect is working:

   ```bash
   nslookup rtupdate.wunderground.com
   # Should return your Home Assistant IP
   ```

3. Ensure weather station is powered on and connected to WiFi
4. Verify firewall allows traffic on port 8099 (or your configured port)

### MQTT connection failed

1. If using HA's internal broker:
   - Ensure Mosquitto broker add-on is installed and running
   - Restart both the Mosquitto broker and this add-on

2. If using external broker:
   - Verify MQTT credentials are correct
   - Check that the MQTT broker is reachable from Home Assistant
   - Test connection with an MQTT client

### Sensors not appearing in Home Assistant

1. Ensure MQTT integration is configured in Home Assistant
2. Check MQTT integration for discovery messages
3. Wait a few minutes for auto-discovery to complete
4. Restart Home Assistant if needed

### Data not forwarding to Weather Underground

1. Verify `wu_forward: true` in configuration
2. Check that `wu_username` and `wu_password` are set correctly
3. Check add-on logs for error messages
4. Ensure your Home Assistant has internet connectivity

## Data Privacy

This add-on runs entirely within your local network. Weather data is only sent to:

1. **Your MQTT broker** (local or external as configured)
2. **Weather Underground** (only if you enable `wu_forward`)

No data is sent to any other third parties.

## Support

For issues, feature requests, or questions:

- GitHub Issues: <https://github.com/C9H13NO3-dev/VevorWeatherbridge/issues>
- Home Assistant Community: <https://community.home-assistant.io/>

## Credits

Developed by the VevorWeatherbridge team. Licensed under the MIT License.
