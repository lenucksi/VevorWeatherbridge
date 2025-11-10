#!/usr/bin/env bash
# ==============================================================================
# Home Assistant Add-on: VEVOR Weather Station Bridge
# Starts the weather station bridge service
# ==============================================================================

set -e

echo "[INFO] Starting VEVOR Weather Station Bridge..."

# Read configuration from Home Assistant addon options
CONFIG_PATH=/data/options.json

# Read device configuration using jq
export DEVICE_NAME=$(jq -r '.device_name' $CONFIG_PATH)
export DEVICE_MANUFACTURER=$(jq -r '.device_manufacturer' $CONFIG_PATH)
export DEVICE_MODEL=$(jq -r '.device_model' $CONFIG_PATH)
export UNITS=$(jq -r '.units' $CONFIG_PATH)
export MQTT_PREFIX=$(jq -r '.mqtt_prefix' $CONFIG_PATH)
export TZ=$(jq -r '.timezone' $CONFIG_PATH)

# Generate device ID from device name (lowercase, replace spaces with underscores)
export DEVICE_ID=$(echo "${DEVICE_NAME}" | tr '[:upper:]' '[:lower:]' | tr ' ' '_')

# Weather Underground forwarding (optional)
WU_FORWARD_CONFIG=$(jq -r '.wu_forward' $CONFIG_PATH)
if [ "$WU_FORWARD_CONFIG" = "true" ]; then
    export WU_FORWARD="true"
    export WU_USERNAME=$(jq -r '.wu_username' $CONFIG_PATH)
    export WU_PASSWORD=$(jq -r '.wu_password' $CONFIG_PATH)
    echo "[INFO] Weather Underground forwarding enabled"
else
    export WU_FORWARD="false"
fi

# MQTT Configuration
# Try to use Home Assistant's internal MQTT broker first
MQTT_HOST_CONFIG=$(jq -r '.mqtt_host' $CONFIG_PATH)

if [ -z "$MQTT_HOST_CONFIG" ] || [ "$MQTT_HOST_CONFIG" = "null" ] || [ "$MQTT_HOST_CONFIG" = "" ]; then
    # Try to detect HA's internal MQTT broker via supervisor API
    echo "[INFO] Attempting to detect Home Assistant's internal MQTT broker..."

    # Try to get MQTT service info from supervisor
    if [ -f /run/secrets/mqtt_host ]; then
        export MQTT_HOST=$(cat /run/secrets/mqtt_host)
        export MQTT_PORT=$(cat /run/secrets/mqtt_port 2>/dev/null || echo "1883")
        export MQTT_USER=$(cat /run/secrets/mqtt_username 2>/dev/null || echo "")
        export MQTT_PASSWORD=$(cat /run/secrets/mqtt_password 2>/dev/null || echo "")
        echo "[INFO] Using Home Assistant's internal MQTT broker from secrets"
    else
        # Try common HA MQTT broker hostnames
        for host in "core-mosquitto" "mosquitto" "homeassistant.local"; do
            if ping -c 1 -W 1 "$host" > /dev/null 2>&1; then
                export MQTT_HOST="$host"
                export MQTT_PORT="1883"
                export MQTT_USER=""
                export MQTT_PASSWORD=""
                echo "[INFO] Found MQTT broker at: $host"
                break
            fi
        done

        if [ -z "$MQTT_HOST" ]; then
            echo "[FATAL] No MQTT broker configured! Please provide MQTT settings in addon configuration."
            exit 1
        fi
    fi
else
    # Use user-provided MQTT configuration
    echo "[INFO] Using external MQTT broker from configuration"
    export MQTT_HOST="$MQTT_HOST_CONFIG"
    export MQTT_PORT=$(jq -r '.mqtt_port' $CONFIG_PATH)
    export MQTT_USER=$(jq -r '.mqtt_user' $CONFIG_PATH)
    export MQTT_PASSWORD=$(jq -r '.mqtt_password' $CONFIG_PATH)
fi

echo "[INFO] Device: ${DEVICE_NAME} (${DEVICE_ID})"
echo "[INFO] MQTT Broker: ${MQTT_HOST}:${MQTT_PORT}"
echo "[INFO] Units: ${UNITS}"
echo "[INFO] Timezone: ${TZ}"

# Start the weather station bridge
echo "[INFO] Starting weather station service on port 80..."
exec python3 /app/weatherstation.py
