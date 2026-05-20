---
id: TASK-5
title: Revisit Dockerfile base image pinning
status: To Do
assignee: []
created_date: '2026-05-20 14:32'
labels:
  - docker
  - security
dependencies: []
references:
  - dev-docs/security/CODEQL_FINDINGS_ANALYSIS.md
priority: low
ordinal: 5000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Dockerfile uses unpinned base image tag. Accepted risk, but worth evaluating if pinning to a digest is now feasible.
<!-- SECTION:DESCRIPTION:END -->
