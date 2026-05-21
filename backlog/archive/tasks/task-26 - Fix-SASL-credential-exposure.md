---
id: TASK-26
title: Fix SASL credential exposure
status: To Do
assignee: []
created_date: '2026-05-20 14:54'
labels:
  - security
  - authentication
dependencies: []
references:
  - .claude/agents/security-auditor.md
priority: high
ordinal: 26000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Credentials may leak via exception messages and debug logs. Implement proper credential wiping and redaction.
<!-- SECTION:DESCRIPTION:END -->
