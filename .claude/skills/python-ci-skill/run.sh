#!/usr/bin/env bash
# Python CI Skill Runner
# This script orchestrates Python quality checks using Poetry environment

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/../../.." && pwd)"

cd "$PROJECT_ROOT"

TARGET="${1:-.}"
MODE="${2:-full}"

echo "🐍 Python CI Skill - Running quality checks on: $TARGET"
echo "=================================================="
echo ""

# Use Poetry to run tools
POETRY_RUN="poetry run"

# Check if Poetry is available
if ! command -v poetry >/dev/null 2>&1; then
    echo "⚠️  Poetry not found - falling back to system tools"
    POETRY_RUN=""
fi

# 1. Linting with Ruff
echo "1️⃣  Running Ruff linter..."
if $POETRY_RUN ruff check "$TARGET" --output-format=text 2>/dev/null || true; then
    echo ""
    echo "   Auto-fixable issues (dry-run):"
    $POETRY_RUN ruff check "$TARGET" --fix --diff 2>/dev/null || true
else
    echo "   ⚠️  Ruff failed or not available"
fi
echo ""

# 2. Formatting check
echo "2️⃣  Checking code formatting..."
$POETRY_RUN ruff format "$TARGET" --check --diff 2>/dev/null || echo "   ⚠️  Format check failed or not available"
echo ""

# 3. Type checking
echo "3️⃣  Running type checker..."
$POETRY_RUN mypy weatherstation.py --no-error-summary 2>/dev/null || echo "   ⚠️  MyPy failed or not available"
echo ""

# 4. Security scan (basic)
echo "4️⃣  Running security scan..."
$POETRY_RUN bandit -r . -ll -f screen --skip B404,B603 --exclude ./.venv,./venv 2>/dev/null || echo "   ⚠️  Bandit failed or not available"
echo ""

# Summary
echo "=================================================="
echo "✅ Python CI Skill - Check complete"
echo ""
echo "📝 Recommendations:"
echo "   - Review any issues reported above"
echo "   - Run 'poetry run ruff check --fix .' to auto-fix linting issues"
echo "   - Run 'poetry run ruff format .' to format code"
echo "   - For security deep-dive: /skills run security-scan-skill"
echo ""
