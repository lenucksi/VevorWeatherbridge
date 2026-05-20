---
id: decision-2
title: Add-on approach over native HA integration
date: '2026-05-20 14:33'
status: accepted
---
## Context

The project is a Home Assistant Add-on but could theoretically be migrated to a native HA integration. HA has a Bronze/Silver/Gold/Platinum Integration Quality Scale.

## Decision

Stay with the Home Assistant Add-on approach. It already works, has broader compatibility (MQTT works with any HA setup), lower maintenance burden. The HA Integration Quality Scale does NOT directly apply to add-ons.

## Consequences

No i18n/translations required. No Python HA framework dependency. Simpler deployment model. Gold/Platinum quality scale targets not applicable. Add-on is more portable across HA versions.

