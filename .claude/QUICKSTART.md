# Claude Code Harness - Quick Start Guide

## Installation

### 1. Install Required Tools

```bash
# Python quality tools (recommended)
pip install ruff mypy bandit pip-audit semgrep

# YAML/JSON tools (for HA addon validation)
pip install yamllint
sudo apt install jq  # or: brew install jq

# Docker linting (optional but recommended)
# hadolint: https://github.com/hadolint/hadolint#install
wget -O /usr/local/bin/hadolint https://github.com/hadolint/hadolint/releases/latest/download/hadolint-Linux-x86_64
chmod +x /usr/local/bin/hadolint

# Container security (optional)
# trivy: https://aquasecurity.github.io/trivy/latest/getting-started/installation/
```

### 2. Verify Harness Installation

```bash
# Check structure
ls -la .claude/

# Test hooks are executable
ls -lh .claude/hooks/*.sh
ls -lh .claude/skills/*/run.sh

# All should show 'x' permission (executable)
```

## Quick Command Reference

### Manual Skill Invocation

```bash
# Python quality check
./.claude/skills/python-ci-skill/run.sh

# HA addon validation
./.claude/skills/ha-addon-skill/run.sh

# Security scan
./.claude/skills/security-scan-skill/run.sh
```

### Via Claude Code (preferred)

```bash
# Python CI
/skills run python-ci-skill

# HA addon check
/skills run ha-addon-skill

# Security scan
/skills run security-scan-skill

# Security scan - dependencies only
/skills run security-scan-skill dependencies
```

## Common Workflows

### Workflow 1: Before Committing Code

```bash
# 1. Run Python quality checks
/skills run python-ci-skill

# 2. Run security scan
/skills run security-scan-skill

# 3. Apply any recommended fixes
ruff check --fix .
ruff format .

# 4. Re-run checks to verify
/skills run python-ci-skill

# 5. Commit if all clear
```

### Workflow 2: Setting Up HA Addon

```bash
# 1. Check current structure
/skills run ha-addon-skill

# 2. Note missing files from output

# 3. Create required files (Claude can help):
# - config.yaml
# - build.json
# - DOCS.md
# - icon.png
# - run.sh

# 4. Validate again
/skills run ha-addon-skill
```

### Workflow 3: Security Audit

```bash
# 1. Full security scan
/skills run security-scan-skill

# 2. Focus on dependencies
/skills run security-scan-skill dependencies

# 3. Check Docker security
/skills run security-scan-skill container

# 4. Look for secrets
/skills run security-scan-skill secrets
```

## What Happens Automatically?

### When You Edit Python Files

Hook triggers automatically and displays:

```text
📋 Python file modified - Quality checks recommended:

Recommended actions:
  1. Run ruff linting:    ruff check --fix .
  2. Run type checking:   mypy weatherstation.py
  3. Run security scan:   bandit -r . -f json
  4. Check formatting:    ruff format --check .

Or invoke the python-ci-skill:
  /skills run python-ci-skill
```

**Note:** Nothing runs automatically - you must invoke the commands or skills manually.

### When You Edit Dockerfile

Hook displays Dockerfile quality recommendations.

### When You Edit requirements.txt

Hook displays security scan recommendations.

## Understanding the Output

### Python CI Skill Output

```text
🐍 Python CI Skill - Running quality checks
==================================================

1️⃣  Running Ruff linter...
   [Shows linting issues and auto-fixable items]

2️⃣  Checking code formatting...
   [Shows formatting differences]

3️⃣  Running type checker...
   [Shows type errors]

4️⃣  Running security scan...
   [Shows security issues]

📝 Recommendations:
   - Review any issues reported above
   - Run 'ruff check --fix .' to auto-fix
```

### HA Addon Skill Output

```text
🏠 Home Assistant Addon Skill
==================================================

📁 Checking addon structure...

Present files:
  ✅ Dockerfile - Required Docker build file
  ✅ README.md - Developer documentation

Missing files (needed for HA addon):
  ❌ config.yaml - Required addon configuration
  ❌ build.json - Required build configuration
  ...
```

### Security Scan Skill Output

```text
🔒 Security Scan Skill
==================================================

1️⃣  Running Bandit (Python security scanner)...
   [Shows Python security issues]

2️⃣  Running Semgrep (SAST)...
   [Shows SAST findings]

3️⃣  Scanning dependencies for vulnerabilities...
   [Shows vulnerable packages]

📋 Security Summary
   [High-level recommendations]
```

## Troubleshooting

### Hook Not Triggering

1. Check [.claude/settings.toml](.claude/settings.toml) exists
2. Verify hook scripts are executable: `chmod +x .claude/hooks/*.sh`
3. Restart Claude Code if needed

### Tool Not Found Errors

Install missing tools:

```bash
pip install ruff mypy bandit pip-audit semgrep yamllint
```

Skills will show which tools are missing and skip those checks gracefully.

### Permission Denied

Make scripts executable:

```bash
chmod +x .claude/hooks/*.sh
chmod +x .claude/skills/*/run.sh
```

## Advanced Usage

### Targeting Specific Files

```bash
# Check only one file
./.claude/skills/python-ci-skill/run.sh weatherstation.py
```

### Custom Tool Options

Edit skill `run.sh` files to customize tool parameters.

### Adding Custom Rules

For Semgrep, add custom rules:

```bash
# Create .semgrep.yml in project root
# See: https://semgrep.dev/docs/writing-rules/overview/
```

For Ruff, add configuration:

```bash
# Create ruff.toml or pyproject.toml
# See: https://docs.astral.sh/ruff/configuration/
```

## Next Steps

1. **Install recommended tools** (see Installation above)
2. **Run initial audit:**

   ```bash
   /skills run python-ci-skill
   /skills run ha-addon-skill
   /skills run security-scan-skill
   ```

3. **Review and fix issues** identified
4. **Read full documentation:**
   - [MANIFEST.md](MANIFEST.md) - Complete harness documentation
   - [CLAUDE.md](../CLAUDE.md) - Project rules and standards
   - Individual skill docs in `.claude/skills/*/SKILL.md`

## Getting Help

- **Harness usage:** Read [MANIFEST.md](MANIFEST.md)
- **Project rules:** Read [CLAUDE.md](../CLAUDE.md)
- **Claude Code:** <https://code.claude.com/docs/>
- **Tools:** See individual tool documentation links in skill SKILL.md files

---

**Pro Tip:** Use hooks as reminders, skills as automation. The harness doesn't change code without your approval - it's your quality assurance assistant, not an auto-fixer.
