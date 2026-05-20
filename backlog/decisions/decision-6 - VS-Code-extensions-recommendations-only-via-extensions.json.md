---
id: decision-6
title: 'VS Code extensions: recommendations-only via extensions.json'
date: '2026-05-20 14:33'
status: accepted
---
## Context

VS Code has a feature request (#40239, open since 2017, 786+ upvotes) for auto-enable/disable of extensions from config. Until resolved, only recommendations are possible.

## Decision

Use .vscode/extensions.json with recommendations and unwantedRecommendations. Accept the limitation that extensions cannot be auto-enabled/disabled. Do not implement workarounds like the GARAIO extension.

## Consequences

Users see recommended extension popup but must manually install. No enforcement. Clean, standard configuration approach.

