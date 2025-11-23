# Version Management Skill

## Purpose

Manages semantic versioning for the VevorWeatherbridge Home Assistant add-on, ensuring consistent version updates across config.yaml and CHANGELOG.md.

## When to Use

- When bumping the version for a new release
- When validating version consistency across files
- When preparing a new release with changelog updates

## What It Does

1. Validates current version format in config.yaml
2. Suggests next version based on semantic versioning (major.minor.patch)
3. Updates version in config.yaml
4. Creates new CHANGELOG.md entry with appropriate template
5. Validates consistency across all version references

## Usage

```bash
./.claude/skills/version-management-skill/run.sh [bump-type]
```

Where `bump-type` is one of:

- `patch` - Bug fixes, small changes (0.1.1 -> 0.1.2)
- `minor` - New features, backward compatible (0.1.2 -> 0.2.0)
- `major` - Breaking changes (0.2.0 -> 1.0.0)
- `check` - Validate current version only

## Workflow Integration

This skill is part of the release workflow:

1. Make code changes
2. Run skill to bump version
3. Update CHANGELOG.md with changes
4. Commit with appropriate message
5. Push to trigger GitHub Actions build and release

## Files Modified

- `vevor-weatherbridge/config.yaml` - Version field
- `vevor-weatherbridge/CHANGELOG.md` - New version entry

## Semantic Versioning Rules

- **Major (X.0.0)**: Breaking changes, incompatible API changes
- **Minor (0.X.0)**: New features, backward compatible
- **Patch (0.0.X)**: Bug fixes, backward compatible

## Example

```bash
# Check current version
./.claude/skills/version-management-skill/run.sh check

# Bump patch version (0.1.2 -> 0.1.3)
./.claude/skills/version-management-skill/run.sh patch

# Bump minor version (0.1.3 -> 0.2.0)
./.claude/skills/version-management-skill/run.sh minor
```
