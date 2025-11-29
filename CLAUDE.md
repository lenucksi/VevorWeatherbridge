# Claude Code Project Rules - VevorWeatherbridge

## Project Context

This is a Home Assistant add-on that intercepts VEVOR weather station data and forwards it to Home Assistant via MQTT.
The add-on runs as a containerized service within Home Assistant OS, with automatic MQTT broker detection and sensor auto-discovery.

External setup required by the user:

- DNS redirect from Weather Underground domains (rtupdate.wunderground.com) to Home Assistant IP
- Weather station configured to send data to the intercepted endpoint

**Tech Stack:**

- Go 1.24+ (net/http, paho.mqtt.golang, log/slog)
- Docker (multi-stage builds)
- Home Assistant Addon Framework
- MQTT protocol

## Code Quality Standards (2025)

### Go

- **Linter:** `golangci-lint` with comprehensive linter suite
- **Testing:** Standard `go test` with table-driven tests
- **Formatting:** `go fmt` / `gofmt` (enforced by golangci-lint)
- **Security:** `golangci-lint` includes gosec for security checks
- **Dependencies:** `go mod` for dependency management

### Docker

- **Linter:** `hadolint` for Dockerfile best practices - <https://github.com/hadolint/hadolint>
- **Base Images:** golang:1.25-alpine for build, Home Assistant base for runtime
- **Security:** Multi-stage builds, minimal layers, AppArmor profile

### Home Assistant Addon

- **Schema Validation:** YAML schema compliance with HA addon standards
- **Documentation:** DOCS.md, README.md, CHANGELOG.md
- **Config:** config.yaml with proper schema, build.json for build config

## Automation Rules

### On Go File Edit (`*.go`)

**Trigger:** `on.edit_file:*.go` OR `on.write:*.go`
**Actions:**

1. Run `go fmt` for formatting (automatic)
2. Run `golangci-lint run` for linting and security
3. Run `go test ./...` to ensure tests pass
4. Run `go mod tidy` to keep dependencies clean

### On Dockerfile Edit

**Trigger:** `on.edit_file:Dockerfile` OR `on.write:Dockerfile`
**Actions:**

1. Run `hadolint vevor-weatherbridge-go/Dockerfile` for best practices
2. Verify multi-stage build is optimal
3. Check for security issues (exposed secrets, running as root)

### On go.mod/go.sum Edit

**Trigger:** `on.edit_file:go.mod` OR `on.edit_file:go.sum`
**Actions:**

1. Run `go mod verify` to verify dependencies
2. Run `go mod tidy` to clean up
3. Check for known vulnerabilities in dependencies

### On Home Assistant Config Edit (`config.yaml`, `build.json`, `apparmor.txt`)

**Trigger:** `on.edit_file:config.yaml` OR similar
**Actions:**

1. Validate YAML syntax
2. Validate against HA addon schema
3. Invoke `/skills run ha-addon-skill --validate-config`

## Development Workflow

### Before Committing

1. Run all quality checks (linting, typing, security)
2. Ensure tests pass (if applicable)
3. Update documentation if API/config changes
4. Review generated hook suggestions

### CI/CD Expectations

- All Go code must pass `golangci-lint` with zero errors
- All tests must pass with `go test ./...`
- Hadolint must pass for Dockerfile
- Docker build must succeed
- Addon config must validate against HA schema

## Skills Available

### `/skills run ha-addon-skill`

Validates Home Assistant addon compliance:

- config.yaml schema validation
- build.json validation
- DOCS.md presence and format
- Icon and logo requirements

### `/skills run security-scan-skill`

Runs security analysis:

- SAST with semgrep
- Python security with bandit
- Dependency vulnerabilities (pip-audit)
- Secret detection (basic patterns)

### `./.claude/skills/version-management-skill/run.sh`

Manages semantic versioning:

- Validates current version format
- Bumps version (major/minor/patch)
- Updates config.yaml and CHANGELOG.md
- Ensures version consistency

Usage:

```bash
./.claude/skills/version-management-skill/run.sh check  # Validate current version
./.claude/skills/version-management-skill/run.sh patch  # Bump patch (0.1.2 -> 0.1.3)
./.claude/skills/version-management-skill/run.sh minor  # Bump minor (0.1.3 -> 0.2.0)
./.claude/skills/version-management-skill/run.sh major  # Bump major (0.2.0 -> 1.0.0)
```

## Token Efficiency Guidelines

### Use MCP References

Prefer `@repo:/path/to/file` over copying file contents in prompts.

### Use Haiku for Simple Tasks

- File searches, basic linting, formatting
- Simple validation tasks
- Hook execution

### Use Sonnet for Complex Tasks

- Refactoring with business logic changes
- Security remediation
- Architectural decisions

### Use Task Tool for Open-Ended Exploration

When searching for patterns or exploring codebase structure, use:

```text
/task explore "Find all MQTT publishing logic" --thoroughness medium
```

## Project-Specific Rules

### MQTT Publishing

- Always use Home Assistant MQTT discovery format
- Group all sensors under single device
- Use retain=True for config messages
- Validate MQTT topics follow pattern: `{prefix}/sensor/{device_id}_{sensor_id}/{topic}`

### Unit Conversion

- Support both metric and imperial via UNITS env var
- Always round to appropriate precision (temps: 1 decimal, pressure: 1 decimal, rain: 2 decimals)
- Use dedicated conversion functions (f_to_c, inhg_to_hpa, etc.)

### Error Handling

- Never expose internal errors to weather station (always return "success")
- Log errors for debugging but maintain service availability
- MQTT connection must be resilient

### Weather Underground Forwarding

- Optional feature via WU_FORWARD env var
- Use DNS resolution to bypass local DNS override
- Timeout after 5 seconds to avoid blocking main response

## Dependency Management

### Automated Updates with Renovate

This project uses [Renovate Bot](https://docs.renovatebot.com/) for automated dependency management:

- **Schedule**: Every weekend (Sunday)
- **Configuration**: `.github/renovate.json`
- **Dashboard**: Maintained as a GitHub issue titled "Dependency Dashboard 🤖"

### What Gets Updated

- **Python packages** (requirements.txt): Flask, paho-mqtt, pytz, dnspython, requests
- **Docker base images**: Home Assistant base-python images
- **GitHub Actions**: All workflow action versions

### Update Strategy

- **Patch/Minor Python updates**: Grouped into single PR
- **Major Python updates**: Separate PRs requiring review
- **Security vulnerabilities**: Immediate PRs with `security` label
- **Auto-merge**: Disabled - all updates require manual review

### Reviewing Dependency PRs

1. Check Dependency Review workflow results (security scan)
2. Review pip-audit output for vulnerabilities
3. Check for breaking changes in major updates
4. Test locally if uncertain
5. Merge and update addon version if needed

### Manual Dependency Updates

```bash
# Check outdated packages
pip list --outdated

# Update and freeze
pip install --upgrade package-name
pip freeze > vevor-weatherbridge/requirements.txt

# Security scan
pip-audit

# Bump version and update changelog
./.claude/skills/version-management-skill/run.sh patch
```

See [.github/DEPENDENCY_MANAGEMENT.md](.github/DEPENDENCY_MANAGEMENT.md) for complete documentation.

## Versioning and Release Workflow

### Semantic Versioning

This project follows [Semantic Versioning 2.0.0](https://semver.org/):

- **Major (X.0.0)**: Breaking changes, incompatible API changes
- **Minor (0.X.0)**: New features, backward compatible additions
- **Patch (0.0.X)**: Bug fixes, backward compatible fixes

### Single Source of Truth

- Version is defined ONLY in `vevor-weatherbridge/config.yaml`
- All other references derive from this file
- GitHub Actions workflows read version from config.yaml dynamically

### Release Process

1. **Make Changes**: Implement features/fixes in code
2. **Bump Version**: Use version management skill:

   ```bash
   ./.claude/skills/version-management-skill/run.sh patch  # or minor/major
   ```

3. **Update CHANGELOG**: Fill in the auto-generated changelog entry with:
   - Added: New features
   - Changed: Changes to existing functionality
   - Fixed: Bug fixes
   - Deprecated: Soon-to-be removed features
   - Removed: Removed features
   - Security: Security fixes
4. **Commit Changes**:

   ```bash
   git add vevor-weatherbridge/config.yaml vevor-weatherbridge/CHANGELOG.md
   git commit -m "Bump version to X.Y.Z"
   ```

5. **Push to GitHub**:

   ```bash
   git push origin main
   ```

6. **Automated Build & Release**:
   - GitHub Actions builds Docker images for all architectures
   - Images are tagged with version from config.yaml
   - GitHub release is created automatically with changelog excerpt
   - Home Assistant users can auto-update to new version

### Version-Related Files

- `vevor-weatherbridge/config.yaml` - Source of truth (version field)
- `vevor-weatherbridge/CHANGELOG.md` - Human-readable changelog
- `.github/workflows/build-addon.yml` - Builds and tags Docker images
- `.github/workflows/release.yml` - Creates GitHub releases

### Important Notes

- NEVER manually edit version in multiple places - use the skill
- ALWAYS update CHANGELOG.md before pushing
- GitHub Actions will fail if version already has a release
- Config.yaml version changes trigger automated releases

## Reference Links

- **Ruff:** <https://docs.astral.sh/ruff/>
- **Hadolint:** <https://github.com/hadolint/hadolint>
- **Semgrep:** <https://semgrep.dev/docs/>
- **Bandit:** <https://bandit.readthedocs.io/>
- **Claude Code Hooks:** <https://code.claude.com/docs/en/hooks>
- **Home Assistant Addon Dev:** <https://developers.home-assistant.io/docs/add-ons>
- **MQTT Discovery:** <https://www.home-assistant.io/integrations/mqtt/#mqtt-discovery>
- **Semantic Versioning:** <https://semver.org/>
