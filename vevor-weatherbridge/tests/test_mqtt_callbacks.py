"""Tests for MQTT callback functions.

Tests for MQTT API v2 callbacks:
- on_connect
- on_disconnect
- on_publish
"""



class TestMQTTCallbacks:
    """Test MQTT callback functions."""

    def test_on_connect_success(self, mock_mqtt_client):
        """Test successful MQTT connection callback."""
        import weatherstation

        weatherstation.mqtt_connected = False

        # Call callback with API v2 signature (reason_code=0 means success)
        weatherstation.on_connect(mock_mqtt_client, None, None, 0, None)

        assert weatherstation.mqtt_connected is True

    def test_on_connect_failure(self, mock_mqtt_client):
        """Test failed MQTT connection callback."""
        import weatherstation

        weatherstation.mqtt_connected = False

        # Call callback with non-zero reason code (connection refused)
        weatherstation.on_connect(mock_mqtt_client, None, None, 1, None)

        assert weatherstation.mqtt_connected is False

    def test_on_connect_various_error_codes(self, mock_mqtt_client):
        """Test various MQTT connection error codes."""
        import weatherstation

        # Test various error codes
        error_codes = [1, 2, 3, 4, 5]  # Various MQTT connection errors

        for code in error_codes:
            weatherstation.mqtt_connected = True
            weatherstation.on_connect(mock_mqtt_client, None, None, code, None)
            assert weatherstation.mqtt_connected is False, f"Should be False for error code {code}"

    def test_on_disconnect(self, mock_mqtt_client):
        """Test MQTT disconnection callback."""
        import weatherstation

        weatherstation.mqtt_connected = True

        # Call callback with API v2 signature
        weatherstation.on_disconnect(mock_mqtt_client, None, None, 1, None)

        assert weatherstation.mqtt_connected is False

    def test_on_disconnect_from_already_disconnected(self, mock_mqtt_client):
        """Test disconnect callback when already disconnected."""
        import weatherstation

        weatherstation.mqtt_connected = False

        # Should not raise any errors
        weatherstation.on_disconnect(mock_mqtt_client, None, None, 1, None)

        assert weatherstation.mqtt_connected is False

    def test_on_publish(self, mock_mqtt_client):
        """Test MQTT publish callback."""
        import weatherstation

        # Should not raise any exceptions
        weatherstation.on_publish(mock_mqtt_client, None, 123, None, None)

    def test_on_publish_various_mids(self, mock_mqtt_client):
        """Test publish callback with various message IDs."""
        import weatherstation

        # Various message IDs - should all succeed without exception
        for mid in [1, 100, 999, 12345]:
            weatherstation.on_publish(mock_mqtt_client, None, mid, None, None)


class TestMQTTConnectionState:
    """Test MQTT connection state management."""

    def test_initial_state_after_connect(self, mock_mqtt_client):
        """Test that mqtt_connected is properly set after successful connect."""
        import weatherstation

        weatherstation.mqtt_connected = False
        weatherstation.on_connect(mock_mqtt_client, None, None, 0, None)

        assert weatherstation.mqtt_connected is True

    def test_state_after_disconnect_then_reconnect(self, mock_mqtt_client):
        """Test connection state through disconnect/reconnect cycle."""
        import weatherstation

        # Start connected
        weatherstation.mqtt_connected = True

        # Disconnect
        weatherstation.on_disconnect(mock_mqtt_client, None, None, 1, None)
        assert weatherstation.mqtt_connected is False

        # Reconnect
        weatherstation.on_connect(mock_mqtt_client, None, None, 0, None)
        assert weatherstation.mqtt_connected is True
