#!/usr/bin/env bash
# SPDX-License-Identifier: GPL-3.0-or-later
# Check for latest Home Assistant base image version on GHCR.
# Usage: ./scripts/check-ha-base-version.sh [current_version]
#   current_version defaults to "3.23" if not provided.
#   Outputs: version=X, newer=true/false

set -euo pipefail

CURRENT="${1:-3.23}"

LATEST=$(curl -s "https://ghcr.io/v2/home-assistant/base/tags/list" \
  | jq -r '.tags[]' \
  | grep -E '^3\.' \
  | sort -V \
  | tail -1)

echo "Current HA base version: ${CURRENT}"
echo "Latest HA base version:  ${LATEST:-unknown}"

if [ -z "$LATEST" ]; then
  echo "newer=false"
  echo "version=${CURRENT}"
  exit 0
fi

echo "version=${LATEST}"

if [ "${LATEST}" != "${CURRENT}" ]; then
  echo "newer=true"
  echo "UPGRADE AVAILABLE: ${CURRENT} -> ${LATEST}"
else
  echo "newer=false"
fi
