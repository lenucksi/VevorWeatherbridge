#!/usr/bin/env bash
# Python Quality Check Hook
# Triggered on Python file edits - recommends quality checks, does NOT auto-execute

set -euo pipefail

echo "📋 Python file modified - Quality checks recommended:"
echo ""
echo "Recommended actions:"
echo "  1. Run ruff linting:    poetry run ruff check --fix ."
echo "  2. Run type checking:   poetry run mypy weatherstation.py"
echo "  3. Run security scan:   poetry run bandit -r . -f json"
echo "  4. Check formatting:    poetry run ruff format --check ."
echo ""
echo "Or invoke the python-ci-skill:"
echo "  ./.claude/skills/python-ci-skill/run.sh"
echo ""
