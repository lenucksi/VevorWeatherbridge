# Testing Guide

This document explains how to test the VevorWeatherbridge addon.

## Test Suite

The project includes comprehensive pytest-based tests:

### Test Files

- `vevor-weatherbridge/test_weatherstation.py` - Tests for weatherstation.py (20 tests)
- `vevor-weatherbridge/test_run_sh.py` - Tests for run.sh configuration logic (9 tests)
- `vevor-weatherbridge/pytest.ini` - Pytest configuration

### Running Tests Locally

```bash
cd vevor-weatherbridge
poetry install
poetry run pytest -v
```

### Test Coverage

**MQTT Callbacks** (test_weatherstation.py):

- on_connect success/failure
- on_disconnect handling
- on_publish callback

**Unit Conversions**:

- Fahrenheit to Celsius
- inHg to hPa
- mph to km/h
- inches to mm

**HTTP Endpoint**:

- Basic response
- Minimal data handling
- Full weather data processing
- MQTT publishing
- Connection state handling

**Metric/Imperial Mode**:

- Temperature conversions
- Pressure conversions
- Wind speed conversions
- Rainfall conversions

**Timezone Handling**:

- UTC to local conversion
- Timezone configuration

**Weather Underground Forwarding**:

- Enabled forwarding
- Disabled forwarding
- DNS resolution
- HTTP requests

**Error Handling**:

- MQTT publish failures
- Malformed timestamps
- Missing sensor values
- Resilience to errors

**Configuration** (test_run_sh.py):

- Config file parsing
- MQTT auto-detection logic
- Environment variable handling
- Device name parsing

## Local Development Container

Use VS Code devcontainer for full Home Assistant environment:

### Setup

1. Install "Remote - Containers" VS Code extension
2. Open project in VS Code
3. Click "Reopen in Container"
4. Run task "Start Home Assistant"
5. Access HA at <http://localhost:7123/>

### Available Tasks

- **Start Home Assistant** - Launches full HA environment
- **Run Addon Tests** - Executes pytest test suite
- **Rebuild Addon** - Builds Docker image locally

## Automated Testing

### Post-Edit Hook

After editing Python files, the hook automatically:

1. Runs ruff linter (with auto-fix)
2. Runs ruff formatter
3. Executes full test suite
4. Fails if any tests fail

Location: `.claude/hooks/post-edit-python.sh`

### Before Committing

Always run:

```bash
cd vevor-weatherbridge
poetry run ruff check . --fix
poetry run ruff format .
poetry run pytest -v
```

## Testing MQTT Integration

### Manual Testing in Home Assistant

1. Install addon in HA
2. Configure or leave MQTT auto-detect enabled
3. Set log_level to "DEBUG" in addon config
4. Send test data to endpoint:

   ```bash
   curl "http://YOUR_HA_IP:8099/weatherstation/updateweatherstation.php?\
ID=test&tempf=70&humidity=50&baromin=29.92"
   ```

1. Check logs for:
   - MQTT connection success
   - Sensor publishing
   - No errors

2. Verify in HA:
   - Go to Settings → Devices & Services
   - Look for "Weather Station" device
   - Check sensors are created and updating

### Testing MQTT Auto-Detection

**With Mosquitto Broker addon installed:**

- Addon should auto-detect via Supervisor API
- No manual MQTT configuration needed
- Check logs for "Using Home Assistant MQTT broker from Supervisor API"

**Without Mosquitto or custom MQTT:**

- Configure MQTT manually in addon settings
- Provide host, port, username, password
- Check logs for "Using external MQTT broker from configuration"

### Testing Callbacks

To verify MQTT callback fixes (v0.1.3+):

1. Enable DEBUG logging
2. Send weather data
3. Look for in logs:
   - "Connected to MQTT broker successfully"
   - "MQTT message published (mid=XXX)"
   - No TypeError exceptions

## Continuous Integration

### GitHub Actions Workflows

**build-addon.yml**:

- Triggers on: push to main, pull requests, tags
- Builds for all 5 architectures
- Tags with version from config.yaml

**dependency-review.yml**:

- Triggers on: PRs with dependencies label
- Runs pip-audit for vulnerabilities
- Runs bandit for security issues

**release.yml**:

- Triggers on: config.yaml version changes
- Creates GitHub release
- Extracts changelog

## Writing New Tests

### Test Structure

```python
class TestFeature:
    """Test description."""

    def test_specific_behavior(self, fixture):
        """Test specific behavior."""
        # Arrange
        ...

        # Act
        ...

        # Assert
        assert expected == actual
```

### Using Mocks

```python
from unittest.mock import Mock, patch

@patch("weatherstation.mqtt_client")
def test_with_mock(mock_client):
    # Mock behavior
    mock_client.publish.return_value = Mock(rc=0)

    # Test code
    ...

    # Verify
    assert mock_client.publish.called
```

### Fixtures

```python
@pytest.fixture
def mock_mqtt_client():
    """Create mock MQTT client."""
    with patch("weatherstation.mqtt.Client") as mock:
        yield mock.return_value
```

## Debugging Tests

### Run specific test

```bash
poetry run pytest test_weatherstation.py::TestMQTTCallbacks::test_on_connect_success -v
```

### Show print output

```bash
poetry run pytest -v -s
```

### Stop on first failure

```bash
poetry run pytest -x
```

### Run with coverage

```bash
poetry run pytest --cov=weatherstation --cov-report=html
```

## Common Issues

### "ModuleNotFoundError: No module named 'paho'"

```bash
poetry install
```

### "SUPERVISOR_TOKEN not available" (local testing)

- This is expected outside HA environment
- Configure MQTT manually in options.json for local testing
- Or use devcontainer with full HA stack

### Tests pass but addon fails in HA

- Check HA logs in Settings → System → Logs
- Enable DEBUG log level
- Verify MQTT broker is running
- Check Supervisor API access

## Best Practices

1. **Write tests BEFORE implementing features** (TDD)
2. **Run tests after EVERY code change**
3. **Never commit failing tests**
4. **Mock external dependencies** (MQTT, HTTP, DNS)
5. **Test edge cases and error conditions**
6. **Keep tests fast** (< 1 second total)
7. **Use descriptive test names**
8. **Document complex test setups**

## Resources

- [pytest Documentation](https://docs.pytest.org/)
- [unittest.mock Guide](https://docs.python.org/3/library/unittest.mock.html)
- [Home Assistant Addon Testing](https://developers.home-assistant.io/docs/add-ons/testing)
- [Flask Testing](https://flask.palletsprojects.com/en/latest/testing/)
