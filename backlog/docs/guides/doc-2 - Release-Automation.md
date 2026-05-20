---
id: doc-2
title: Release Automation
type: guide
created_date: '2026-05-20 14:36'
updated_date: '2026-05-20 14:36'
---
# Release Automation

## Workflow
1. Bump version: `./.claude/skills/version-management-skill/run.sh patch|minor|major`
2. Update CHANGELOG.md with release notes
3. Commit and push to main
4. GitHub Actions handles: multi-arch Docker build + GHCR push + GitHub release

## Version Management
- Source of truth: vevor-weatherbridge-go/config.yaml (version field)
- Semantic versioning (SemVer 2.0.0)
- Version injected via ldflags at build time

## Changelog
- Keep a Changelog format (Added, Changed, Fixed, etc.)
- Custom changelog sections (not auto-contributors)

**Source:** dev-docs/guides/RELEASE-AUTOMATION.md
