---
id: decision-7
title: pyproject.toml kept at root
date: '2026-05-20 14:33'
status: superseded
---
## Context

The pyproject.toml file location was questioned during project consolidation. Moving it would affect tooling and build commands.

## Decision

Keep pyproject.toml at project root per PEP 517/518/621 standard location. Moving would break 'uv sync' and require --config flags for all tools.

## Consequences

Standard location, no tooling friction. Note: Mostly historical now as the project is Go-only and pyproject.toml is no longer actively used.

