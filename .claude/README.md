# Claude Code Harness for VevorWeatherbridge

This directory contains the Claude Code automation harness - a quality assurance and development assistance framework.

## What's Inside

```text
.claude/
├── README.md              # This file
├── QUICKSTART.md          # Quick reference guide
├── MANIFEST.md            # Complete harness documentation
├── settings.toml          # Hook configuration
├── hooks/                 # Event-triggered scripts
├── skills/                # Reusable automation modules
├── generated-hooks/       # Generated artifacts (gitignored)
└── prompts/               # Prompt templates (future use)
```

## Quick Links

- **Getting Started:** Read [QUICKSTART.md](QUICKSTART.md)
- **Full Documentation:** Read [MANIFEST.md](MANIFEST.md)
- **Project Rules:** Read [../CLAUDE.md](../CLAUDE.md)

## What This Harness Does

### 🎯 Purpose

Provides automated quality assurance recommendations through:

1. **Hooks** - Detect code changes and suggest checks
2. **Skills** - Run comprehensive quality analysis
3. **Safety** - Never modifies code without explicit approval

### 🔧 Available Skills

| Skill | Purpose | Command |
|-------|---------|---------|
| **python-ci-skill** | Python quality checks (linting, typing, security) | `/skills run python-ci-skill` |
| **ha-addon-skill** | Home Assistant addon validation | `/skills run ha-addon-skill` |
| **security-scan-skill** | Security and vulnerability scanning | `/skills run security-scan-skill` |

### 🪝 Active Hooks

| Trigger | Action |
|---------|--------|
| Edit `*.py` | Recommends Python quality checks |
| Edit `Dockerfile` | Recommends Dockerfile linting |
| Edit `config.yaml` | Recommends HA config validation |
| Edit `requirements.txt` | Recommends security scan |

## First-Time Setup

### 1. Install Tools (Recommended)

```bash
# Python quality tools
pip install ruff mypy bandit pip-audit semgrep

# Config validation
pip install yamllint
sudo apt install jq

# Docker linting (optional)
# See: https://github.com/hadolint/hadolint#install
```

### 2. Test the Harness

```bash
# Run each skill to verify setup
./.claude/skills/python-ci-skill/run.sh
./.claude/skills/ha-addon-skill/run.sh
./.claude/skills/security-scan-skill/run.sh
```

### 3. Verify Hooks Work

```bash
# Check hook configuration
cat .claude/settings.toml

# Hooks trigger automatically when you edit files in Claude Code
```

## Daily Usage

### Before Committing

```bash
# Quick quality check
/skills run python-ci-skill

# Apply any fixes
ruff check --fix .
ruff format .

# Verify security
/skills run security-scan-skill
```

### Developing HA Addon

```bash
# Check addon structure
/skills run ha-addon-skill

# Create missing files as recommended
# Validate again
```

### Security Audit

```bash
# Full scan
/skills run security-scan-skill

# Or targeted scans
/skills run security-scan-skill dependencies
/skills run security-scan-skill container
```

## How It Works

### Hook Workflow

```text
You edit file → Hook detects → Shows recommendations → You decide to act
```

**Example:**

```bash
# You edit weatherstation.py
# Claude Code hook triggers automatically and shows:

📋 Python file modified - Quality checks recommended:
  1. Run ruff linting:    ruff check --fix .
  2. Run type checking:   mypy weatherstation.py
  ...
  Or invoke: /skills run python-ci-skill
```

### Skill Workflow

```text
You invoke skill → Skill runs tools → Returns results → You review & apply
```

**Example:**

```bash
# You run: /skills run python-ci-skill
# Skill executes: ruff, mypy, bandit
# Returns: JSON with findings + recommendations
# You apply fixes manually
```

## Safety Principles

✅ **Safe:**

- Hooks display recommendations only
- Skills analyze and report findings
- You review before applying changes

❌ **Never Does:**

- Auto-modify code without approval
- Execute destructive commands
- Change files automatically

## Token Efficiency

The harness is designed for token efficiency:

- **Hooks:** Lightweight shell scripts, minimal tokens
- **Skills:** Use Haiku for simple tasks, Sonnet for complex reasoning
- **MCP:** Can reference files with `@repo:/path` instead of copying content
- **Batching:** Skills run multiple checks in one invocation

## Troubleshooting

### Tools Not Found

Skills gracefully handle missing tools and show installation instructions.

### Hooks Not Triggering

1. Verify `.claude/settings.toml` exists
2. Check hooks are executable: `ls -lh .claude/hooks/`
3. Restart Claude Code if needed

### Permission Errors

```bash
chmod +x .claude/hooks/*.sh
chmod +x .claude/skills/*/run.sh
```

## Customization

### Modify Hook Behavior

Edit `.claude/settings.toml` to add/remove hook triggers.

### Modify Skill Behavior

Edit individual `run.sh` scripts in skill directories.

### Add Custom Rules

Create tool-specific config files:

- `ruff.toml` - Ruff configuration
- `.semgrep.yml` - Custom Semgrep rules
- `pyproject.toml` - Python project metadata

## Documentation

| Document | Purpose |
|----------|---------|
| [README.md](README.md) | This file - overview |
| [QUICKSTART.md](QUICKSTART.md) | Quick reference commands |
| [MANIFEST.md](MANIFEST.md) | Complete technical documentation |
| [../CLAUDE.md](../CLAUDE.md) | Project coding standards |
| `skills/*/SKILL.md` | Individual skill documentation |

## Version

- **Harness Version:** 1.0.0
- **Created:** 2025-11-10
- **Compatible With:** Claude Code (latest), Sonnet 4.5, Haiku 3.5

## Support

- **Harness Issues:** Review documentation above
- **Claude Code:** <https://code.claude.com/docs/>
- **Tool Docs:** See individual skill SKILL.md files
- **Project:** See [../README.md](../README.md)

---

**Remember:** This harness is your assistant, not an auto-pilot. It recommends, you decide. 🚀
