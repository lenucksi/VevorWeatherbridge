---
id: TASK-23
title: Migrate von pre-commit zu prek (Rust-basiert)
status: To Do
assignee: []
created_date: '2026-05-22 13:15'
labels:
  - tooling
  - quality
  - prek
dependencies: []
priority: medium
ordinal: 25000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Git Hooks von legacy pre-commit (Python/YAML) zu prek (Rust/TOML) migrieren. prek ist schneller, nativ in Rust und Drop-in-kompatibel. Vorbild: rig-for-red/prek.toml
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 prek.toml mit built-in Hooks + gitleaks + golangci-lint + hadolint + rumdl + cog verify
- [ ] #2 .pre-commit-config.yaml entfernt
- [ ] #3 prek validate-config erfolgreich
- [ ] #4 prek install funktioniert (shims in .git/hooks)
- [ ] #5 CLAUDE.md aktualisiert
<!-- AC:END -->
