#!/usr/bin/env bash
# Post Edit/Write Dispatcher Hook
# Routes to appropriate check based on file type

set -euo pipefail

FILE_PATH="${1:-}"

# If no file path provided, exit silently
[ -z "$FILE_PATH" ] && exit 0

# Get the basename and extension
FILENAME=$(basename "$FILE_PATH")
EXTENSION="${FILENAME##*.}"

# Get script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Route based on file type
case "$FILENAME" in
  *.py)
    "$SCRIPT_DIR/python_quality_check.sh"
    ;;
  Dockerfile)
    "$SCRIPT_DIR/dockerfile_check.sh"
    ;;
  config.yaml|build.json)
    "$SCRIPT_DIR/ha_config_check.sh"
    ;;
  requirements.txt|pyproject.toml|poetry.lock)
    "$SCRIPT_DIR/security_check.sh"
    ;;
  *)
    # Unknown file type, exit silently
    exit 0
    ;;
esac
