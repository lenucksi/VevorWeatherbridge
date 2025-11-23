# python-ci-skill

## Role
Provides comprehensive Python code quality automation for the VevorWeatherbridge project.

## Triggers
- Manual: `/skills run python-ci-skill`
- Hook-based: Automatically suggested when `*.py` files are edited
- Can be invoked with specific targets: `/skills run python-ci-skill --target <file>`

## Tools Required
Available via system installation or docker:
- `ruff` - Fast Python linter and formatter
- `mypy` - Static type checker
- `bandit` - Security issue detector
- `pip-audit` - Dependency vulnerability scanner

## Responsibilities

### 1. Linting and Formatting
- Run `ruff check --fix .` to identify and auto-fix linting issues
- Run `ruff format .` to ensure consistent code formatting
- Generate unified diff patches for suggested fixes
- Report unfixable issues with file:line references

### 2. Type Checking
- Run `mypy weatherstation.py --strict` for type safety
- Report type errors with context
- Suggest type annotations for untyped code

### 3. Security Scanning
- Run `bandit -r . -f json -ll` for Python security issues
- Focus on high-severity findings first
- Delegate comprehensive security scan to `security-scan-skill` if needed

### 4. Code Quality Metrics
- Report cyclomatic complexity for functions
- Identify code smells (long functions, too many arguments, etc.)
- Suggest refactoring opportunities

## Usage Examples

### Basic usage
```bash
/skills run python-ci-skill
```

### Target specific file
```bash
/skills run python-ci-skill --target weatherstation.py
```

### Generate patches only
```bash
/skills run python-ci-skill --mode generate-patches
```

### Security focus
```bash
/skills run python-ci-skill --focus security
```

## Output Format

JSON structure:
```json
{
  "linting": {
    "status": "passed|failed",
    "errors": 0,
    "warnings": 5,
    "fixed": 3,
    "details": [...]
  },
  "formatting": {
    "status": "passed|failed",
    "files_reformatted": 0
  },
  "type_checking": {
    "status": "passed|failed",
    "errors": [...],
    "coverage_percent": 85
  },
  "security": {
    "status": "passed|failed",
    "high_severity": 0,
    "medium_severity": 2,
    "findings": [...]
  },
  "patches": [
    "--- a/weatherstation.py\n+++ b/weatherstation.py\n..."
  ],
  "recommendations": [
    "Add type hints to safe_get function",
    "Consider extracting MQTT logic to separate module"
  ]
}
```

## Integration with Hooks

This skill is referenced in `.claude/hooks/python_quality_check.sh` which triggers on Python file edits. The hook provides a recommendation to run this skill, but does not execute it automatically for safety.

## Model Recommendation
Use **Haiku** for basic linting/formatting runs, **Sonnet** for complex refactoring suggestions.
