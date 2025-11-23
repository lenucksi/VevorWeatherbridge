# Claude Code Harness Architecture

## System Overview

```
┌─────────────────────────────────────────────────────────────┐
│                     Claude Code IDE                          │
│  ┌────────────────────────────────────────────────────────┐ │
│  │  Developer edits file (e.g., weatherstation.py)        │ │
│  └───────────────────────┬────────────────────────────────┘ │
│                          │                                   │
│                          ▼                                   │
│  ┌────────────────────────────────────────────────────────┐ │
│  │         .claude/settings.toml (Hook Config)            │ │
│  │   • on.edit_file:*.py → python_quality_check.sh        │ │
│  │   • on.edit_file:Dockerfile → dockerfile_check.sh      │ │
│  │   • on.edit_file:config.yaml → ha_config_check.sh      │ │
│  │   • on.edit_file:requirements.txt → security_check.sh  │ │
│  └───────────────────────┬────────────────────────────────┘ │
│                          │                                   │
│                          ▼                                   │
│  ┌────────────────────────────────────────────────────────┐ │
│  │           Hook Script Executes (Safe)                  │ │
│  │  • Displays recommendations only                       │ │
│  │  • Does NOT modify code                                │ │
│  │  • Suggests skill invocation                           │ │
│  └───────────────────────┬────────────────────────────────┘ │
└──────────────────────────┼────────────────────────────────────┘
                           │
              ┌────────────┴────────────┐
              │  Developer Decision     │
              └────────────┬────────────┘
                           │
         ┌─────────────────┼─────────────────┐
         │                 │                 │
         ▼                 ▼                 ▼
    Manual Command    Skill Invocation   Ignore
    (e.g., ruff)      (/skills run ...)
         │                 │
         │                 ▼
         │    ┌──────────────────────────┐
         │    │   Skill Dispatcher       │
         │    │  (Claude Code Agent)     │
         │    └────────┬─────────────────┘
         │             │
         │    ┌────────┼────────┐
         │    │        │        │
         │    ▼        ▼        ▼
         │  ┌────┐  ┌────┐  ┌────┐
         │  │ 🐍 │  │ 🏠 │  │ 🔒 │
         │  └────┘  └────┘  └────┘
         │  python   ha      security
         │    ci    addon     scan
         │  skill   skill    skill
         │    │       │        │
         │    └───────┼────────┘
         │            │
         │            ▼
         │  ┌──────────────────┐
         │  │  Tool Execution  │
         │  │  • ruff          │
         │  │  • mypy          │
         │  │  • bandit        │
         │  │  • semgrep       │
         │  │  • yamllint      │
         │  │  • hadolint      │
         │  │  • trivy         │
         │  │  • pip-audit     │
         │  └─────────┬────────┘
         │            │
         │            ▼
         │  ┌──────────────────┐
         │  │  Results         │
         │  │  (JSON + Text)   │
         │  └─────────┬────────┘
         │            │
         └────────────┼─────────
                      │
                      ▼
         ┌──────────────────────┐
         │  Developer Reviews   │
         │  and Applies Fixes   │
         └──────────────────────┘
```

## Component Details

### 1. Hook Layer (Event Detection)

**Purpose:** Detect file changes and recommend actions

**Characteristics:**
- Triggered automatically by Claude Code
- Read-only, non-destructive
- Fast execution (shell scripts)
- Zero token usage for detection

**Files:**
```
.claude/hooks/
├── python_quality_check.sh    # Python file edits
├── dockerfile_check.sh         # Dockerfile edits
├── ha_config_check.sh          # HA config edits
└── security_check.sh           # Dependency edits
```

### 2. Skill Layer (Analysis & Automation)

**Purpose:** Comprehensive quality analysis and recommendations

**Characteristics:**
- Manually invoked (safe)
- Token-efficient (use Haiku when possible)
- Modular and reusable
- JSON output for structured results

**Skills:**

```
python-ci-skill
├── Purpose: Python code quality
├── Tools: ruff, mypy, bandit
├── Model: Haiku (default), Sonnet (refactoring)
└── Output: Linting issues, type errors, security findings

ha-addon-skill
├── Purpose: HA addon compliance
├── Tools: yamllint, jq, hadolint
├── Model: Haiku (validation), Sonnet (generation)
└── Output: Missing files, config errors, recommendations

security-scan-skill
├── Purpose: Security & vulnerabilities
├── Tools: semgrep, bandit, pip-audit, trivy
├── Model: Haiku (scanning), Sonnet (remediation)
└── Output: Vulnerabilities, SAST findings, recommendations
```

### 3. Tool Layer (Execution)

**Purpose:** Actual quality checking tools

**Categories:**

**Python Quality:**
- `ruff` - Fast linter/formatter (2025 standard)
- `mypy` - Static type checker
- `black` - Alternative formatter (deprecated in favor of ruff)

**Security:**
- `bandit` - Python AST security scanner
- `semgrep` - Pattern-based SAST
- `pip-audit` - Dependency vulnerability scanner
- `trivy` - Container security scanner

**Config/Infrastructure:**
- `yamllint` - YAML linter
- `hadolint` - Dockerfile linter
- `jq` - JSON processor

## Data Flow

### Hook Trigger Flow

```
File Edit Event
    ↓
settings.toml matches pattern
    ↓
Execute hook script
    ↓
Display recommendations (stdout)
    ↓
No further action (waits for human)
```

### Skill Invocation Flow

```
Developer command: /skills run <skill-name>
    ↓
Claude Code resolves skill path
    ↓
Execute skill run.sh with arguments
    ↓
Skill checks tool availability
    ↓
For each tool:
  - Run with appropriate flags
  - Capture output
  - Parse results
    ↓
Aggregate results
    ↓
Format as JSON + human-readable text
    ↓
Return to Claude Code
    ↓
Display to developer
    ↓
Developer reviews and acts
```

## Token Efficiency Strategy

### Model Selection Decision Tree

```
Task Required?
    ├─ File search / pattern matching
    │  └─> Use Task tool (Explore agent) with Haiku
    │
    ├─ Linting / formatting check
    │  └─> Run skill directly (Haiku)
    │
    ├─ Type checking
    │  └─> Run skill directly (Haiku)
    │
    ├─ Security scan (detection only)
    │  └─> Run skill directly (Haiku)
    │
    ├─ Simple bug fix
    │  └─> Sonnet
    │
    ├─ Refactoring suggestion
    │  └─> Sonnet
    │
    ├─ Security remediation
    │  └─> Sonnet
    │
    └─ Architectural decision
       └─> Sonnet
```

### MCP Usage

**When to use MCP references:**

```python
# Instead of this (high token cost):
Read weatherstation.py → 200 lines → ~800 tokens
Process content
Analyze

# Use this (low token cost):
Reference @repo:/weatherstation.py → ~50 tokens
Tool reads directly
Analyze
```

**Benefit:**
- 90%+ token reduction for file operations
- Faster execution
- Exact file content (no copy-paste errors)

## Safety Mechanisms

### Multi-Layer Safety

```
Layer 1: Hook Scripts
    → Read-only
    → No code modification
    → Display recommendations only

Layer 2: Skill Scripts
    → Analysis only
    → Generate patches as text
    → No automatic application

Layer 3: Tool Execution
    → Isolated in skill context
    → Results captured and parsed
    → No direct file writes

Layer 4: Human Review
    → Developer reviews findings
    → Developer applies fixes manually
    → Full control and visibility
```

### Generated Artifacts

```
.claude/generated-hooks/
    ├── python-edit.hook.sh         # Generated but not executed
    ├── dockerfile-edit.hook.sh     # Safe stubs for review
    └── ...

→ gitignored (not committed)
→ Reviewed before execution
→ Safe by default
```

## Integration Points

### Claude Code Integration

```
Claude Code Editor
    ↓
    ├─> Watches file edits (via LSP/filesystem)
    ├─> Triggers hooks (via settings.toml)
    ├─> Provides /skills command
    └─> Supports MCP references
```

### External Tool Integration

```
Skills
    ↓
    ├─> Checks if tool available (command -v)
    ├─> Skips gracefully if missing
    ├─> Shows installation instructions
    └─> Continues with available tools
```

## Extensibility

### Adding New Hook

```bash
# 1. Create hook script
cat > .claude/hooks/new_hook.sh << 'EOF'
#!/usr/bin/env bash
echo "Recommendation for this file type..."
