#!/usr/bin/env bash
# ==============================================================================
# Version Management Skill for VevorWeatherbridge
# Manages semantic versioning across config.yaml and CHANGELOG.md
# ==============================================================================

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

CONFIG_FILE="vevor-weatherbridge/config.yaml"
CHANGELOG_FILE="vevor-weatherbridge/CHANGELOG.md"

# Get current version from config.yaml
get_current_version() {
    grep '^version:' "$CONFIG_FILE" | awk '{print $2}' | tr -d '"'
}

# Validate semver format
validate_version() {
    local version=$1
    if [[ ! $version =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
        echo -e "${RED}ERROR: Invalid version format: $version${NC}"
        echo "Version must follow semantic versioning (major.minor.patch)"
        exit 1
    fi
}

# Parse version components
parse_version() {
    local version=$1
    IFS='.' read -r -a parts <<< "$version"
    echo "${parts[0]} ${parts[1]} ${parts[2]}"
}

# Bump version
bump_version() {
    local current=$1
    local bump_type=$2

    read -r major minor patch <<< "$(parse_version "$current")"

    case $bump_type in
        major)
            major=$((major + 1))
            minor=0
            patch=0
            ;;
        minor)
            minor=$((minor + 1))
            patch=0
            ;;
        patch)
            patch=$((patch + 1))
            ;;
        *)
            echo -e "${RED}ERROR: Invalid bump type: $bump_type${NC}"
            echo "Valid types: major, minor, patch"
            exit 1
            ;;
    esac

    echo "$major.$minor.$patch"
}

# Update config.yaml
update_config() {
    local new_version=$1
    local current_version=$2

    # Use sed to update version in config.yaml
    if [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS
        sed -i '' "s/^version: $current_version/version: $new_version/" "$CONFIG_FILE"
    else
        # Linux
        sed -i "s/^version: $current_version/version: $new_version/" "$CONFIG_FILE"
    fi

    echo -e "${GREEN}✓ Updated $CONFIG_FILE: $current_version -> $new_version${NC}"
}

# Add CHANGELOG entry
add_changelog_entry() {
    local new_version=$1
    local date=$(date +%Y-%m-%d)

    # Create new entry template
    local entry="## [$new_version] - $date

### Added
-

### Changed
-

### Fixed
-

"

    # Insert after the semver link line
    if [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS - use temp file approach
        awk -v entry="$entry" '
            /^and this project adheres to/ {
                print;
                print "";
                print entry;
                next
            }
            {print}
        ' "$CHANGELOG_FILE" > "$CHANGELOG_FILE.tmp" && mv "$CHANGELOG_FILE.tmp" "$CHANGELOG_FILE"
    else
        # Linux
        sed -i "/^and this project adheres to/a\\
\\
$entry" "$CHANGELOG_FILE"
    fi

    echo -e "${GREEN}✓ Added changelog entry for v$new_version${NC}"
    echo -e "${YELLOW}⚠ Please update the CHANGELOG.md with actual changes${NC}"
}

# Check version consistency
check_consistency() {
    local config_version=$(get_current_version)

    echo -e "${GREEN}Current version: $config_version${NC}"

    # Check if version exists in CHANGELOG
    if grep -q "## \[$config_version\]" "$CHANGELOG_FILE"; then
        echo -e "${GREEN}✓ Version $config_version found in CHANGELOG.md${NC}"
    else
        echo -e "${YELLOW}⚠ Version $config_version NOT found in CHANGELOG.md${NC}"
    fi
}

# Main
main() {
    local action=${1:-check}

    # Ensure we're in the project root
    if [[ ! -f "$CONFIG_FILE" ]]; then
        echo -e "${RED}ERROR: Must run from project root${NC}"
        exit 1
    fi

    local current_version=$(get_current_version)
    validate_version "$current_version"

    case $action in
        check)
            check_consistency
            ;;
        major|minor|patch)
            local new_version=$(bump_version "$current_version" "$action")
            echo -e "${YELLOW}Bumping version: $current_version -> $new_version ($action)${NC}"

            update_config "$new_version" "$current_version"
            add_changelog_entry "$new_version"

            echo ""
            echo -e "${GREEN}✓ Version bump complete!${NC}"
            echo ""
            echo "Next steps:"
            echo "1. Edit $CHANGELOG_FILE and fill in changes for v$new_version"
            echo "2. Commit: git add $CONFIG_FILE $CHANGELOG_FILE"
            echo "3. Commit: git commit -m 'Bump version to $new_version'"
            echo "4. Push: git push origin main"
            echo "5. GitHub Actions will build images and create release"
            ;;
        *)
            echo "Usage: $0 [check|major|minor|patch]"
            echo ""
            echo "Actions:"
            echo "  check  - Validate current version consistency"
            echo "  patch  - Bump patch version (0.1.2 -> 0.1.3)"
            echo "  minor  - Bump minor version (0.1.2 -> 0.2.0)"
            echo "  major  - Bump major version (0.2.0 -> 1.0.0)"
            exit 1
            ;;
    esac
}

main "$@"
