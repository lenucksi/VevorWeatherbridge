#!/usr/bin/env bash
# SPDX-License-Identifier: GPL-3.0-or-later
# Check for latest Home Assistant base image version on GHCR.
# Usage: ./scripts/check-ha-base-version.sh [current_version]
#   current_version defaults to "3.23" if not provided.
#   Outputs: version=X, newer=true/false

set -euo pipefail

CURRENT="${1:-3.23}"

RESPONSE=$(curl -s -w "\n%{http_code}" "https://ghcr.io/v2/home-assistant/base/tags/list")
HTTP_CODE=$(echo "$RESPONSE" | tail -1)
BODY=$(echo "$RESPONSE" | sed '$d')

if [ "$HTTP_CODE" != "200" ]; then
  echo "Warning: GHCR API returned HTTP ${HTTP_CODE} (auth required?)" >&2
  echo "newer=false"
  echo "version=${CURRENT}"
  exit 0
fi

LATEST=$(echo "$BODY" | jq -r '.tags // [] | .[]' \
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
