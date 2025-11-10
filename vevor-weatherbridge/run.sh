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
export LOG_LEVEL=$(jq -r '.log_level' $CONFIG_PATH)

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
    # Use Supervisor Services API to get MQTT credentials
    echo "[INFO] Attempting to auto-detect MQTT broker via Supervisor API..."

    # SUPERVISOR_TOKEN is set by HA addon environment
    # Fallback to empty if not available (for local testing)
    SUPERVISOR_TOKEN="${SUPERVISOR_TOKEN:-}"

    # Check if we have supervisor access
    if [ -n "$SUPERVISOR_TOKEN" ]; then
        echo "[DEBUG] SUPERVISOR_TOKEN is available, querying API..."

        # Query Supervisor Services API for MQTT config
        MQTT_CONFIG=$(curl -sSL -H "Authorization: Bearer ${SUPERVISOR_TOKEN}" \
                      http://supervisor/services/mqtt 2>/dev/null || echo "")

        if [ -n "$MQTT_CONFIG" ] && [ "$MQTT_CONFIG" != "null" ]; then
            export MQTT_HOST=$(echo "$MQTT_CONFIG" | jq -r '.data.host // empty')
            export MQTT_PORT=$(echo "$MQTT_CONFIG" | jq -r '.data.port // 1883')
            export MQTT_USER=$(echo "$MQTT_CONFIG" | jq -r '.data.username // empty')
            export MQTT_PASSWORD=$(echo "$MQTT_CONFIG" | jq -r '.data.password // empty')

            if [ -n "$MQTT_HOST" ]; then
                echo "[INFO] Using Home Assistant MQTT broker from Supervisor API"
                echo "[INFO] MQTT Host: ${MQTT_HOST}:${MQTT_PORT}"
                echo "[INFO] MQTT User: ${MQTT_USER}"
            else
                echo "[ERROR] Supervisor API returned empty MQTT configuration"
                echo "[FATAL] No MQTT broker configured! Please install Mosquitto broker addon or configure MQTT manually."
                exit 1
            fi
        else
            echo "[ERROR] Could not retrieve MQTT config from Supervisor API"
            echo "[FATAL] No MQTT broker configured! Please install Mosquitto broker addon or configure MQTT manually."
            exit 1
        fi
    else
        echo "[ERROR] SUPERVISOR_TOKEN not available"
        echo "[FATAL] Cannot auto-detect MQTT broker. Please configure MQTT settings in addon configuration."
        exit 1
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
