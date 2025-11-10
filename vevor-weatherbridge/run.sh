#!/usr/bin/with-contenv bashio
# ==============================================================================
# Home Assistant Add-on: VEVOR Weather Station Bridge
# Starts the weather station bridge service
# ==============================================================================

bashio::log.info "Starting VEVOR Weather Station Bridge..."

# Read configuration from Home Assistant addon options
CONFIG_PATH=/data/options.json

# Read device configuration
export DEVICE_NAME=$(bashio::config 'device_name')
export DEVICE_MANUFACTURER=$(bashio::config 'device_manufacturer')
export DEVICE_MODEL=$(bashio::config 'device_model')
export UNITS=$(bashio::config 'units')
export MQTT_PREFIX=$(bashio::config 'mqtt_prefix')
export TZ=$(bashio::config 'timezone')

# Generate device ID from device name (lowercase, replace spaces with underscores)
export DEVICE_ID=$(echo "${DEVICE_NAME}" | tr '[:upper:]' '[:lower:]' | tr ' ' '_')

# Weather Underground forwarding (optional)
export WU_FORWARD=$(bashio::config 'wu_forward')
if bashio::config.true 'wu_forward'; then
    export WU_USERNAME=$(bashio::config 'wu_username')
    export WU_PASSWORD=$(bashio::config 'wu_password')
    bashio::log.info "Weather Underground forwarding enabled"
fi

# MQTT Configuration
# Try to use Home Assistant's internal MQTT broker first
if bashio::services.available "mqtt"; then
    bashio::log.info "Using Home Assistant's internal MQTT broker"
    export MQTT_HOST=$(bashio::services "mqtt" "host")
    export MQTT_PORT=$(bashio::services "mqtt" "port")
    export MQTT_USER=$(bashio::services "mqtt" "username")
    export MQTT_PASSWORD=$(bashio::services "mqtt" "password")
else
    # Use user-provided MQTT configuration
    MQTT_HOST_CONFIG=$(bashio::config 'mqtt_host')
    if bashio::var.has_value "${MQTT_HOST_CONFIG}"; then
        bashio::log.info "Using external MQTT broker"
        export MQTT_HOST="${MQTT_HOST_CONFIG}"
        export MQTT_PORT=$(bashio::config 'mqtt_port')
        export MQTT_USER=$(bashio::config 'mqtt_user')
        export MQTT_PASSWORD=$(bashio::config 'mqtt_password')
    else
        bashio::log.fatal "No MQTT broker configured! Please configure either Home Assistant's MQTT integration or provide external MQTT settings."
        exit 1
    fi
fi

bashio::log.info "Device: ${DEVICE_NAME} (${DEVICE_ID})"
bashio::log.info "MQTT Broker: ${MQTT_HOST}:${MQTT_PORT}"
bashio::log.info "Units: ${UNITS}"
bashio::log.info "Timezone: ${TZ}"

# Start the weather station bridge
bashio::log.info "Starting weather station service on port 80..."
exec python3 /app/weatherstation.py
