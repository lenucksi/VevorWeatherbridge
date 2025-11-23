#!/usr/bin/env bash
# Home Assistant Addon Skill Runner
# Validates HA addon structure and configuration

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/../../.." && pwd)"

cd "$PROJECT_ROOT"

MODE="${1:---check-structure}"

echo "🏠 Home Assistant Addon Skill"
echo "=================================================="
echo ""

# Check for addon files
echo "📁 Checking addon structure..."
echo ""

FILES_TO_CHECK=(
    "config.yaml:Required addon configuration"
    "build.json:Required build configuration"
    "Dockerfile:Required Docker build file"
    "DOCS.md:User documentation"
    "README.md:Developer documentation"
    "CHANGELOG.md:Version history"
    "icon.png:Addon icon (256x256)"
    "logo.png:Addon logo (optional)"
    "run.sh:Startup script"
)

MISSING=()
PRESENT=()

for item in "${FILES_TO_CHECK[@]}"; do
    IFS=':' read -r file desc <<< "$item"
    if [ -f "$file" ]; then
        PRESENT+=("✅ $file - $desc")
    else
        MISSING+=("❌ $file - $desc")
    fi
done

echo "Present files:"
for item in "${PRESENT[@]}"; do
    echo "  $item"
done
echo ""

if [ ${#MISSING[@]} -gt 0 ]; then
    echo "Missing files (needed for HA addon):"
    for item in "${MISSING[@]}"; do
        echo "  $item"
    done
    echo ""
fi

# Validate config.yaml if it exists
if [ -f "config.yaml" ]; then
    echo "🔍 Validating config.yaml..."
    if command -v yamllint >/dev/null 2>&1; then
        yamllint config.yaml || echo "  ⚠️  YAML validation found issues"
    else
        echo "  ⚠️  yamllint not installed - install with: pip install yamllint"
    fi

    echo ""
    echo "  Checking required fields..."
    REQUIRED_FIELDS=("name" "version" "slug" "description" "arch")
    for field in "${REQUIRED_FIELDS[@]}"; do
        if grep -q "^$field:" config.yaml; then
            echo "    ✅ $field"
        else
            echo "    ❌ $field (MISSING - REQUIRED)"
        fi
    done
    echo ""
fi

# Validate build.json if it exists
if [ -f "build.json" ]; then
    echo "🔍 Validating build.json..."
    if command -v jq >/dev/null 2>&1; then
        if jq empty build.json 2>/dev/null; then
            echo "  ✅ Valid JSON"
            echo ""
            echo "  Checking build_from architectures..."
            ARCHS=("amd64" "armhf" "armv7" "aarch64" "i386")
            for arch in "${ARCHS[@]}"; do
                if jq -e ".build_from.${arch}" build.json >/dev/null 2>&1; then
                    IMAGE=$(jq -r ".build_from.${arch}" build.json)
                    echo "    ✅ $arch: $IMAGE"
                else
                    echo "    ⚠️  $arch: not defined"
                fi
            done
        else
            echo "  ❌ Invalid JSON format"
        fi
    else
        echo "  ⚠️  jq not installed - install with: sudo apt install jq"
    fi
    echo ""
fi

# Check Dockerfile
if [ -f "Dockerfile" ]; then
    echo "🐳 Checking Dockerfile..."
    if command -v hadolint >/dev/null 2>&1; then
        hadolint Dockerfile || echo "  ⚠️  Dockerfile has issues"
    else
        echo "  ℹ️  Install hadolint for Dockerfile linting"
    fi
    echo ""
fi

# Summary and recommendations
echo "=================================================="
echo "📝 Recommendations for HA Addon:"
echo ""

if [ ! -f "config.yaml" ]; then
    echo "1. Create config.yaml with addon configuration"
    echo "   See: https://developers.home-assistant.io/docs/add-ons/configuration"
    echo ""
fi

if [ ! -f "build.json" ]; then
    echo "2. Create build.json with multi-arch build configuration"
    echo "   Example: https://developers.home-assistant.io/docs/add-ons/configuration#add-on-advanced-builds"
    echo ""
fi

if [ ! -f "DOCS.md" ]; then
    echo "3. Create DOCS.md with user-facing documentation"
    echo "   Include: configuration options, setup instructions, troubleshooting"
    echo ""
fi

if [ ! -f "icon.png" ]; then
    echo "4. Add icon.png (256x256 pixels) for addon store visibility"
    echo ""
fi

if [ ! -f "run.sh" ]; then
    echo "5. Create run.sh as the addon entry point"
    echo "   This script should start your application"
    echo ""
fi

echo "For MQTT-based addons like this:"
echo "  - Support both internal HA MQTT broker and external brokers"
echo "  - Use MQTT auto-discovery for seamless integration"
echo "  - Handle connection failures gracefully with retries"
echo "  - Provide clear config validation and error messages"
echo ""

echo "✅ HA Addon structure check complete"
echo ""
