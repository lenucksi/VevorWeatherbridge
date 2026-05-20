---
id: doc-5
title: GitHub Actions Workflows
type: other
created_date: '2026-05-20 14:36'
updated_date: '2026-05-20 14:36'
---
# GitHub Actions Workflows

## Workflows
- **ci.yml:** Lint (golangci-lint, hadolint) + Test (go test, coverage, SonarQube) on PR/push
- **build-addon.yml:** Build multi-arch Docker images for 5 architectures → GHCR
- **release-please.yml:** Automated releases from version tags
- **dependency-review.yml:** Security review for dependency PRs (govulncheck, gosec)
- **scorecard.yml:** Weekly OpenSSF Scorecard
- **actionlint.yml:** Lint workflow files on PR

**Source:** dev-docs/github/README.md, .github/workflows/
