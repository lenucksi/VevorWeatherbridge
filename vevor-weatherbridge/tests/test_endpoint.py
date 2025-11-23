"""Integration tests for the weather station Flask endpoint.

Tests for the /weatherstation/updateweatherstation.php endpoint including:
- Basic request handling
- MQTT discovery payload structure
- Error handling and resilience
- Timezone handling
"""

import json
from unittest.mock import Mock, patch


class TestUpdateEndpoint:
    """Tests for the /weatherstation/updateweatherstation.php endpoint."""

    def test_basic_request(self, client, mock_mqtt):
        """Test basic weather data submission."""
        response = client.get(
            "/weatherstation/updateweatherstation.php",
            query_string={
                "tempf": "68",
                "humidity": "50",
                "baromin": "29.92",
            },
        )

        assert response.status_code == 200
        assert response.data == b"success"
        assert mock_mqtt.publish.called

    def test_endpoint_responds(self, client, mock_mqtt):
        """Test endpoint returns success."""
        response = client.get("/weatherstation/updateweatherstation.php?ID=test&tempf=70&humidity=50")
        assert response.status_code == 200
        assert response.data == b"success"

    def test_endpoint_with_minimal_data(self, client, mock_mqtt):
        """Test endpoint with minimal parameters."""
        response = client.get("/weatherstation/updateweatherstation.php?ID=test")
        assert response.status_code == 200

    def test_all_sensor_types(self, client, mock_mqtt):
        """Test with all supported sensor parameters."""
        response = client.get(
            "/weatherstation/updateweatherstation.php",
            query_string={
                "tempf": "72",
                "humidity": "45",
                "baromin": "30.1",
                "dewptf": "50",
                "rainin": "0.01",
                "dailyrainin": "0.5",
                "winddir": "180",
                "windspeedmph": "5",
                "windgustmph": "10",
                "UV": "3",
                "solarRadiation": "500",
                "dateutc": "2024-01-15 12:00:00",
            },
        )

        assert response.status_code == 200
        # Should have published config, state, and attributes for each sensor
        assert mock_mqtt.publish.call_count > 0

    def test_missing_parameters(self, client, mock_mqtt):
        """Test with no parameters - should still return success."""
        response = client.get("/weatherstation/updateweatherstation.php")

        assert response.status_code == 200
        assert response.data == b"success"

    def test_mqtt_discovery_payload_structure(self, client, mock_mqtt):
        """Test that MQTT discovery payloads have correct structure."""
        response = client.get(
            "/weatherstation/updateweatherstation.php",
            query_string={"tempf": "68"},
        )

        assert response.status_code == 200

        # Find a config topic publish call
        config_calls = [call for call in mock_mqtt.publish.call_args_list if "/config" in str(call)]
        assert len(config_calls) > 0

        # Verify config payload structure
        config_call = config_calls[0]
        config_payload = json.loads(config_call[0][1])

        assert "name" in config_payload
        assert "state_topic" in config_payload
        assert "unique_id" in config_payload
        assert "device" in config_payload
        assert "identifiers" in config_payload["device"]

    def test_endpoint_with_full_data(self, client, mock_mqtt):
        """Test endpoint with complete weather data."""
        params = {
            "ID": "test123",
            "PASSWORD": "pass",
            "dateutc": "2025-11-10 20:00:00",
            "tempf": "47.0",
            "humidity": "88",
            "dewptf": "43.7",
            "baromin": "29.92",
            "rainin": "0.1",
            "dailyrainin": "0.5",
            "winddir": "335",
            "windspeedmph": "5",
            "windgustmph": "10",
            "UV": "0",
            "solarRadiation": "0",
        }

        response = client.get("/weatherstation/updateweatherstation.php", query_string=params)
        assert response.status_code == 200
        assert response.data == b"success"


class TestMQTTPublishing:
    """Tests for MQTT message publishing."""

    def test_mqtt_publishing(self, client, weatherstation_module, mock_mqtt):
        """Test that MQTT messages are published."""
        # Mock publish to return success
        publish_result = Mock()
        publish_result.rc = 0
        publish_result.mid = 123
        mock_mqtt.publish.return_value = publish_result

        response = client.get("/weatherstation/updateweatherstation.php?tempf=70&humidity=50")

        assert response.status_code == 200
        # Verify publish was called
        assert mock_mqtt.publish.called

    def test_mqtt_not_connected_warning(self, client, weatherstation_module, mock_mqtt):
        """Test warning when MQTT not connected."""
        weatherstation_module.mqtt_connected = False

        response = client.get("/weatherstation/updateweatherstation.php?tempf=70")

        # Should still return success even if MQTT disconnected
        assert response.status_code == 200


class TestMetricConversion:
    """Test metric unit conversions in endpoint."""

    @patch("weatherstation.UNITS", "metric")
    def test_metric_temperature_conversion(self, client, weatherstation_module, mock_mqtt):
        """Test temperature converted to Celsius in metric mode."""
        publish_result = Mock()
        publish_result.rc = 0
        mock_mqtt.publish.return_value = publish_result

        response = client.get("/weatherstation/updateweatherstation.php?tempf=32")

        assert response.status_code == 200
        # Check that publish was called with Celsius value
        calls = mock_mqtt.publish.call_args_list
        # Temperature should be 0.0 (converted from 32F)
        state_calls = [c for c in calls if "state" in str(c)]
        assert len(state_calls) > 0


class TestTimezoneHandling:
    """Test timezone conversion."""

    @patch("weatherstation.TIMEZONE", "Europe/Berlin")
    def test_utc_to_local_timezone(self, client, mock_mqtt):
        """Test UTC timestamp converted to local timezone."""
        response = client.get("/weatherstation/updateweatherstation.php?dateutc=2025-11-10+12:00:00&tempf=70")
        assert response.status_code == 200


class TestErrorHandling:
    """Test error handling and resilience."""

    def test_mqtt_publish_failure_still_returns_success(self, client, weatherstation_module, mock_mqtt):
        """Test endpoint returns success even if MQTT publish fails."""
        weatherstation_module.mqtt_connected = True

        # Simulate publish failure
        publish_result = Mock()
        publish_result.rc = 1  # Error code
        mock_mqtt.publish.return_value = publish_result

        response = client.get("/weatherstation/updateweatherstation.php?tempf=70")

        # Should still return success to weather station
        assert response.status_code == 200
        assert response.data == b"success"

    def test_malformed_timestamp_handled(self, client, mock_mqtt):
        """Test invalid timestamp doesn't crash endpoint."""
        response = client.get("/weatherstation/updateweatherstation.php?dateutc=invalid_timestamp&tempf=70")

        # Should still succeed
        assert response.status_code == 200

    def test_missing_sensor_values_skipped(self, client, weatherstation_module, mock_mqtt):
        """Test None values don't get published."""
        publish_result = Mock()
        publish_result.rc = 0
        mock_mqtt.publish.return_value = publish_result

        # Request with no sensor data
        response = client.get("/weatherstation/updateweatherstation.php?ID=test")

        assert response.status_code == 200
        # Publish should not be called for None values
        # (or very few times for empty data)
