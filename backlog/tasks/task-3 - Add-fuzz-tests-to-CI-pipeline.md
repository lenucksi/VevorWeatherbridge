---
id: TASK-3
title: Add fuzz tests to CI pipeline
status: To Do
assignee: []
created_date: '2026-05-20 14:32'
labels:
  - ci
  - testing
dependencies: []
references:
  - dev-docs/FUZZING_RESEARCH.md
priority: medium
ordinal: 3000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Go fuzz tests exist in the codebase but are not wired into CI. Add go test -fuzz=. -fuzztime=30s step to ci.yml (run on schedule, not per-PR).
<!-- SECTION:DESCRIPTION:END -->
