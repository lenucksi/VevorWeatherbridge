#!/usr/bin/env bash
# Home Assistant Config Check Hook
# Triggered on config.yaml edits

set -euo pipefail

echo "🏠 Home Assistant config modified - Validation recommended:"
echo ""
echo "Recommended actions:"
echo "  1. Validate YAML:       yamllint config.yaml"
echo "  2. Check HA schema:     /skills run ha-addon-skill --validate-config"
echo ""
