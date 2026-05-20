---
id: doc-4
title: AppArmor Profile Reference
type: other
created_date: '2026-05-20 14:36'
updated_date: '2026-05-20 14:36'
---
# AppArmor Profile Reference

## Key Insight
Custom AppArmor profiles completely REPLACE Docker's default profile. Everything needed must be explicitly allowed.

## Working Profile
Requires `network unix,` permission for Supervisor API access (bashio::services → HTTP → supervisor hostname → DNS → Unix sockets).

## Modes
- SECURITY_DISABLE: No AppArmor
- SECURITY_PROFILE: Custom profile (current)
- SECURITY_DEFAULT: HA default

## Debugging
`journalctl _TRANSPORT=audit -g 'apparmor=DENIED'`

**Source:** dev-docs/homeassistant/APPARMOR_RESEARCH.md
