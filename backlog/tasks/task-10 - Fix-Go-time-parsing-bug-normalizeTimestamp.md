---
id: TASK-10
title: Fix Go time parsing bug (normalizeTimestamp)
status: Done
assignee: []
created_date: '2026-05-20 14:32'
labels:
  - parsing
  - bugfix
dependencies: []
references:
  - TODO.md
priority: high
ordinal: 10000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Weather station sends timestamps with single-digit time components. Added normalizeTimestamp() to pad single-digit hours/minutes/seconds.
<!-- SECTION:DESCRIPTION:END -->
