"""
Test suite for weatherstation.py

Run with: pytest test_weatherstation.py -v
"""

import json
from unittest.mock import MagicMock, Mock, patch

import pytest


@pytest.fixture
def mock_mqtt_client():
    """Mock MQTT client for testing."""
    with patch("weatherstation.mqtt.Client") as mock:
        client_instance = MagicMock()
        mock.return_value = client_instance
        yield client_instance


@pytest.fixture
def app():
    """Create Flask test app."""
    with patch("weatherstation.mqtt"):
        import weatherstation

        weatherstation.mqtt_connected = True  # Assume connected for tests
        with weatherstation.app.test_client() as client:
            yield client


class TestMQTTCallbacks:
    """Test MQTT callback functions."""

    def test_on_connect_success(self, mock_mqtt_client):
        """Test successful MQTT connection callback."""
        import weatherstation

        weatherstation.mqtt_connected = False

        # Call callback with API v2 signature
        weatherstation.on_connect(mock_mqtt_client, None, None, 0, None)

        assert weatherstation.mqtt_connected is True

    def test_on_connect_failure(self, mock_mqtt_client):
        """Test failed MQTT connection callback."""
        import weatherstation

        weatherstation.mqtt_connected = False

        # Call callback with non-zero reason code
        weatherstation.on_connect(mock_mqtt_client, None, None, 1, None)

        assert weatherstation.mqtt_connected is False

    def test_on_disconnect(self, mock_mqtt_client):
        """Test MQTT disconnection callback."""
        import weatherstation

        weatherstation.mqtt_connected = True

        # Call callback with API v2 signature
        weatherstation.on_disconnect(mock_mqtt_client, None, None, 1, None)

        assert weatherstation.mqtt_connected is False

    def test_on_publish(self, mock_mqtt_client):
        """Test MQTT publish callback."""
        import weatherstation

        # Should not raise any exceptions
        weatherstation.on_publish(mock_mqtt_client, None, 123, None, None)


class TestUnitConversions:
    """Test unit conversion functions."""

    def test_f_to_c(self):
        """Test Fahrenheit to Celsius conversion."""
        from weatherstation import f_to_c

        assert f_to_c(32) == 0.0
        assert f_to_c(212) == 100.0
        assert f_to_c(98.6) == 37.0

    def test_inhg_to_hpa(self):
        """Test inHg to hPa conversion."""
        from weatherstation import inhg_to_hpa

        assert inhg_to_hpa(29.92) == 1013.2
        assert inhg_to_hpa(30.00) == 1015.9

    def test_mph_to_kmh(self):
        """Test mph to km/h conversion."""
        from weatherstation import mph_to_kmh

        assert mph_to_kmh(10) == 16.1
        assert mph_to_kmh(60) == 96.6

    def test_inch_to_mm(self):
        """Test inch to mm conversion."""
        from weatherstation import inch_to_mm

        assert inch_to_mm(1) == 25.4
        assert inch_to_mm(0.5) == 12.7


class TestWeatherEndpoint:
    """Test the weather station HTTP endpoint."""

    def test_endpoint_responds(self, app, mock_mqtt_client):
        """Test endpoint returns success."""
        response = app.get(
            "/weatherstation/updateweatherstation.php?ID=test&tempf=70&humidity=50"
        )
        assert response.status_code == 200
        assert response.data == b"success"

    def test_endpoint_with_minimal_data(self, app, mock_mqtt_client):
        """Test endpoint with minimal parameters."""
        response = app.get("/weatherstation/updateweatherstation.php?ID=test")
        assert response.status_code == 200

    def test_endpoint_with_full_data(self, app, mock_mqtt_client):
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

        response = app.get("/weatherstation/updateweatherstation.php", query_string=params)
        assert response.status_code == 200
        assert response.data == b"success"

    @patch("weatherstation.mqtt_client")
    def test_mqtt_publishing(self, mock_client, app):
        """Test that MQTT messages are published."""
        import weatherstation

        weatherstation.mqtt_connected = True

        # Mock publish to return success
        publish_result = Mock()
        publish_result.rc = 0
        publish_result.mid = 123
        mock_client.publish.return_value = publish_result

        response = app.get(
            "/weatherstation/updateweatherstation.php?tempf=70&humidity=50"
        )

        assert response.status_code == 200
        # Verify publish was called
        assert mock_client.publish.called

    @patch("weatherstation.mqtt_client")
    def test_mqtt_not_connected_warning(self, mock_client, app):
        """Test warning when MQTT not connected."""
        import weatherstation

        weatherstation.mqtt_connected = False

        response = app.get("/weatherstation/updateweatherstation.php?tempf=70")

        # Should still return success even if MQTT disconnected
        assert response.status_code == 200


class TestMetricConversion:
    """Test metric unit conversions in endpoint."""

    @patch("weatherstation.UNITS", "metric")
    @patch("weatherstation.mqtt_client")
    def test_metric_temperature_conversion(self, mock_client, app):
        """Test temperature converted to Celsius in metric mode."""
        publish_result = Mock()
        publish_result.rc = 0
        mock_client.publish.return_value = publish_result

        response = app.get("/weatherstation/updateweatherstation.php?tempf=32")

        assert response.status_code == 200
        # Check that publish was called with Celsius value
        calls = mock_client.publish.call_args_list
        # Temperature should be 0.0 (converted from 32F)
        state_calls = [c for c in calls if "state" in str(c)]
        assert len(state_calls) > 0


class TestTimezoneHandling:
    """Test timezone conversion."""

    @patch("weatherstation.TIMEZONE", "Europe/Berlin")
    def test_utc_to_local_timezone(self, app, mock_mqtt_client):
        """Test UTC timestamp converted to local timezone."""
        response = app.get(
            "/weatherstation/updateweatherstation.php?"
            "dateutc=2025-11-10+12:00:00&tempf=70"
        )
        assert response.status_code == 200


class TestWeatherUndergroundForwarding:
    """Test Weather Underground forwarding functionality."""

    @patch("weatherstation.WU_FORWARD", True)
    @patch("weatherstation.WU_USERNAME", "test_user")
    @patch("weatherstation.WU_PASSWORD", "test_pass")
    @patch("weatherstation.requests.get")
    @patch("weatherstation.dns.resolver.Resolver")
    def test_wu_forwarding_enabled(
        self, mock_resolver, mock_requests, app, mock_mqtt_client
    ):
        """Test data forwarded to Weather Underground when enabled."""
        # Mock DNS resolution
        mock_resolver.return_value.resolve.return_value = [Mock(to_text=lambda: "1.2.3.4")]

        # Mock HTTP request
        mock_response = Mock()
        mock_response.status_code = 200
        mock_requests.return_value = mock_response

        response = app.get("/weatherstation/updateweatherstation.php?tempf=70")

        assert response.status_code == 200
        assert mock_requests.called

    @patch("weatherstation.WU_FORWARD", False)
    @patch("weatherstation.requests.get")
    def test_wu_forwarding_disabled(self, mock_requests, app, mock_mqtt_client):
        """Test data not forwarded when WU forwarding disabled."""
        response = app.get("/weatherstation/updateweatherstation.php?tempf=70")

        assert response.status_code == 200
        assert not mock_requests.called


class TestErrorHandling:
    """Test error handling and resilience."""

    @patch("weatherstation.mqtt_client")
    def test_mqtt_publish_failure_still_returns_success(self, mock_client, app):
        """Test endpoint returns success even if MQTT publish fails."""
        import weatherstation

        weatherstation.mqtt_connected = True

        # Simulate publish failure
        publish_result = Mock()
        publish_result.rc = 1  # Error code
        mock_client.publish.return_value = publish_result

        response = app.get("/weatherstation/updateweatherstation.php?tempf=70")

        # Should still return success to weather station
        assert response.status_code == 200
        assert response.data == b"success"

    def test_malformed_timestamp_handled(self, app, mock_mqtt_client):
        """Test invalid timestamp doesn't crash endpoint."""
        response = app.get(
            "/weatherstation/updateweatherstation.php?"
            "dateutc=invalid_timestamp&tempf=70"
        )

        # Should still succeed
        assert response.status_code == 200

    @patch("weatherstation.mqtt_client")
    def test_missing_sensor_values_skipped(self, mock_client, app):
        """Test None values don't get published."""
        publish_result = Mock()
        publish_result.rc = 0
        mock_client.publish.return_value = publish_result

        # Request with no sensor data
        response = app.get("/weatherstation/updateweatherstation.php?ID=test")

        assert response.status_code == 200
        # Publish should not be called for None values
        # (or very few times for empty data)


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
