#!/usr/bin/env bash
# Security Check Hook
# Triggered on requirements.txt edits

set -euo pipefail

echo "🔒 Dependencies modified - Security scan recommended:"
echo ""
echo "Recommended actions:"
echo "  1. Check for vulnerabilities:  pip-audit"
echo "  2. Run security scan:          /skills run security-scan-skill"
echo "  3. Check license compliance:   pip-licenses"
echo ""
