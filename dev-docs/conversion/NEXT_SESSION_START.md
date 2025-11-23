# VevorWeatherbridge - Home Assistant Addon Conversion Start Prompt

## Session Context

You are Claude, working on converting the **VevorWeatherbridge** project from a standalone Docker container into a proper **Home Assistant Add-on** that can be installed directly from the Home Assistant add-on store with minimal user configuration.

## Project Status

### Current State
- ✅ Working standalone Docker container (see [weatherstation.py](weatherstation.py), [Dockerfile](Dockerfile), [docker-compose.yml](docker-compose.yml))
- ✅ MQTT integration with Home Assistant auto-discovery
- ✅ Weather Underground data interception and forwarding
- ✅ Comprehensive Claude Code harness with hooks, skills, and quality automation
- ✅ Poetry development environment with all quality tools (ruff, mypy, bandit, semgrep, pip-audit, yamllint)

### What's Missing for HA Addon
- ❌ `config.yaml` - Home Assistant addon configuration with schema
- ❌ `build.json` - Multi-architecture build configuration
- ❌ `DOCS.md` - User-facing documentation
- ❌ `run.sh` - Addon entry point script
- ❌ `icon.png` - Addon icon (256x256)
- ❌ `CHANGELOG.md` - Version history
- ❌ Integration with HA's MQTT broker (currently requires external MQTT config)
- ❌ Configuration via HA addon options (currently uses environment variables)

## Ultimate Goal

**Create a "zero-configuration" Home Assistant Add-on** where users can:
1. Install the addon from the store with a few clicks
2. Optionally configure MQTT settings (or use HA's built-in MQTT broker automatically)
3. Configure their weather station device details
4. Start the addon

**External user action required:** DNS redirect from Weather Underground domains to their Home Assistant IP address (for the weather station).

## Claude Code Harness Available

This project has a **comprehensive quality assurance harness**. Always leverage it:

### Skills (Invoke Before/After Changes)
```bash
# Python quality checks
./.claude/skills/python-ci-skill/run.sh

# HA addon validation
./.claude/skills/ha-addon-skill/run.sh

# Security scanning
./.claude/skills/security-scan-skill/run.sh
```

### Quality Tools (via Poetry)
```bash
# All tools are installed in Poetry environment:
poetry run ruff check --fix .          # Linting
poetry run ruff format .                # Formatting
poetry run mypy weatherstation.py       # Type checking
poetry run bandit -r .                  # Security
poetry run pip-audit                    # Dependency vulnerabilities
poetry run yamllint config.yaml         # YAML validation
```

### Hooks
- Automatically trigger on file edits via `.claude/settings.local.json`
- Display recommendations (don't auto-execute)
- See [.claude/hooks/](.claude/hooks/) for individual hook scripts

## Key Documentation References

### Project Documentation
- **Project Rules:** [CLAUDE.md](CLAUDE.md) - Coding standards, automation rules, project context
- **Harness Guide:** [.claude/MANIFEST.md](.claude/MANIFEST.md) - Complete harness documentation
- **Quick Reference:** [.claude/QUICKSTART.md](.claude/QUICKSTART.md) - Command reference
- **Architecture:** [.claude/ARCHITECTURE.md](.claude/ARCHITECTURE.md) - System design
- **Summary:** [HARNESS_SUMMARY.md](HARNESS_SUMMARY.md) - Implementation overview

### Home Assistant Resources
- **Addon Development:** https://developers.home-assistant.io/docs/add-ons
- **Addon Configuration:** https://developers.home-assistant.io/docs/add-ons/configuration
- **Addon Tutorial:** https://developers.home-assistant.io/docs/add-ons/tutorial
- **MQTT Discovery:** https://www.home-assistant.io/integrations/mqtt/#mqtt-discovery
- **Supervisor API:** https://developers.home-assistant.io/docs/add-ons/communication

### Quality Standards (2025)
- **Ruff (Python):** https://docs.astral.sh/ruff/
- **Hadolint (Docker):** https://github.com/hadolint/hadolint
- **Semgrep (SAST):** https://semgrep.dev/docs/
- **Bandit (Security):** https://bandit.readthedocs.io/

## Immediate Next Tasks

### Phase 1: Core Addon Files (PRIORITY)

1. **Create `config.yaml`**
   - Define addon metadata (name, version, slug, description, arch)
   - Define configuration options with JSON schema
   - Set startup type, boot behavior, ports
   - Configure MQTT integration flags
   - Use skill: `./.claude/skills/ha-addon-skill/run.sh` to validate

2. **Create `build.json`**
   - Multi-architecture build configuration
   - Base images for: amd64, armv7, aarch64, armhf, i386
   - Use official HA base images

3. **Create `DOCS.md`**
   - User-facing documentation
   - Configuration options explained
   - Setup instructions
   - Troubleshooting guide
   - DNS redirect instructions (critical for weather station)

4. **Create `run.sh`**
   - Addon entry point
   - Read configuration from `/data/options.json`
   - Handle HA's built-in MQTT broker vs external broker
   - Start the Python application with proper environment variables

### Phase 2: Adaptation

5. **Modify `weatherstation.py`**
   - Read config from HA addon options (not environment variables)
   - Support HA's internal MQTT broker (hostname: `core-mosquitto`)
   - Add better error handling and logging for addon environment
   - Use skill: `./.claude/skills/python-ci-skill/run.sh` before/after

6. **Update `Dockerfile`**
   - Adapt for HA addon base images
   - Add `run.sh` as entrypoint
   - Ensure multi-architecture support
   - Add non-root user (security)
   - Use skill: `./.claude/skills/security-scan-skill/run.sh` to check

### Phase 3: Documentation & Polish

7. **Create `CHANGELOG.md`**
   - Version history
   - Semantic versioning

8. **Add `icon.png`**
   - 256x256 addon icon
   - Weather-themed icon

9. **Update `README.md`**
   - Add HA addon installation instructions
   - Keep Docker standalone instructions as fallback

## Important Implementation Notes

### MQTT Integration Strategy
```python
# In run.sh, detect MQTT configuration:
if [ -z "$MQTT_HOST" ]; then
    # Use HA's built-in MQTT broker
    export MQTT_HOST="core-mosquitto"
    export MQTT_PORT="1883"
    # Get credentials from Supervisor API
fi
```

### Configuration Schema Pattern
```yaml
# In config.yaml
options:
  device_name:
    type: str
    default: "Weather Station"
  units:
    type: list(metric|imperial)
    default: "metric"
  mqtt_host:
    type: str?  # Optional, defaults to HA internal
  mqtt_port:
    type: port?
    default: 1883
```

### Addon Lifecycle
1. User installs addon
2. User configures options via HA UI
3. HA calls `run.sh`
4. `run.sh` reads `/data/options.json`
5. `run.sh` sets environment variables
6. `run.sh` starts `weatherstation.py`
7. Application runs and logs to stdout (captured by HA)

## Token Efficiency Reminders

### When Working with Files
- Use `@repo:/path/to/file` MCP references instead of reading files when possible
- This reduces token usage by 90%

### When Searching/Exploring
- Use Task tool with Explore agent for open-ended searches
- Don't use multiple Grep/Glob calls sequentially

### Model Selection
- **Haiku:** Config validation, linting, simple file operations
- **Sonnet (you):** Refactoring, architecture, security remediation, addon creation

## Quality Checklist (Run Before Committing)

```bash
# 1. Python quality
./.claude/skills/python-ci-skill/run.sh

# 2. HA addon structure
./.claude/skills/ha-addon-skill/run.sh

# 3. Security
./.claude/skills/security-scan-skill/run.sh

# 4. Fix any issues found
poetry run ruff check --fix .
poetry run ruff format .

# 5. Validate again
./.claude/skills/python-ci-skill/run.sh
```

## Workflow Recommendations

### Creating New Files
1. Read existing similar files for context
2. Check HA addon examples/documentation
3. Create the file
4. Run appropriate skill to validate
5. Fix any issues
6. Validate again

### Modifying Existing Files
1. Read the file first (required by Claude Code)
2. Make changes using Edit tool
3. Run appropriate skill
4. Address findings
5. Test if possible

## Current Technical Stack

- **Python:** 3.12
- **Framework:** Flask (lightweight HTTP server)
- **MQTT Library:** paho-mqtt
- **Environment:** Docker (will become HA addon container)
- **Dependencies:** See [pyproject.toml](pyproject.toml)
- **Development:** Poetry for dependency management
- **Quality:** Ruff, MyPy, Bandit, Semgrep, pip-audit (all via Poetry)

## Expected Challenges & Solutions

### Challenge: MQTT Broker Detection
**Solution:** Use Supervisor API to detect if internal MQTT is available, fallback to user config

### Challenge: Configuration Format Change
**Solution:** Create adapter in `run.sh` that transforms `/data/options.json` to environment variables

### Challenge: Multi-Architecture Builds
**Solution:** Use HA's base images which handle this automatically

### Challenge: User DNS Configuration
**Solution:** Clear documentation in DOCS.md with examples for Pi-hole, router DNS, etc.

## Success Criteria

The addon is complete when:
- ✅ `ha-addon-skill` reports all required files present and valid
- ✅ All quality tools pass (python-ci-skill, security-scan-skill)
- ✅ Config schema validates correctly
- ✅ Documentation is complete and clear
- ✅ Multi-architecture build configuration is correct
- ✅ Integration with HA's MQTT broker works
- ✅ User can install and configure via HA UI
- ✅ Weather data flows from station → addon → MQTT → HA

## How to Start This Session

Begin with:
```
I'm continuing work on converting VevorWeatherbridge to a Home Assistant add-on.
I've read @NEXT_SESSION_START.md and @CLAUDE.md.
Let me first validate the current project state by running the ha-addon-skill.
```

Then invoke:
```bash
./.claude/skills/ha-addon-skill/run.sh
```

This will show you exactly what files are missing, and you can proceed from there.

## Remember

- **Always use the harness** - Skills are there to help you maintain quality
- **Read CLAUDE.md** for project-specific rules and standards
- **Check HA documentation** before making assumptions about addon behavior
- **Test incrementally** - Validate after each major change
- **Security first** - Run security-scan-skill before finalizing
- **Token efficient** - Use MCP references, Task tool for exploration, appropriate models

---

**Current Session Goal:** Create all required Home Assistant addon files and adapt the application to read configuration from HA's addon options system.

**Next Step:** Run `./.claude/skills/ha-addon-skill/run.sh` to see current status and missing files.
