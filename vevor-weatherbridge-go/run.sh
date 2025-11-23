#!/usr/bin/with-contenv bashio
# ==============================================================================
# Home Assistant Add-on: VEVOR Weather Station Bridge (Go)
# Starts the weather station bridge service
# ==============================================================================

set -e

bashio::log.info "Starting VEVOR Weather Station Bridge (Go)..."

# Read device configuration using bashio
export DEVICE_NAME=$(bashio::config 'device_name')
export DEVICE_MANUFACTURER=$(bashio::config 'device_manufacturer')
export DEVICE_MODEL=$(bashio::config 'device_model')
export UNITS=$(bashio::config 'units')
export MQTT_PREFIX=$(bashio::config 'mqtt_prefix')
export TZ=$(bashio::config 'timezone')
export LOG_LEVEL=$(bashio::config 'log_level')

# Generate device ID from device name (lowercase, replace spaces with underscores)
export DEVICE_ID=$(echo "${DEVICE_NAME}" | tr '[:upper:]' '[:lower:]' | tr ' ' '_')

# Weather Underground forwarding (optional)
if bashio::config.true 'wu_forward'; then
    export WU_FORWARD="true"
    export WU_USERNAME=$(bashio::config 'wu_username')
    export WU_PASSWORD=$(bashio::config 'wu_password')
    bashio::log.info "Weather Underground forwarding enabled"
else
    export WU_FORWARD="false"
fi

# MQTT Configuration
# Try to use Home Assistant's internal MQTT broker first
if ! bashio::config.has_value 'mqtt_host'; then
    # Use bashio to get MQTT service credentials
    bashio::log.info "Attempting to auto-detect MQTT broker via Supervisor..."

    if bashio::services.available "mqtt"; then
        export MQTT_HOST=$(bashio::services "mqtt" "host")
        export MQTT_PORT=$(bashio::services "mqtt" "port")
        export MQTT_USER=$(bashio::services "mqtt" "username")
        export MQTT_PASSWORD=$(bashio::services "mqtt" "password")

        bashio::log.info "Using Home Assistant MQTT broker from Supervisor"
        bashio::log.info "MQTT Host: ${MQTT_HOST}:${MQTT_PORT}"
        bashio::log.info "MQTT User: ${MQTT_USER}"
    else
        bashio::log.fatal "No MQTT service available!"
        bashio::log.fatal "Please install Mosquitto broker addon or configure MQTT manually."
        exit 1
    fi
else
    # Use user-provided MQTT configuration
    bashio::log.info "Using external MQTT broker from configuration"
    export MQTT_HOST=$(bashio::config 'mqtt_host')
    export MQTT_PORT=$(bashio::config 'mqtt_port')
    export MQTT_USER=$(bashio::config 'mqtt_user')
    export MQTT_PASSWORD=$(bashio::config 'mqtt_password')
fi

bashio::log.info "Device: ${DEVICE_NAME} (${DEVICE_ID})"
bashio::log.info "MQTT Broker: ${MQTT_HOST}:${MQTT_PORT}"
bashio::log.info "Units: ${UNITS}"
bashio::log.info "Timezone: ${TZ}"

# Start the weather station bridge
bashio::log.info "Starting weather station service on port 80..."
exec /weatherbridge
