---
id: decision-3
title: 'License: CC0 root + GPL-3.0-or-later Go + EPL-2.0/EDL-1.0 paho'
date: '2026-05-20 14:33'
status: accepted
---
## Context

Project started with Python implementation under CC0 1.0 Universal. After Go migration, new Go code and third-party dependencies (paho.mqtt.golang) introduced different licenses.

## Decision

Keep root repository CC0 1.0 Universal. Go implementation under GPL-3.0-or-later. paho.mqtt.golang dependency remains under EPL-2.0/EDL-1.0 (weak copyleft, linking exempted). No license conflict.

## Consequences

SPDX headers added to Go source files. Go add-on is GPL-3.0-or-later while Python was CC0. Third-party attribution document generated. NOTICE file still needed.

