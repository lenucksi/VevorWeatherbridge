# VEVOR Weather Station Bridge

This project provides a **Home Assistant Add-on** (and standalone Docker container) for ingesting weather data from a VEVOR 7-in-1 Wi-Fi Solar Self-Charging Weather Station (Model YT60234, or any station sending data in Weather Underground format) and forwarding it to Home Assistant via **MQTT**.

---

## Features

- **Home Assistant Add-on**: One-click installation from the add-on store
- Accepts Weather Underground (WU) station GET requests (as sent by the VEVOR weather station)
- Converts measurements to **metric** or **imperial** units
- Publishes sensor data to Home Assistant via MQTT using the auto-discovery format
- All sensors appear under one device in Home Assistant
- Automatic MQTT broker detection (works with HA's internal Mosquitto broker)
- Optional Weather Underground forwarding
- Multi-architecture support (amd64, armv7, aarch64, armhf, i386)
- Responds with `success` so the weather station doesn't retry

---

## Installation

### Option 1: Home Assistant Add-on (Recommended)

1. **Add this repository to your Home Assistant**:
   - Go to **Settings** → **Add-ons** → **Add-on Store**
   - Click the **⋮** menu (top right) → **Repositories**
   - Add: `https://github.com/C9H13NO3-dev/VevorWeatherbridge`
   - Click **Add**

2. **Install the add-on**:
   - Find "VEVOR Weather Station Bridge" in the add-on store
   - Click on it and press **Install**

3. **Configure the add-on**:
   - Go to the **Configuration** tab
   - Set your device name, units (metric/imperial), timezone
   - If you have an external MQTT broker, configure it (otherwise it auto-detects HA's internal broker)
   - Optionally enable Weather Underground forwarding

4. **Start the add-on**:
   - Click **Start**
   - Check the **Log** tab to ensure it's running

5. **Configure DNS redirect** (see [DNS Setup](#dns-setup) below)

For detailed add-on documentation, see [DOCS.md](DOCS.md).

### Option 2: Standalone Docker Container

If you prefer to run this outside of Home Assistant:

#### 1. Clone the repository

```bash
git clone https://github.com/C9H13NO3-dev/VevorWeatherbridge.git
cd VevorWeatherbridge
```

#### 2. Configure the environment

Edit the `docker-compose.yml` file and set the following variables:

- `MQTT_HOST`: Hostname or IP of your MQTT broker
- `MQTT_PORT`: Broker port (default `1883`)
- `MQTT_USER` / `MQTT_PASSWORD`: Credentials if required
- `DEVICE_ID`: Unique identifier for the weather station device (default `weather_station`)
- `DEVICE_NAME`: Display name for the device in Home Assistant (default `Weather Station`)
- `DEVICE_MANUFACTURER`: (optional) Manufacturer name shown in Home Assistant (default `VEVOR`)
- `DEVICE_MODEL`: (optional) Model name (default `7-in-1 Weather Station`)
- `UNITS`: `metric` (default) or `imperial`
- `WU_FORWARD`: Set to `true` to also forward data to Weather Underground (default `false`)
- `WU_USERNAME` / `WU_PASSWORD`: Credentials for Weather Underground (optional)

Example:

```yaml
environment:
  TZ: Europe/Berlin
  MQTT_HOST: 192.168.1.100
  MQTT_PORT: 1883
  MQTT_USER: youruser
  MQTT_PASSWORD: yourpass
  DEVICE_ID: weather_station
  DEVICE_NAME: "Backyard Weather"
  DEVICE_MANUFACTURER: VEVOR
  DEVICE_MODEL: "7-in-1 Weather Station"
  # optional: "metric" (default) or "imperial"
  UNITS: metric
  # forward data to Weather Underground
  WU_FORWARD: "false"
  # credentials if forwarding
  WU_USERNAME: yourWUuser
  WU_PASSWORD: yourWUpass
```

#### 3. Build and run

```bash
docker-compose up --build -d
```

The service now listens on port `80` for requests to `/weatherstation/updateweatherstation.php`.

---

## DNS Setup

**Critical Step**: Your weather station needs to connect to this service instead of Weather Underground. You must configure your network to redirect Weather Underground traffic to your Home Assistant (or Docker host) IP address.

The weather station connects to: `rtupdate.wunderground.com`

### Using Pi-hole (Recommended)

1. Log into your Pi-hole admin interface
2. Go to **Local DNS** → **DNS Records**
3. Add a new record:
   - **Domain**: `rtupdate.wunderground.com`
   - **IP Address**: Your Home Assistant IP (e.g., `192.168.1.50`)
4. Save the record

### Using Router DNS Override

Many routers support local DNS overrides:

1. Log into your router admin interface
2. Look for **DNS Settings**, **Local DNS**, or **Host Name Mapping**
3. Add an entry mapping `rtupdate.wunderground.com` to your Home Assistant IP
4. Save and restart router if needed

**Note:** If your router has DNS rebind protection enabled, you must allow this domain in your router settings when overriding DNS with Pi-hole.

### Weather Station Configuration

On the weather station itself, enable Weather Underground uploads with any station ID/key. The DNS redirect will automatically route this traffic to your local service.

---

## Endpoints

The service listens for GET requests at:

```text
/weatherstation/updateweatherstation.php
```

With query parameters matching the WU format, e.g.:

```text
http://<your-server-ip>/weatherstation/updateweatherstation.php?ID=XXXXX&PASSWORD=XXXXX&dateutc=xxxx-x-xx+xx:xx:xx&baromin=x&tempf=x&humidity=x&dewptf=x&rainin=x&dailyrainin=x&winddir=x&windspeedmph=x&windgustmph=x&UV=x&solarRadiation=x
```

---

## How it Works

1. The weather station uploads data to this endpoint.
2. The service converts units to metric or keeps imperial values depending on `UNITS`.
3. Each value is published to Home Assistant via MQTT, auto-discovered as a sensor, and grouped under the configured device.
4. The endpoint returns `success` to acknowledge the update.

### Home Assistant Sensor Entities

The following sensors are created or updated and will appear under the device specified by `DEVICE_NAME`:

*Units in parentheses assume `UNITS=metric`; values switch to imperial when `UNITS=imperial`.*

- `sensor.weather_station_barometric_pressure` (hPa)
- `sensor.weather_station_temperature` (°C)
- `sensor.weather_station_humidity` (%)
- `sensor.weather_station_dew_point` (°C)
- `sensor.weather_station_rainfall` (mm)
- `sensor.weather_station_daily_rainfall` (mm)
- `sensor.weather_station_wind_direction` (°)
- `sensor.weather_station_wind_speed` (km/h)
- `sensor.weather_station_wind_gust_speed` (km/h)
- `sensor.weather_station_uv_index` (index)
- `sensor.weather_station_solar_radiation` (W/m²)

You can use these entities directly in your Home Assistant dashboards or automations.

---

## Troubleshooting

- Ensure the container is reachable on port 80 from your network.
- Check that your DNS or Pi-hole setup correctly redirects the WU domain.
- Verify your MQTT broker settings.
- Review logs with:

```bash
docker-compose logs -f
```

Add logging or print statements to `weatherstation.py` for further debugging.

---

## License

This project is licensed under the [CC0 1.0 Universal](LICENSE).

---

## Acknowledgements

- Original Weather Underground relay script inspiration by [@vlovmx](https://github.com/vlovmx)
- Python rewrite and containerization by C9H13NO3-dev
