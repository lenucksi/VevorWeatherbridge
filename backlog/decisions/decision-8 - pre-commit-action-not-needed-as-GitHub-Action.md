---
id: decision-8
title: pre-commit/action not needed as GitHub Action
date: '2026-05-20 14:33'
status: accepted
---
## Context

Evaluated whether pre-commit/action should be added as a GitHub Action step for CI enforcement.

## Decision

Not adding pre-commit/action to CI. It is in maintenance mode. CI already runs the same linters (hadolint, golangci-lint) directly in workflow steps. Local pre-commit hooks remain available for developers.

## Consequences

One less CI dependency and complexity point. CI and local pre-commit remain independent. No duplication of linting configuration.

