#!/usr/bin/env bash
# ==============================================================================
# Post-Edit Hook for Python Files
# Runs quality checks and tests after editing Python files
# ==============================================================================

set -e

EDITED_FILE="$1"

# Only run for Python files in vevor-weatherbridge/
if [[ ! "$EDITED_FILE" =~ vevor-weatherbridge/.*\.py$ ]]; then
    exit 0
fi

# Skip test files themselves
if [[ "$EDITED_FILE" =~ test_.*\.py$ ]]; then
    echo "⏭️  Skipping hook for test file"
    exit 0
fi

echo "🔍 Running post-edit checks for $EDITED_FILE..."

# Navigate to addon directory
ADDON_DIR="$(dirname "$EDITED_FILE")"
cd "$ADDON_DIR"

# Run ruff linter
echo "  ✓ Linting with ruff..."
poetry run ruff check "$EDITED_FILE" --fix || true

# Run ruff formatter
echo "  ✓ Formatting with ruff..."
poetry run ruff format "$EDITED_FILE"

# Run full test suite
echo "  ✓ Running test suite..."
poetry run pytest test_weatherstation.py test_run_sh.py -v --tb=short -q || {
    echo ""
    echo "❌ Tests failed! Please fix before committing."
    exit 1
}

echo ""
echo "✅ All checks passed!"
