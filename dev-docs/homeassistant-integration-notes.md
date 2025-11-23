# Home Assistant Integration Notes

This document clarifies the relationship between VevorWeatherbridge and Home Assistant's quality standards.

## TL;DR

**VevorWeatherbridge is a Home Assistant Add-on** that communicates via MQTT Discovery.

- It is NOT a native integration (Python module in HA Core)
- It is NOT a standalone Docker container (it's managed by HA Supervisor)
- The Integration Quality Scale does NOT directly apply, but we assess against its intent

## Architecture

```text
┌─────────────────────────────────────────────────────────────────────┐
│                    Home Assistant Supervisor                        │
│  ┌──────────────────────────────────────────────────────────────┐   │
│  │              VevorWeatherbridge Add-on                       │   │
│  │  ┌─────────────────┐    MQTT    ┌──────────────────────┐    │   │
│  │  │ Flask HTTP      │ ────────►  │ HA MQTT Integration  │    │   │
│  │  │ (port 8099)     │  Discovery │ (auto-discovery)     │    │   │
│  │  └─────────────────┘            └──────────────────────┘    │   │
│  │         ▲                                                    │   │
│  │         │ HTTP POST                                          │   │
│  └─────────┼────────────────────────────────────────────────────┘   │
│            │                                                         │
└────────────┼─────────────────────────────────────────────────────────┘
             │
    ┌────────┴────────┐
    │ VEVOR Weather   │  (DNS redirected to HA IP)
    │ Station         │
    └─────────────────┘
```

### Integration Types Comparison

| Type | Location | Quality Scale | This Project |
|------|----------|---------------|--------------|
| **Native Integration** | `homeassistant/components/` | Yes - Bronze/Silver/Gold/Platinum | No |
| **HACS Integration** | Custom components | Community guidelines | No |
| **Home Assistant Add-on** | Docker container in Supervisor | Add-on guidelines | **Yes** |
| **Standalone Docker** | External to HA | None | No |

## Home Assistant Add-on Guidelines Compliance

Reference: [Home Assistant Add-on Development](https://developers.home-assistant.io/docs/add-ons)

### Required Files

| File | Status | Notes |
|------|--------|-------|
| `config.yaml` | ✅ Present | Full schema with options validation |
| `Dockerfile` | ✅ Present | Builds from HA base images |
| `run.sh` | ✅ Present | Uses bashio for HA option integration |
| `repository.yaml` | ✅ Present | At repo root for add-on discovery |

### Recommended Files

| File | Status | Notes |
|------|--------|-------|
| `DOCS.md` | ✅ Present | User documentation |
| `CHANGELOG.md` | ✅ Present | Version history |
| `README.md` | ✅ Present | Project overview |
| `icon.png` | ✅ Present | 256x256 add-on icon |
| `logo.png` | ❌ Missing | Should be added for store display |
| `apparmor.txt` | ❌ Missing | Optional, adds +1 security point |

### Security Configuration

Reference: [Add-on Security](https://developers.home-assistant.io/docs/add-ons/configuration/)

Current `config.yaml` security settings:

- `services: [mqtt:need]` - Declares MQTT dependency
- `ports: 80/tcp: 8099` - Exposes HTTP endpoint
- `map: [config:rw]` - Minimal filesystem access

**AppArmor Profile:**

- Not currently implemented
- Would add +1 to security rating
- For network-only add-on, default policy is sufficient
- Custom profile recommended for production

## Quality Assessment (Intent Transfer)

Since the Integration Quality Scale applies to native integrations, not add-ons, we assess VevorWeatherbridge against the **intent** of those requirements:

### Bronze-Equivalent (Baseline)

| Requirement Intent | Status | Implementation |
|-------------------|--------|----------------|
| Easy setup via UI | ✅ Pass | `config.yaml` schema with HA options UI |
| Basic documentation | ✅ Pass | `DOCS.md`, `README.md` |
| Automated tests | ✅ Pass | pytest suite with 85%+ coverage |
| Code quality tooling | ✅ Pass | ruff, mypy, bandit in CI |

**Bronze Result:** PASS

### Silver-Equivalent (Reliability)

| Requirement Intent | Status | Implementation |
|-------------------|--------|----------------|
| Identified maintainer | ✅ Pass | `repository.yaml` with maintainer |
| Error recovery | ✅ Pass | MQTT reconnection on disconnect |
| Graceful degradation | ✅ Pass | Returns "success" even on MQTT failure |
| Connection re-establishment | ✅ Pass | `on_disconnect` callback triggers reconnect |

**Silver Result:** PASS

### Gold-Equivalent (Best UX)

| Requirement Intent | Status | Implementation |
|-------------------|--------|----------------|
| Auto-discovery | ✅ Pass | Full MQTT Discovery with device grouping |
| Device information | ✅ Pass | Manufacturer, model, identifiers in discovery |
| Proper device classes | ✅ Pass | temperature, humidity, pressure, etc. |
| Translations/i18n | ❌ Fail | Not implemented |
| Reconfigurable without restart | ⚠️ Partial | Requires add-on restart |

**Gold Result:** PARTIAL (missing translations)

### Platinum-Equivalent (Excellence)

| Requirement Intent | Status | Implementation |
|-------------------|--------|----------------|
| Full type annotations | ⚠️ Partial | Some typing, not comprehensive |
| Fully asynchronous | ❌ Fail | Flask is synchronous WSGI |
| Optimized data handling | ✅ Pass | Efficient MQTT publishing |

**Platinum Result:** FAIL (architectural - Flask is synchronous)

### Summary

| Tier | Status |
|------|--------|
| Bronze | ✅ PASS |
| Silver | ✅ PASS |
| Gold | ⚠️ PARTIAL |
| Platinum | ❌ FAIL |

## What We Provide via MQTT Discovery

VevorWeatherbridge publishes to Home Assistant using [MQTT Discovery](https://www.home-assistant.io/integrations/mqtt/#mqtt-discovery):

### Device Registration

```json
{
  "device": {
    "identifiers": ["vevor_weather_station"],
    "name": "Weather Station",
    "manufacturer": "VEVOR",
    "model": "7-in-1 Weather Station"
  }
}
```

### Sensors Published

- Temperature (°C/°F)
- Humidity (%)
- Dew Point (°C/°F)
- Barometric Pressure (hPa/inHg)
- Wind Speed (km/h/mph)
- Wind Gust (km/h/mph)
- Wind Direction (degrees + cardinal)
- Rain Rate (mm/h or in/h)
- Daily Rain (mm or in)
- UV Index
- Solar Radiation (W/m²)

## If You Wanted a Native Integration

Creating a native Home Assistant integration would require:

### Minimum Requirements (Bronze)

1. **Config Flow**: UI-based setup in `config_flow.py`
2. **Manifest**: `manifest.json` with dependencies, codeowners
3. **Sensor Platform**: `sensor.py` with entity definitions
4. **Tests**: pytest tests in `tests/components/vevor_weather/`
5. **Documentation**: Entry in Home Assistant docs

### Estimated Effort

| Task | Hours |
|------|-------|
| Learn HA integration architecture | 4-8 |
| Implement config flow | 4-6 |
| Implement sensor platform | 4-6 |
| Write tests | 4-8 |
| Documentation | 2-4 |
| Review process | Variable |
| **Total** | **18-32+ hours** |

### Tradeoffs: Add-on vs Native Integration

| Aspect | Add-on (current) | Native Integration |
|--------|-----------------|-------------------|
| Setup | Add-on repository URL | Built-in or HACS |
| Configuration | Supervisor UI | HA Config Flow |
| Maintenance | Independent | Tied to HA releases |
| User reach | Anyone with Supervisor | All HA users |
| Review process | None required | HA Core review |
| Update control | Full control | HA release cycle |
| Resource isolation | Docker container | In HA process |

## Recommendation

**Stay with the Add-on approach** because:

1. **Already works**: MQTT auto-discovery is fully supported
2. **Lower maintenance**: No dependency on HA release cycles
3. **Broader compatibility**: Works with any MQTT-capable system
4. **Simpler architecture**: Single container, clear responsibilities
5. **User-friendly**: GUI configuration via Supervisor

A native integration would only be beneficial if:

- You want "official" status in Home Assistant
- You need deep HA-specific features (areas, device registry manipulation)
- You want to participate in the Home Assistant ecosystem directly

## Gaps to Address

### High Priority

- [ ] Create `logo.png` for add-on store display
- [ ] Add `apparmor.txt` security profile

### Medium Priority

- [ ] Add translations support (if feasible for add-on)
- [ ] Consider async framework migration (Quart/FastAPI) for Platinum

### Low Priority

- [ ] Full type annotations for all functions

## References

- [Home Assistant Add-on Development](https://developers.home-assistant.io/docs/add-ons)
- [Add-on Configuration](https://developers.home-assistant.io/docs/add-ons/configuration/)
- [Add-on Security / AppArmor](https://developers.home-assistant.io/docs/add-ons/security/)
- [Integration Quality Scale](https://developers.home-assistant.io/docs/core/integration-quality-scale/)
- [MQTT Discovery](https://www.home-assistant.io/integrations/mqtt/#mqtt-discovery)
- [Quality Scale Overview](https://www.home-assistant.io/docs/quality_scale/)
