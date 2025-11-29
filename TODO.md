# TODO

## Open

- Test MQTT auto-detection (without manual config) to fully validate AppArmor `network unix,` permission

## Completed - 2025-11-29

### Fixed Go time parsing bug
- Issue: Weather station sends timestamps with single-digit time components (e.g., "2025-11-28 18:58:7" instead of "18:58:07")
- Solution: Added `normalizeTimestamp()` function to pad single-digit hours/minutes/seconds
- All tests passing, no more parsing errors

### Re-added AppArmor profile with Supervisor API support
- Added `network unix,` permission required for Home Assistant Supervisor API access
- Enables MQTT auto-detection via bashio::services
- Based on APPARMOR_RESEARCH.md findings
- ✅ Deployed and tested: Works with manual MQTT config, data processing confirmed
- Auto-detection without manual config still needs testing

### Migrated from Python+Go to Go-only
- Preserved Python implementation in `addon-python` branch
- Removed all Python files, tests, and dependencies from main
- Updated CI/CD workflows: removed ruff, pytest, Python linting
- Updated renovate.json: removed pip_requirements manager
- Updated .pre-commit-config.yaml: removed ruff hooks
- Updated .gitignore: replaced Python patterns with Go patterns
- Kept `vevor-weatherbridge-go/` directory name to avoid breaking installed addons

### Previous logs (kept for reference)
time=2025-11-28T19:58:13.467+01:00 level=INFO msg="Starting VEVOR Weather Station Bridge (Go)" version=0.1.0 device_name="Weather Station" device_id=weather_station units=metric timezone=Europe/Berlin
time=2025-11-28T19:58:13.468+01:00 level=INFO msg="Connecting to MQTT broker" host=core-mosquitto port=1883
time=2025-11-28T19:58:13.848+01:00 level=INFO msg="MQTT connected" host=core-mosquitto port=1883
time=2025-11-28T19:58:13.849+01:00 level=INFO msg="HTTP server starting" addr=:80
time=2025-11-28T19:58:16.759+01:00 level=DEBUG msg="Received weather update request" path=/weatherstation/updateweatherstation.php query="ID=1234567890&PASSWORD=abcdefghijklmnopt&dateutc=2025-11-28+18%3A58%3A7&baromin=30.18&tempf=42.7&humidity=81&dewptf=37.2&rainin=0&dailyrainin=0&winddir=73&windspeedmph=0&windgustmph=0&UV=0&solarRadiation=0"
time=2025-11-28T19:58:16.759+01:00 level=WARN msg="Failed to parse dateutc" value="2025-11-28 18:58:7" error="parsing time \"2025-11-28 18:58:7\" as \"2006-01-02 15:04:05\": cannot parse \"7\" as \"05\""
time=2025-11-28T19:58:16.761+01:00 level=DEBUG msg="Published sensor config" sensor=barometric_pressure topic=homeassistant/sensor/weather_station_barometric_pressure/config
time=2025-11-28T19:58:16.763+01:00 level=DEBUG msg="Published sensor state" sensor=barometric_pressure value=1022.0
time=2025-11-28T19:58:16.765+01:00 level=DEBUG msg="Published sensor attributes" sensor=barometric_pressure
time=2025-11-28T19:58:16.765+01:00 level=DEBUG msg="Published sensor data" sensor=barometric_pressure value=1022.0
time=2025-11-28T19:58:16.766+01:00 level=DEBUG msg="Published sensor config" sensor=temperature topic=homeassistant/sensor/weather_station_temperature/config
time=2025-11-28T19:58:16.767+01:00 level=DEBUG msg="Published sensor state" sensor=temperature value=5.9
time=2025-11-28T19:58:16.769+01:00 level=DEBUG msg="Published sensor attributes" sensor=temperature
time=2025-11-28T19:58:16.769+01:00 level=DEBUG msg="Published sensor data" sensor=temperature value=5.9
time=2025-11-28T19:58:16.771+01:00 level=DEBUG msg="Published sensor config" sensor=humidity topic=homeassistant/sensor/weather_station_humidity/config
time=2025-11-28T19:58:16.772+01:00 level=DEBUG msg="Published sensor state" sensor=humidity value=81
time=2025-11-28T19:58:16.773+01:00 level=DEBUG msg="Published sensor attributes" sensor=humidity
time=2025-11-28T19:58:16.773+01:00 level=DEBUG msg="Published sensor data" sensor=humidity value=81
time=2025-11-28T19:58:16.773+01:00 level=DEBUG msg="Published sensor config" sensor=dew_point topic=homeassistant/sensor/weather_station_dew_point/config
time=2025-11-28T19:58:16.774+01:00 level=DEBUG msg="Published sensor state" sensor=dew_point value=2.9
time=2025-11-28T19:58:16.775+01:00 level=DEBUG msg="Published sensor attributes" sensor=dew_point
time=2025-11-28T19:58:16.775+01:00 level=DEBUG msg="Published sensor data" sensor=dew_point value=2.9
time=2025-11-28T19:58:16.775+01:00 level=DEBUG msg="Published sensor config" sensor=rainfall topic=homeassistant/sensor/weather_station_rainfall/config
time=2025-11-28T19:58:16.777+01:00 level=DEBUG msg="Published sensor state" sensor=rainfall value=0.0
time=2025-11-28T19:58:16.778+01:00 level=DEBUG msg="Published sensor attributes" sensor=rainfall
time=2025-11-28T19:58:16.779+01:00 level=DEBUG msg="Published sensor data" sensor=rainfall value=0.0
time=2025-11-28T19:58:16.780+01:00 level=DEBUG msg="Published sensor config" sensor=daily_rainfall topic=homeassistant/sensor/weather_station_daily_rainfall/config
time=2025-11-28T19:58:16.781+01:00 level=DEBUG msg="Published sensor state" sensor=daily_rainfall value=0.0
time=2025-11-28T19:58:16.782+01:00 level=DEBUG msg="Published sensor attributes" sensor=daily_rainfall
time=2025-11-28T19:58:16.782+01:00 level=DEBUG msg="Published sensor data" sensor=daily_rainfall value=0.0
time=2025-11-28T19:58:16.783+01:00 level=DEBUG msg="Published sensor config" sensor=wind_direction topic=homeassistant/sensor/weather_station_wind_direction/config
time=2025-11-28T19:58:16.783+01:00 level=DEBUG msg="Published sensor state" sensor=wind_direction value=73
time=2025-11-28T19:58:16.784+01:00 level=DEBUG msg="Published sensor attributes" sensor=wind_direction
time=2025-11-28T19:58:16.784+01:00 level=DEBUG msg="Published sensor data" sensor=wind_direction value=73
time=2025-11-28T19:58:16.785+01:00 level=DEBUG msg="Published sensor config" sensor=wind_speed topic=homeassistant/sensor/weather_station_wind_speed/config
time=2025-11-28T19:58:16.785+01:00 level=DEBUG msg="Published sensor state" sensor=wind_speed value=0.0
time=2025-11-28T19:58:16.787+01:00 level=DEBUG msg="Published sensor attributes" sensor=wind_speed
time=2025-11-28T19:58:16.787+01:00 level=DEBUG msg="Published sensor data" sensor=wind_speed value=0.0
time=2025-11-28T19:58:16.788+01:00 level=DEBUG msg="Published sensor config" sensor=wind_gust_speed topic=homeassistant/sensor/weather_station_wind_gust_speed/config
time=2025-11-28T19:58:16.789+01:00 level=DEBUG msg="Published sensor state" sensor=wind_gust_speed value=0.0
time=2025-11-28T19:58:16.790+01:00 level=DEBUG msg="Published sensor attributes" sensor=wind_gust_speed
time=2025-11-28T19:58:16.790+01:00 level=DEBUG msg="Published sensor data" sensor=wind_gust_speed value=0.0
time=2025-11-28T19:58:16.790+01:00 level=DEBUG msg="Published sensor config" sensor=uv_index topic=homeassistant/sensor/weather_station_uv_index/config
time=2025-11-28T19:58:16.791+01:00 level=DEBUG msg="Published sensor state" sensor=uv_index value=0
time=2025-11-28T19:58:16.792+01:00 level=DEBUG msg="Published sensor attributes" sensor=uv_index
time=2025-11-28T19:58:16.792+01:00 level=DEBUG msg="Published sensor data" sensor=uv_index value=0
time=2025-11-28T19:58:16.792+01:00 level=DEBUG msg="Published sensor config" sensor=solar_radiation topic=homeassistant/sensor/weather_station_solar_radiation/config
time=2025-11-28T19:58:16.793+01:00 level=DEBUG msg="Published sensor state" sensor=solar_radiation value=0.0
time=2025-11-28T19:58:16.793+01:00 level=DEBUG msg="Published sensor attributes" sensor=solar_radiation
time=2025-11-28T19:58:16.793+01:00 level=DEBUG msg="Published sensor data" sensor=solar_radiation value=0.0
time=2025-11-28T19:58:16.793+01:00 level=INFO msg="Processed weather update" sensors_published=11
time=2025-11-28T19:59:21.550+01:00 level=DEBUG msg="Received weather update request" path=/weatherstation/updateweatherstation.php query="ID=1234567890&PASSWORD=abcdefghijklmnopt&dateutc=2025-11-28+18%3A59%3A18&baromin=30.18&tempf=42.5&humidity=81&dewptf=37.0&rainin=0&dailyrainin=0&winddir=73&windspeedmph=0&windgustmph=0&UV=0&solarRadiation=0"
time=2025-11-28T19:59:21.552+01:00 level=DEBUG msg="Published sensor config" sensor=barometric_pressure topic=homeassistant/sensor/weather_station_barometric_pressure/config
time=2025-11-28T19:59:21.554+01:00 level=DEBUG msg="Published sensor state" sensor=barometric_pressure value=1022.0
time=2025-11-28T19:59:21.555+01:00 level=DEBUG msg="Published sensor attributes" sensor=barometric_pressure
time=2025-11-28T19:59:21.555+01:00 level=DEBUG msg="Published sensor data" sensor=barometric_pressure value=1022.0
time=2025-11-28T19:59:21.556+01:00 level=DEBUG msg="Published sensor config" sensor=temperature topic=homeassistant/sensor/weather_station_temperature/config
time=2025-11-28T19:59:21.557+01:00 level=DEBUG msg="Published sensor state" sensor=temperature value=5.8
time=2025-11-28T19:59:21.558+01:00 level=DEBUG msg="Published sensor attributes" sensor=temperature
time=2025-11-28T19:59:21.558+01:00 level=DEBUG msg="Published sensor data" sensor=temperature value=5.8
time=2025-11-28T19:59:21.558+01:00 level=DEBUG msg="Published sensor config" sensor=humidity topic=homeassistant/sensor/weather_station_humidity/config
time=2025-11-28T19:59:21.559+01:00 level=DEBUG msg="Published sensor state" sensor=humidity value=81
time=2025-11-28T19:59:21.559+01:00 level=DEBUG msg="Published sensor attributes" sensor=humidity
time=2025-11-28T19:59:21.559+01:00 level=DEBUG msg="Published sensor data" sensor=humidity value=81
time=2025-11-28T19:59:21.559+01:00 level=DEBUG msg="Published sensor config" sensor=dew_point topic=homeassistant/sensor/weather_station_dew_point/config
time=2025-11-28T19:59:21.560+01:00 level=DEBUG msg="Published sensor state" sensor=dew_point value=2.8
time=2025-11-28T19:59:21.560+01:00 level=DEBUG msg="Published sensor attributes" sensor=dew_point
time=2025-11-28T19:59:21.560+01:00 level=DEBUG msg="Published sensor data" sensor=dew_point value=2.8
time=2025-11-28T19:59:21.560+01:00 level=DEBUG msg="Published sensor config" sensor=rainfall topic=homeassistant/sensor/weather_station_rainfall/config
time=2025-11-28T19:59:21.560+01:00 level=DEBUG msg="Published sensor state" sensor=rainfall value=0.0
time=2025-11-28T19:59:21.562+01:00 level=DEBUG msg="Published sensor attributes" sensor=rainfall
time=2025-11-28T19:59:21.562+01:00 level=DEBUG msg="Published sensor data" sensor=rainfall value=0.0
time=2025-11-28T19:59:21.563+01:00 level=DEBUG msg="Published sensor config" sensor=daily_rainfall topic=homeassistant/sensor/weather_station_daily_rainfall/config
time=2025-11-28T19:59:21.564+01:00 level=DEBUG msg="Published sensor state" sensor=daily_rainfall value=0.0
time=2025-11-28T19:59:21.565+01:00 level=DEBUG msg="Published sensor attributes" sensor=daily_rainfall
time=2025-11-28T19:59:21.565+01:00 level=DEBUG msg="Published sensor data" sensor=daily_rainfall value=0.0
time=2025-11-28T19:59:21.565+01:00 level=DEBUG msg="Published sensor config" sensor=wind_direction topic=homeassistant/sensor/weather_station_wind_direction/config
time=2025-11-28T19:59:21.565+01:00 level=DEBUG msg="Published sensor state" sensor=wind_direction value=73
time=2025-11-28T19:59:21.566+01:00 level=DEBUG msg="Published sensor attributes" sensor=wind_direction
time=2025-11-28T19:59:21.566+01:00 level=DEBUG msg="Published sensor data" sensor=wind_direction value=73
time=2025-11-28T19:59:21.566+01:00 level=DEBUG msg="Published sensor config" sensor=wind_speed topic=homeassistant/sensor/weather_station_wind_speed/config
time=2025-11-28T19:59:21.566+01:00 level=DEBUG msg="Published sensor state" sensor=wind_speed value=0.0
time=2025-11-28T19:59:21.567+01:00 level=DEBUG msg="Published sensor attributes" sensor=wind_speed
time=2025-11-28T19:59:21.567+01:00 level=DEBUG msg="Published sensor data" sensor=wind_speed value=0.0
time=2025-11-28T19:59:21.567+01:00 level=DEBUG msg="Published sensor config" sensor=wind_gust_speed topic=homeassistant/sensor/weather_station_wind_gust_speed/config
time=2025-11-28T19:59:21.567+01:00 level=DEBUG msg="Published sensor state" sensor=wind_gust_speed value=0.0
time=2025-11-28T19:59:21.567+01:00 level=DEBUG msg="Published sensor attributes" sensor=wind_gust_speed
time=2025-11-28T19:59:21.567+01:00 level=DEBUG msg="Published sensor data" sensor=wind_gust_speed value=0.0
time=2025-11-28T19:59:21.568+01:00 level=DEBUG msg="Published sensor config" sensor=uv_index topic=homeassistant/sensor/weather_station_uv_index/config
time=2025-11-28T19:59:21.568+01:00 level=DEBUG msg="Published sensor state" sensor=uv_index value=0
time=2025-11-28T19:59:21.568+01:00 level=DEBUG msg="Published sensor attributes" sensor=uv_index
time=2025-11-28T19:59:21.568+01:00 level=DEBUG msg="Published sensor data" sensor=uv_index value=0
time=2025-11-28T19:59:21.569+01:00 level=DEBUG msg="Published sensor config" sensor=solar_radiation topic=homeassistant/sensor/weather_station_solar_radiation/config
time=2025-11-28T19:59:21.569+01:00 level=DEBUG msg="Published sensor state" sensor=solar_radiation value=0.0
time=2025-11-28T19:59:21.569+01:00 level=DEBUG msg="Published sensor attributes" sensor=solar_radiation
time=2025-11-28T19:59:21.569+01:00 level=DEBUG msg="Published sensor data" sensor=solar_radiation value=0.0
time=2025-11-28T19:59:21.569+01:00 level=INFO msg="Processed weather update" sensors_published=11


## Completed

- [x] Fix Go warnings with golangci-lint (4 errcheck warnings fixed)

- [x] Add Go tests to CI workflow (added `go test -v ./...` step to ci.yml)

- [x] Clean up RELEASE_NOTES files from .github/
  - Deleted `RELEASE_NOTES_v0.1.6.md` and `RELEASE_NOTES_v0.1.7.md`
  - These were historical artifacts not used by automated release workflow

- [x] pyproject.toml location decision
  - **Result**: Keep at root (standard location per PEP 517/518/621)
  - Moving would break `uv sync` and require `--config` flags for all tools

- [x] pre-commit/action as GitHub Action
  - **Result**: Not needed - `pre-commit/action` is in maintenance mode
  - CI already runs the same linters (ruff, hadolint, golangci-lint)
  - Local pre-commit hooks are sufficient

- [x] Research HomeAssistant bronze/silver/gold integration quality requirements
  - **Result**: Quality Scale does NOT apply to this project (it's for native HA integrations, not MQTT bridges)
  - See [dev-docs/homeassistant-integration-notes.md](dev-docs/homeassistant-integration-notes.md) for details

- [x] Consider Go rewrite for smaller Docker images (~5-10 MB vs ~150 MB)
  - See [dev-docs/go-rust-migration-analysis.md](dev-docs/go-rust-migration-analysis.md) for analysis

- [x] Remediate CodeQL security findings
  - Fixed token permissions in release-please.yml and build-addon.yml
  - Replaced `go install @latest` with GitHub Actions for govulncheck and gosec
  - Documented Dockerfile base image as accepted risk
  - Full analysis in [dev-docs/security/CODEQL_FINDINGS_ANALYSIS.md](dev-docs/security/CODEQL_FINDINGS_ANALYSIS.md)

## Future Considerations
