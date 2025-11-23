# Release Automation Guide

## Overview

This project uses automated GitHub Actions workflows to create comprehensive release notes whenever a new version tag is pushed. This ensures consistent, detailed releases with minimal manual effort.

## How It Works

### Automated Release Workflow

The release process is fully automated via [.github/workflows/release.yml](../../.github/workflows/release.yml):

1. **Trigger**: Workflow activates when a version tag matching `v*.*.*` is pushed (e.g., `v0.1.7`)
2. **Changelog Extraction**: Automatically extracts the relevant section from [CHANGELOG.md](../../vevor-weatherbridge/CHANGELOG.md)
3. **Git History Analysis**:
   - Lists all contributors since the previous release
   - Generates commit list with hashes
   - Creates comparison link to previous version
4. **Release Note Generation**: Combines all information into comprehensive release notes
5. **GitHub Release**: Automatically publishes the release with complete documentation

### What Gets Included

Each automated release includes:

- **Version Header**: Clear version identification
- **Changelog**: Extracted from CHANGELOG.md for that version
- **Installation Instructions**:
  - Home Assistant addon update steps
  - Docker image pull commands for all architectures
- **Full Changelog Link**: Reference to complete version history
- **Contributors**: List of all contributors for this release
- **Commits**: Detailed list of all commits with hashes
- **Comparison Link**: GitHub diff between this and previous version

## Creating a New Release

### Step-by-Step Process

1. **Update Version Number**

   Edit [vevor-weatherbridge/config.yaml](../../vevor-weatherbridge/config.yaml):

   ```yaml
   version: 0.1.8  # Increment version
   ```

2. **Update Software Version in Code**

   Edit [vevor-weatherbridge/weatherstation.py](../../vevor-weatherbridge/weatherstation.py):

   ```python
   "origin": {
       "name": "VEVOR Weatherbridge",
       "sw_version": "0.1.8",  # Update version here
       "support_url": "https://github.com/C9H13NO3-dev/VevorWeatherbridge",
   }
   ```

3. **Add Changelog Entry**

   Edit [vevor-weatherbridge/CHANGELOG.md](../../vevor-weatherbridge/CHANGELOG.md):

   ```markdown
   ## [0.1.8] - 2025-11-XX

   ### Added
   - New feature description

   ### Fixed
   - Bug fix description

   ### Changed
   - Change description
   ```

4. **Commit Changes**

   ```bash
   git add vevor-weatherbridge/config.yaml vevor-weatherbridge/weatherstation.py vevor-weatherbridge/CHANGELOG.md
   git commit -m "Bump version to 0.1.8"
   ```

5. **Create and Push Tag**

   ```bash
   git tag v0.1.8
   git push origin main
   git push origin v0.1.8
   ```

6. **Wait for Automation**

   The GitHub Actions workflow will automatically:

   - Extract the `[0.1.8]` section from CHANGELOG.md
   - Generate contributor and commit lists
   - Create the release with comprehensive notes
   - Publish to GitHub Releases page

## Manual Release (If Needed)

If you need to create or edit a release manually:

```bash
# Create release with notes file
gh release create v0.1.8 --title "v0.1.8" --notes-file release_notes.md

# Edit existing release
gh release edit v0.1.8 --notes-file updated_notes.md
```

## Release Note Format

The automated workflow generates release notes in this format:

```markdown
# VEVOR Weatherbridge v0.1.X

[Changelog content extracted from CHANGELOG.md]

## Installation

### Home Assistant Addon

1. Update the addon to v0.1.X in Home Assistant
2. Restart the addon
3. Check the logs for any errors

### Docker Image

```bash
docker pull ghcr.io/lenucksi/vevor-weatherbridge-amd64:0.1.X
```

Available architectures: `amd64`, `armv7`, `aarch64`, `armhf`, `i386`

## Full Changelog

See [CHANGELOG.md](link) for complete version history.

## Contributors

- @contributor1
- @contributor2

## Commits

- Commit message 1 (hash1)
- Commit message 2 (hash2)

**Full Diff**: <https://github.com/lenucksi/VevorWeatherbridge/compare/v0.1.X-1...v0.1.X>

```text

```text

## Workflow Details

### File Structure

```

.github/
├── workflows/
│   └── release.yml           # Automated release workflow
├── RELEASE_NOTES_v0.1.7.md   # Historical release notes
└── RELEASE_NOTES_v0.1.6.md   # Historical release notes

```text

### Workflow Steps Breakdown

1. **Checkout Code** (`actions/checkout@v4`)
   - Fetches full git history (`fetch-depth: 0`)
   - Required for comparing with previous tags

2. **Extract Version**
   - Parses version from tag (removes `v` prefix)
   - Stores both version number and full tag

3. **Extract Changelog**
   - Uses `awk` to extract the specific version section
   - Looks for `## [VERSION]` header in CHANGELOG.md
   - Stops at the next version header

4. **Generate Git Information**
   - Finds previous tag: `git describe --abbrev=0 --tags`
   - Determines commit range: `PREV_TAG..CURRENT_TAG`
   - Extracts unique contributors: `git log --format='%an'`
   - Lists all commits with hashes

5. **Create Release**
   - Uses `softprops/action-gh-release@v2`
   - Publishes as stable release (not draft or prerelease)
   - Uploads generated release notes

## Best Practices

### Changelog Guidelines

Follow [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) format:

- Use **Added** for new features
- Use **Changed** for changes in existing functionality
- Use **Deprecated** for soon-to-be removed features
- Use **Removed** for now removed features
- Use **Fixed** for any bug fixes
- Use **Security** for vulnerability fixes

Mark critical changes:
```markdown

### Fixed

- **CRITICAL**: Description of critical fix
```

### Version Numbering

Follow [Semantic Versioning](https://semver.org/):

- **MAJOR** (1.0.0): Incompatible API changes
- **MINOR** (0.1.0): New functionality, backwards compatible
- **PATCH** (0.0.1): Backwards compatible bug fixes

For this addon project (pre-1.0.0):

- **MINOR** (0.1.x → 0.2.x): New features, improvements
- **PATCH** (0.1.1 → 0.1.2): Bug fixes, small improvements

### Commit Messages

Good commit messages help generate useful release notes:

```

```text

```text

```text

```bash

## Good

git commit -m "Add MQTT discovery support for automatic sensor detection"
git commit -m "Fix SUPERVISOR_TOKEN handling in bashio integration"

## Less helpful

git commit -m "Update code"
git commit -m "Fix bug"
```

## Troubleshooting

### Workflow Not Triggering

**Problem**: Pushed tag but no release created

**Solutions**:

1. Verify tag format matches `v*.*.*` pattern
2. Check GitHub Actions page for errors
3. Ensure `GITHUB_TOKEN` has write permissions

### Missing Changelog Content

**Problem**: Release notes missing changelog section

**Solutions**:

1. Verify CHANGELOG.md has entry for this version
2. Check version format matches: `## [X.Y.Z] - YYYY-MM-DD`
3. Ensure CHANGELOG.md is in `vevor-weatherbridge/` directory

### Contributors Not Listed

**Problem**: Contributors missing from release

**Solutions**:

1. Ensure commits have proper author information
2. Check git config: `git config user.name` and `git config user.email`
3. Verify commits are in the range between tags

## Examples

### Example Release v0.1.7

**Tag**: `v0.1.7`
**Release URL**: <https://github.com/lenucksi/VevorWeatherbridge/releases/tag/v0.1.7>

This release fixed MQTT Discovery by:

- Adding required `origin` field
- Adding availability tracking
- Configuring Last Will and Testament

The automated workflow:

1. Extracted changelog from CHANGELOG.md
2. Listed commit: `cc325be`
3. Credited contributors
4. Generated comparison link

### Example Release v0.1.6

**Tag**: `v0.1.6`
**Release URL**: <https://github.com/lenucksi/VevorWeatherbridge/releases/tag/v0.1.6>

This release migrated to bashio library following HA best practices.

## Future Improvements

Potential enhancements to the release process:

- [ ] Docker image build and push automation
- [ ] Automated testing before release creation
- [ ] Release candidate (RC) support for pre-releases
- [ ] Automated Home Assistant addon repository update
- [ ] Release notes preview in PRs
- [ ] Semantic version validation
- [ ] Breaking change detection

## References

- [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)
- [Semantic Versioning](https://semver.org/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [softprops/action-gh-release](https://github.com/softprops/action-gh-release)
