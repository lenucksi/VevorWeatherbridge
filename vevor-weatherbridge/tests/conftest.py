"""Pytest configuration and fixtures for VevorWeatherbridge tests."""

import sys
from pathlib import Path
from unittest.mock import MagicMock, patch

import pytest

# Add the vevor-weatherbridge directory to the path so we can import weatherstation
sys.path.insert(0, str(Path(__file__).parent.parent))

# Mock MQTT client before any weatherstation imports
_mock_mqtt_client = MagicMock()
_mock_mqtt_class = MagicMock(return_value=_mock_mqtt_client)


@pytest.fixture(scope="session", autouse=True)
def mock_mqtt_globally():
    """Mock MQTT client globally before weatherstation module is imported."""
    with patch.dict("sys.modules", {"paho.mqtt.client": MagicMock()}):
        import paho.mqtt.client as mqtt

        mqtt.Client = _mock_mqtt_class
        yield _mock_mqtt_client


# Patch before weatherstation is imported anywhere
patch("paho.mqtt.client.Client", _mock_mqtt_class).start()


@pytest.fixture
def mock_mqtt():
    """Provide access to the mocked MQTT client."""
    _mock_mqtt_client.reset_mock()
    return _mock_mqtt_client


@pytest.fixture
def mock_mqtt_client():
    """Alternative fixture name for MQTT client mock."""
    _mock_mqtt_client.reset_mock()
    return _mock_mqtt_client


@pytest.fixture
def app(mock_mqtt):
    """Create Flask test application with mocked MQTT."""
    import weatherstation

    weatherstation.mqtt_client = mock_mqtt
    weatherstation.mqtt_connected = True
    weatherstation.app.config["TESTING"] = True
    return weatherstation.app


@pytest.fixture
def client(app):
    """Create Flask test client."""
    return app.test_client()


@pytest.fixture
def weatherstation_module(mock_mqtt):
    """Provide the weatherstation module with mocked MQTT."""
    import weatherstation

    weatherstation.mqtt_client = mock_mqtt
    weatherstation.mqtt_connected = True
    return weatherstation
