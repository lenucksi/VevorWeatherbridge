---
id: decision-5
title: 'Release notes: keep custom changelog over auto-contributors'
date: '2026-05-20 14:33'
status: accepted
---
## Context

release-please offers two changelog-type options: 'default' (custom sections) and 'github' (auto-contributors). The repo needs to decide which one to use.

## Decision

Keep current release-please configuration with custom changelog sections (default type). Manually enhance important releases with contributor credits. Forego auto-contributors from github changelog type.

## Consequences

Better organized changelogs with semantic sections (Added, Changed, Fixed). Manual effort needed for contributor shout-outs on major releases.

