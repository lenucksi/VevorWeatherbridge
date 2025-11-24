#!/usr/bin/with-contenv bashio
# ==============================================================================
# Home Assistant Add-on: VEVOR Weather Station Bridge
# Starts the weather station bridge service
# ==============================================================================

set -e

bashio::log.info "Starting VEVOR Weather Station Bridge..."

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
# Check if user provided manual MQTT configuration
CONFIGURED_HOST=$(bashio::config 'mqtt_host')
CONFIGURED_PORT=$(bashio::config 'mqtt_port')
CONFIGURED_USER=$(bashio::config 'mqtt_user')
CONFIGURED_PASSWORD=$(bashio::config 'mqtt_password')

if bashio::var.has_value "${CONFIGURED_HOST}"; then
    # Use user-provided MQTT configuration
    bashio::log.info "Using manually configured MQTT broker"
    export MQTT_HOST="${CONFIGURED_HOST}"
    export MQTT_PORT="${CONFIGURED_PORT}"
    export MQTT_USER="${CONFIGURED_USER}"
    export MQTT_PASSWORD="${CONFIGURED_PASSWORD}"
else
    # Try to auto-detect Home Assistant's internal MQTT broker
    bashio::log.info "Attempting to auto-detect MQTT broker via Supervisor..."
    bashio::log.debug "Checking if MQTT service is available..."

    if bashio::services.available "mqtt"; then
        bashio::log.debug "MQTT service is available, fetching credentials..."
        export MQTT_HOST=$(bashio::services "mqtt" "host")
        export MQTT_PORT=$(bashio::services "mqtt" "port")
        export MQTT_USER=$(bashio::services "mqtt" "username")
        export MQTT_PASSWORD=$(bashio::services "mqtt" "password")

        bashio::log.info "Using Home Assistant MQTT broker from Supervisor"
        bashio::log.info "MQTT Host: ${MQTT_HOST}:${MQTT_PORT}"
        bashio::log.info "MQTT User: ${MQTT_USER}"
    else
        bashio::log.warning "MQTT auto-detection failed!"
        bashio::log.warning "The Supervisor MQTT service is not available."
        bashio::log.warning ""
        bashio::log.warning "Note: Auto-detection only works with the Mosquitto broker addon."
        bashio::log.warning "If using an external MQTT broker or the MQTT integration,"
        bashio::log.warning "you must configure MQTT manually in the addon settings."
        bashio::log.warning ""
        bashio::log.warning "Required settings for manual configuration:"
        bashio::log.warning "  - mqtt_host: your MQTT broker hostname/IP"
        bashio::log.warning "  - mqtt_port: MQTT port (default: 1883)"
        bashio::log.warning "  - mqtt_user: MQTT username (if required)"
        bashio::log.warning "  - mqtt_password: MQTT password (if required)"
        bashio::log.fatal "Cannot start without MQTT configuration."
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
