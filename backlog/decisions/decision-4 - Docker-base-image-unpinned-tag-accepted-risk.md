---
id: decision-4
title: Docker base image unpinned tag accepted risk
date: '2026-05-20 14:33'
status: accepted
---
## Context

CodeQL alert #37 flagged the Dockerfile using an unpinned base image tag (v2.x channel). Pinning a digest would make builds reproducible but break the HA add-on update mechanism.

## Decision

Accept CodeQL alert #37 as standard HA add-on pattern. Do not pin base image to a digest. Renovate is configured to update base images automatically on a schedule.

## Consequences

OSSF Scorecard stays at ~6/10 instead of higher. Dockerfile follows standard HA add-on conventions. Renovate provides automated but delayed updates.

