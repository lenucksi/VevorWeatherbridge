# Claude Code Harness Manifest - VevorWeatherbridge

## Overview

This manifest describes the Claude Code automation harness for the VevorWeatherbridge project. The harness provides quality assurance, security scanning, and Home Assistant addon development assistance through hooks, skills, and agents.

## Purpose

The harness automates and orchestrates development quality checks without executing them automatically. Instead, it:

1. Detects code/config changes via hooks
2. Recommends appropriate quality checks
3. Provides skills that can be manually invoked for comprehensive analysis
4. Ensures token-efficient operations by delegating to appropriate Claude models

## Directory Structure

```text
.claude/
├── settings.toml                 # Hook configuration
├── MANIFEST.md                   # This file
├── hooks/                        # Hook scripts (triggered on file edits)
│   ├── python_quality_check.sh
│   ├── dockerfile_check.sh
│   ├── ha_config_check.sh
│   └── security_check.sh
├── generated-hooks/              # Generated hook artifacts (safe stubs)
├── skills/                       # Reusable skill modules
│   ├── python-ci-skill/
│   │   ├── SKILL.md
│   │   └── run.sh
│   ├── ha-addon-skill/
│   │   ├── SKILL.md
│   │   └── run.sh
│   └── security-scan-skill/
│       ├── SKILL.md
│       └── run.sh
└── prompts/                      # Prompt templates (future use)
```

## Workflow

### Typical Developer Flow

```text
Developer edits Python file (e.g., weatherstation.py)
              ↓
Hook detects edit (via settings.toml: on.edit_file:*.py)
              ↓
Hook script runs (.claude/hooks/python_quality_check.sh)
              ↓
Displays recommended actions:
  - ruff check --fix .
  - mypy weatherstation.py
  - bandit -r .
  - Or: /skills run python-ci-skill
              ↓
Developer decides to invoke skill manually
              ↓
Claude runs python-ci-skill (uses Haiku for efficiency)
              ↓
Results returned as JSON + recommendations
              ↓
Developer reviews and applies fixes
```

## Hooks

Hooks are defined in [.claude/settings.toml](.claude/settings.toml) and trigger on specific file operations.

### Available Hooks

| Hook Trigger | Script | Purpose |
|--------------|--------|---------|
| `on.edit_file:*.py` | `python_quality_check.sh` | Recommends linting, typing, security checks |
| `on.write:*.py` | `python_quality_check.sh` | Same as above for new files |
| `on.edit_file:Dockerfile` | `dockerfile_check.sh` | Recommends Dockerfile linting |
| `on.edit_file:config.yaml` | `ha_config_check.sh` | Recommends HA config validation |
| `on.edit_file:requirements.txt` | `security_check.sh` | Recommends dependency security scan |

### Hook Behavior

**Important:** Hooks in this project are **non-executing**. They:

- Display recommended commands
- Do NOT run linters/formatters automatically
- Suggest skill invocations
- Maintain safety by requiring manual confirmation

## Skills

Skills are reusable, encapsulated automation modules that Claude can invoke.

### 1. python-ci-skill

**Location:** [.claude/skills/python-ci-skill/](.claude/skills/python-ci-skill/)

**Purpose:** Comprehensive Python code quality checks

**Tools Used:**

- `ruff` - Linting and formatting
- `mypy` - Type checking
- `bandit` - Security scanning

**Invocation:**

```bash
# Basic usage
/skills run python-ci-skill

# Target specific file
/skills run python-ci-skill --target weatherstation.py

# Security focus
/skills run python-ci-skill --focus security
```

**Model:** Haiku for basic runs, Sonnet for refactoring suggestions

**Documentation:** [.claude/skills/python-ci-skill/SKILL.md](.claude/skills/python-ci-skill/SKILL.md)

---

### 2. ha-addon-skill

**Location:** [.claude/skills/ha-addon-skill/](.claude/skills/ha-addon-skill/)

**Purpose:** Home Assistant addon structure validation and compliance

**Tools Used:**

- `yamllint` - YAML validation
- `jq` - JSON validation
- `hadolint` - Dockerfile linting
- HA addon schema knowledge

**Invocation:**

```bash
# Check addon structure
/skills run ha-addon-skill

# Validate config only
/skills run ha-addon-skill --validate-config

# Generate template
/skills run ha-addon-skill --generate-template
```

**Model:** Haiku for validation, Sonnet for template generation

**Documentation:** [.claude/skills/ha-addon-skill/SKILL.md](.claude/skills/ha-addon-skill/SKILL.md)

---

### 3. security-scan-skill

**Location:** [.claude/skills/security-scan-skill/](.claude/skills/security-scan-skill/)

**Purpose:** Security analysis and vulnerability detection

**Tools Used:**

- `semgrep` - SAST pattern matching
- `bandit` - Python security
- `pip-audit` - Dependency vulnerabilities
- `trivy` - Container scanning

**Invocation:**

```bash
# Full security scan
/skills run security-scan-skill

# Dependencies only
/skills run security-scan-skill --focus dependencies

# Container security
/skills run security-scan-skill --focus container

# Secret detection
/skills run security-scan-skill --focus secrets
```

**Model:** Haiku for scanning, Sonnet for remediation

**Documentation:** [.claude/skills/security-scan-skill/SKILL.md](.claude/skills/security-scan-skill/SKILL.md)

## Usage Examples

### Scenario 1: Editing Python Code

```bash
# 1. Edit weatherstation.py
# 2. Hook displays recommendations
# 3. Invoke skill
/skills run python-ci-skill

# 4. Review output and apply fixes
ruff check --fix .
ruff format .
```

### Scenario 2: Creating HA Addon

```bash
# 1. Check current structure
/skills run ha-addon-skill

# 2. Review missing files
# 3. Create required files (config.yaml, build.json, etc.)

# 4. Validate again
/skills run ha-addon-skill --validate-config
```

### Scenario 3: Security Review

```bash
# 1. Full security audit
/skills run security-scan-skill

# 2. Check dependencies specifically
/skills run security-scan-skill --focus dependencies

# 3. Review Docker image
/skills run security-scan-skill --focus container
```

### Scenario 4: Pre-Commit Checks

```bash
# Run all quality checks before committing
/skills run python-ci-skill
/skills run security-scan-skill
/skills run ha-addon-skill

# Fix any issues
# Commit changes
```

## MCP Integration

### Using MCP References

When invoking skills or working with large files, use MCP references to avoid token waste:

```bash
# Instead of pasting file content
/skills run python-ci-skill --target @repo:/weatherstation.py

# Reference multiple files
/skills run security-scan-skill --files @repo:/weatherstation.py,@repo:/Dockerfile
```

### Benefits

- Reduced token usage
- Faster processing
- Access to exact file content
- No manual copy-paste

## Token Efficiency Strategy

### Model Selection

| Task Type | Recommended Model | Rationale |
|-----------|------------------|-----------|
| Linting/Formatting | Haiku | Fast, deterministic |
| Config validation | Haiku | Schema-based, simple |
| Security scanning | Haiku | Pattern matching |
| Refactoring suggestions | Sonnet | Complex reasoning |
| Architecture decisions | Sonnet | Strategic thinking |
| Documentation generation | Sonnet | Quality writing |

### Best Practices

1. **Use Task Tool for Exploration**

   ```bash
   # Instead of multiple Grep/Glob calls
   /task explore "Find all MQTT publishing logic" --thoroughness medium
   ```

2. **Batch Operations**
   - Run multiple checks in single skill invocation
   - Use `--focus all` rather than multiple separate runs

3. **Incremental Checks**
   - Use `--target specific_file.py` when only one file changed
   - Avoid full-project scans if unnecessary

## Tool Installation

### Required Tools

Install recommended tools for full harness functionality:

```bash
# Python quality tools
pip install ruff mypy bandit pip-audit

# Security tools
pip install semgrep safety

# HA addon tools
pip install yamllint
sudo apt install jq

# Docker tools
# hadolint: https://github.com/hadolint/hadolint#install
# trivy: https://aquasecurity.github.io/trivy/latest/getting-started/installation/
```

### Optional Tools

```bash
# Additional security
pip install trufflehog

# Docker alternative scanners
# grype: https://github.com/anchore/grype#installation
```

## Safety Notes

### Hook Safety

- All hooks are **read-only and recommendation-based**
- No automatic code modifications
- No automatic command execution
- Generated artifacts stored in `.claude/generated-hooks/` for review

### Skill Safety

- Skills run analysis and return results
- Code changes require manual application
- Patches are generated as diffs for review
- No side effects without explicit confirmation

## Project-Specific Considerations

### VevorWeatherbridge Specifics

1. **MQTT Focus**
   - Validate MQTT topic structure
   - Ensure HA auto-discovery format compliance
   - Check for proper device grouping

2. **Unit Conversion**
   - Verify metric/imperial conversion accuracy
   - Check rounding precision

3. **Weather Underground Forwarding**
   - Validate DNS resolution security
   - Check timeout handling

4. **Docker Deployment**
   - Ensure multi-architecture support
   - Validate environment variable usage
   - Check port exposure

## Continuous Improvement

### Extending the Harness

To add new skills:

1. Create skill directory: `.claude/skills/my-new-skill/`
2. Write `SKILL.md` documentation
3. Create `run.sh` executable script
4. Add hook trigger in `settings.toml` if needed
5. Update this manifest

### Feedback Loop

The harness evolves based on:

- Actual findings from tools
- Developer pain points
- New best practices
- Tool updates

## Reference Documentation

### Official Docs

- **Claude Code Hooks:** <https://code.claude.com/docs/en/hooks>
- **Claude Code Skills:** <https://code.claude.com/docs/en/skills>
- **Home Assistant Addons:** <https://developers.home-assistant.io/docs/add-ons>

### Tool Documentation

- **Ruff:** <https://docs.astral.sh/ruff/>
- **Hadolint:** <https://github.com/hadolint/hadolint>
- **Semgrep:** <https://semgrep.dev/docs/>
- **Bandit:** <https://bandit.readthedocs.io/>
- **pip-audit:** <https://github.com/pypa/pip-audit>

### Project Docs

- **Main Rules:** [CLAUDE.md](../CLAUDE.md)
- **README:** [README.md](../README.md)

## Support

For issues with:

- **Harness itself:** Review this manifest and skill documentation
- **Claude Code:** <https://code.claude.com/docs/>
- **Tools:** Refer to individual tool documentation above
- **Project code:** See README.md

---

**Last Updated:** 2025-11-10
**Harness Version:** 1.0.0
**Compatible with:** Claude Code latest, Sonnet 4.5, Haiku 3.5
