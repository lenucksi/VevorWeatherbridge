---
id: doc-3
title: Dependency Management
type: guide
created_date: '2026-05-20 14:36'
updated_date: '2026-05-20 14:36'
---
# Dependency Management

## Schedule
- Weekly (Sunday), Timezone: Europe/Berlin
- Auto-merge for patches and minor updates
- Security vulns get immediate PRs with security label

## Managed Packages
- Go modules (gomod)
- Docker base images (docker)
- GitHub Actions

## Manual Updates
- `go get -u ./... && go mod tidy`
- Review CHANGELOG for breaking changes
- Run `go test ./...` to verify

**Source:** dev-docs/github/DEPENDENCY_MANAGEMENT.md, dev-docs/github/RENOVATE_SETUP.md
