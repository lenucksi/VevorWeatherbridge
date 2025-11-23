#!/usr/bin/env bash
# Dockerfile Quality Check Hook
# Triggered on Dockerfile edits - recommends hadolint check

set -euo pipefail

echo "🐳 Dockerfile modified - Quality checks recommended:"
echo ""
echo "Recommended actions:"
echo "  1. Run hadolint:        hadolint Dockerfile"
echo "  2. Check build:         docker build -t weatherbridge-test ."
echo "  3. Security scan:       docker scan weatherbridge-test"
echo ""
