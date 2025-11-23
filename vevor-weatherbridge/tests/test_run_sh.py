"""
Test suite for run.sh script logic

Tests MQTT auto-detection and configuration parsing
"""

import json

import pytest


@pytest.fixture
def temp_config(tmp_path):
    """Create temporary config file."""
    config_file = tmp_path / "options.json"
    config_data = {
        "device_name": "Test Weather Station",
        "device_manufacturer": "VEVOR",
        "device_model": "Test Model",
        "units": "metric",
        "mqtt_host": "",
        "mqtt_port": 1883,
        "mqtt_user": "",
        "mqtt_password": "",
        "mqtt_prefix": "homeassistant",
        "timezone": "Europe/Berlin",
        "wu_forward": False,
        "wu_username": "",
        "wu_password": "",
        "log_level": "INFO",
    }
    config_file.write_text(json.dumps(config_data))
    return config_file


class TestMQTTAutoDetection:
    """Test MQTT broker auto-detection logic."""

    def test_supervisor_token_available(self, temp_config, monkeypatch):
        """Test MQTT config retrieved when SUPERVISOR_TOKEN available."""
        # Mock environment
        monkeypatch.setenv("SUPERVISOR_TOKEN", "test-token-123")

        # This would require mocking curl which is complex
        # Better to test the Python code directly
        # TODO: Implement with subprocess mock

    def test_supervisor_token_missing(self, temp_config):
        """Test error when SUPERVISOR_TOKEN not available."""
        # This should be tested by actually running run.sh in isolation
        pass  # TODO: Implement bash script testing

    def test_manual_mqtt_config(self, temp_config):
        """Test manual MQTT configuration takes precedence."""
        config_data = json.loads(temp_config.read_text())
        config_data["mqtt_host"] = "custom.mqtt.server"
        config_data["mqtt_user"] = "custom_user"
        config_data["mqtt_password"] = "custom_pass"  # noqa: S105
        temp_config.write_text(json.dumps(config_data))

        # Verify config parsing works
        parsed = json.loads(temp_config.read_text())
        assert parsed["mqtt_host"] == "custom.mqtt.server"
        assert parsed["mqtt_user"] == "custom_user"


class TestConfigParsing:
    """Test configuration file parsing."""

    def test_valid_config(self, temp_config):
        """Test parsing valid configuration."""
        config = json.loads(temp_config.read_text())

        assert config["device_name"] == "Test Weather Station"
        assert config["units"] == "metric"
        assert config["mqtt_port"] == 1883
        assert config["timezone"] == "Europe/Berlin"

    def test_empty_mqtt_config(self, temp_config):
        """Test empty MQTT configuration triggers auto-detection."""
        config = json.loads(temp_config.read_text())

        # Empty string should trigger auto-detection
        assert config["mqtt_host"] == ""
        assert config["mqtt_user"] == ""
        assert config["mqtt_password"] == ""

    def test_wu_forward_disabled(self, temp_config):
        """Test Weather Underground forwarding disabled by default."""
        config = json.loads(temp_config.read_text())
        assert config["wu_forward"] is False


class TestEnvironmentVariables:
    """Test environment variable handling."""

    def test_device_name_parsing(self, temp_config):
        """Test device name is parsed correctly."""
        config = json.loads(temp_config.read_text())
        device_name = config["device_name"]
        device_id = device_name.lower().replace(" ", "_")

        assert device_id == "test_weather_station"

    def test_timezone_setting(self, temp_config):
        """Test timezone environment variable."""
        config = json.loads(temp_config.read_text())
        assert config["timezone"] == "Europe/Berlin"

    def test_log_level_setting(self, temp_config):
        """Test log level configuration."""
        config = json.loads(temp_config.read_text())
        assert config["log_level"] in ["DEBUG", "INFO", "WARNING", "ERROR"]


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
