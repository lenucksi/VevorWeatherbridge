import json
import logging
import os
import sys
from datetime import datetime

import dns.resolver
import paho.mqtt.client as mqtt
import pytz
import requests
from flask import Flask, request

# Configure logging for Home Assistant addon environment
# Allow LOG_LEVEL to be configured via environment variable
LOG_LEVEL = os.environ.get("LOG_LEVEL", "INFO").upper()
logging.basicConfig(
    level=getattr(logging, LOG_LEVEL, logging.INFO),
    format="%(asctime)s - %(name)s - %(levelname)s - %(message)s",
    stream=sys.stdout,
)
logger = logging.getLogger(__name__)

# MQTT settings
MQTT_HOST = os.environ.get("MQTT_HOST", "localhost")
MQTT_PORT = int(os.environ.get("MQTT_PORT", 1883))
MQTT_USER = os.environ.get("MQTT_USER")
MQTT_PASSWORD = os.environ.get("MQTT_PASSWORD")
MQTT_PREFIX = os.environ.get("MQTT_PREFIX", "homeassistant")
DEVICE_ID = os.environ.get("DEVICE_ID", "weather_station")
DEVICE_NAME = os.environ.get("DEVICE_NAME", "Weather Station")
DEVICE_MANUFACTURER = os.environ.get("DEVICE_MANUFACTURER", "VEVOR")
DEVICE_MODEL = os.environ.get("DEVICE_MODEL", "7-in-1 Weather Station")
TIMEZONE = os.environ.get("TZ", "Europe/Berlin")
UNITS = os.environ.get("UNITS", "metric").lower()
WU_FORWARD = os.environ.get("WU_FORWARD", "false").lower() == "true"
WU_USERNAME = os.environ.get("WU_USERNAME")
WU_PASSWORD = os.environ.get("WU_PASSWORD")

logger.info("Starting VEVOR Weather Station Bridge")
logger.info(f"Log level: {LOG_LEVEL}")
logger.info(f"Device: {DEVICE_NAME} ({DEVICE_ID})")
logger.info(f"MQTT: {MQTT_HOST}:{MQTT_PORT}")
logger.info(f"Units: {UNITS}")

app = Flask(__name__)

# Global MQTT connection state
mqtt_connected = False


def on_connect(client, userdata, flags, reason_code, properties):
    """MQTT connection callback (API v2)."""
    global mqtt_connected
    # reason_code is a ReasonCode instance in v2
    if reason_code == 0:
        mqtt_connected = True
        logger.info(f"Connected to MQTT broker successfully (reason_code={reason_code})")
    else:
        mqtt_connected = False
        logger.error(f"Failed to connect to MQTT broker (reason_code={reason_code})")


def on_disconnect(client, userdata, flags, reason_code, properties):
    """MQTT disconnection callback (API v2)."""
    global mqtt_connected
    mqtt_connected = False
    if reason_code != 0:
        logger.warning(f"Unexpected MQTT disconnection (reason_code={reason_code})")


def on_publish(client, userdata, mid, reason_codes, properties):
    """MQTT publish callback (API v2)."""
    logger.debug(f"MQTT message published (mid={mid})")


# Create MQTT client with callback API version 2
mqtt_client = mqtt.Client(callback_api_version=mqtt.CallbackAPIVersion.VERSION2)
mqtt_client.on_connect = on_connect
mqtt_client.on_disconnect = on_disconnect
mqtt_client.on_publish = on_publish

if MQTT_USER:
    logger.debug(f"Setting MQTT credentials for user: {MQTT_USER}")
    mqtt_client.username_pw_set(MQTT_USER, MQTT_PASSWORD)

try:
    logger.info(f"Connecting to MQTT broker at {MQTT_HOST}:{MQTT_PORT}...")
    mqtt_client.connect(MQTT_HOST, MQTT_PORT, 60)
    mqtt_client.loop_start()
    logger.info("MQTT connection initiated")
except Exception as e:
    logger.error(f"Failed to initiate MQTT connection: {e}")
    logger.error("The addon will continue but data will not be published until MQTT connection is established")


def f_to_c(f):
    return round((float(f) - 32) * 5.0 / 9.0, 1)


def inhg_to_hpa(inhg):
    return round(float(inhg) * 33.8639, 1)


def mph_to_kmh(mph):
    return round(float(mph) * 1.60934, 1)


def inch_to_mm(inch):
    return round(float(inch) * 25.4, 1)


def safe_get(key):
    return request.args.get(key, None)


@app.route("/weatherstation/updateweatherstation.php")
def update():
    attributes = {
        "Barometric Pressure": {
            "value": (inhg_to_hpa(safe_get("baromin")) if UNITS == "metric" else round(float(safe_get("baromin")), 1))
            if safe_get("baromin")
            else None,
            "unit": "hPa" if UNITS == "metric" else "inHg",
            "device_class": "atmospheric_pressure",
        },
        "Temperature": {
            "value": (f_to_c(safe_get("tempf")) if UNITS == "metric" else round(float(safe_get("tempf")), 1))
            if safe_get("tempf")
            else None,
            "unit": "°C" if UNITS == "metric" else "°F",
            "device_class": "temperature",
        },
        "Humidity": {"value": safe_get("humidity"), "unit": "%", "device_class": "humidity"},
        "Dew Point": {
            "value": (f_to_c(safe_get("dewptf")) if UNITS == "metric" else round(float(safe_get("dewptf")), 1))
            if safe_get("dewptf")
            else None,
            "unit": "°C" if UNITS == "metric" else "°F",
            "device_class": "temperature",
        },
        "Rainfall": {
            "value": (inch_to_mm(safe_get("rainin")) if UNITS == "metric" else round(float(safe_get("rainin")), 2))
            if safe_get("rainin")
            else None,
            "unit": "mm" if UNITS == "metric" else "in",
            "device_class": "precipitation",
        },
        "Daily Rainfall": {
            "value": (
                inch_to_mm(safe_get("dailyrainin")) if UNITS == "metric" else round(float(safe_get("dailyrainin")), 2)
            )
            if safe_get("dailyrainin")
            else None,
            "unit": "mm" if UNITS == "metric" else "in",
            "device_class": "precipitation",
        },
        "Wind Direction": {"value": safe_get("winddir"), "unit": "°", "device_class": None},
        "Wind Speed": {
            "value": (
                mph_to_kmh(safe_get("windspeedmph")) if UNITS == "metric" else round(float(safe_get("windspeedmph")), 1)
            )
            if safe_get("windspeedmph")
            else None,
            "unit": "km/h" if UNITS == "metric" else "mph",
            "device_class": "wind_speed",
        },
        "Wind Gust Speed": {
            "value": (
                mph_to_kmh(safe_get("windgustmph")) if UNITS == "metric" else round(float(safe_get("windgustmph")), 1)
            )
            if safe_get("windgustmph")
            else None,
            "unit": "km/h" if UNITS == "metric" else "mph",
            "device_class": "wind_speed",
        },
        "UV Index": {"value": safe_get("UV"), "unit": "index", "device_class": None},
        "Solar Radiation": {"value": safe_get("solarRadiation"), "unit": "W/m²", "device_class": "irradiance"},
    }

    dateutc = safe_get("dateutc")
    local_time = ""
    if dateutc:
        try:
            dt = datetime.strptime(dateutc, "%Y-%m-%d %H:%M:%S")
            dt = pytz.utc.localize(dt).astimezone(pytz.timezone(TIMEZONE))
            local_time = dt.strftime("%Y-%m-%d %H:%M:%S")
        except Exception:
            local_time = dateutc  # fallback

    # Check MQTT connection status
    if not mqtt_connected:
        logger.warning("MQTT not connected - attempting to reconnect...")
        try:
            mqtt_client.reconnect()
        except Exception as e:
            logger.error(f"MQTT reconnection failed: {e}")

    logger.debug(f"Processing weather data - MQTT connected: {mqtt_connected}")
    logger.debug(f"Received parameters: {request.args.to_dict()}")

    # Publish each sensor to MQTT using HA auto-discovery
    published_count = 0
    for name, data in attributes.items():
        if data["value"] is None:
            logger.debug(f"Skipping {name} - no value")
            continue

        sensor_id = name.lower().replace(" ", "_")
        base = f"{MQTT_PREFIX}/sensor/{DEVICE_ID}_{sensor_id}"
        state_topic = f"{base}/state"
        attr_topic = f"{base}/attributes"
        config_topic = f"{base}/config"

        config_payload = {
            "name": f"{DEVICE_NAME} {name}",
            "state_topic": state_topic,
            "unit_of_measurement": data["unit"],
            "device_class": data["device_class"],
            "unique_id": f"{DEVICE_ID}_{sensor_id}",
            "json_attributes_topic": attr_topic,
            "device": {
                "identifiers": [DEVICE_ID],
                "name": DEVICE_NAME,
                "manufacturer": DEVICE_MANUFACTURER,
                "model": DEVICE_MODEL,
            },
        }

        logger.debug(f"Publishing {name}: {data['value']} {data['unit']} to {state_topic}")

        try:
            mqtt_client.publish(config_topic, json.dumps(config_payload), retain=True)
            result_state = mqtt_client.publish(state_topic, str(data["value"]), retain=True)
            mqtt_client.publish(attr_topic, json.dumps({"measured_on": local_time}), retain=True)

            if result_state.rc == mqtt.MQTT_ERR_SUCCESS:
                published_count += 1
                logger.debug(f"Successfully queued {name} for publishing (mid={result_state.mid})")
            else:
                logger.error(f"Failed to queue {name} for publishing (rc={result_state.rc})")
        except Exception as e:
            logger.error(f"Exception publishing {name}: {e}")

    if WU_FORWARD:
        params = request.args.to_dict()
        if WU_USERNAME:
            params["ID"] = WU_USERNAME
        if WU_PASSWORD:
            params["PASSWORD"] = WU_PASSWORD
        try:
            logger.debug("Forwarding to Weather Underground...")
            resolver = dns.resolver.Resolver()
            resolver.nameservers = ["8.8.8.8", "8.8.4.4"]
            wu_ip = resolver.resolve("rtupdate.wunderground.com")[0].to_text()
            url = f"http://{wu_ip}/weatherstation/updateweatherstation.php"
            headers = {"Host": "rtupdate.wunderground.com"}
            response = requests.get(url, params=params, headers=headers, timeout=5)
            logger.debug(f"Weather Underground response: {response.status_code}")
        except Exception as e:
            logger.warning(f"Failed to forward to Weather Underground: {e}")

    if mqtt_connected and published_count > 0:
        logger.info(f"Weather data received - published {published_count} sensors to MQTT")
    elif published_count > 0:
        logger.warning(f"Weather data queued ({published_count} sensors) but MQTT not confirmed connected")
    else:
        logger.error("Weather data received but no sensors were published")

    return "success", 200


if __name__ == "__main__":
    logger.info("Starting Flask server on 0.0.0.0:80")
    app.run(host="0.0.0.0", port=80)  # nosec B104 - binding to all interfaces is intentional for container
