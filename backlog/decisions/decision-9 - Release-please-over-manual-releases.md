---
id: decision-9
title: Release-please over manual releases
date: '2026-05-20 14:33'
status: accepted
---
## Context

Manual releases were error-prone and required multiple steps. An automated release process was needed for consistent multi-architecture Docker builds and GitHub releases.

## Decision

Use release-please for automated release management. Triggers on version tag push. Extracts changelog from CHANGELOG.md, creates GitHub release with contributor credits, triggers multi-arch Docker build.

## Consequences

Consistent releases. Automated Docker image publishing to GHCR for 5 architectures. Release notes include changelog excerpts and contributor lists.

