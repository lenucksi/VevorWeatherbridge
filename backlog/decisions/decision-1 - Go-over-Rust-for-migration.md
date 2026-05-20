---
id: decision-1
title: Go over Rust for migration
date: '2026-05-20 14:33'
status: accepted
---
## Context

The project needed to migrate from Python (~150 MB image, ~50 MB RAM) to a more efficient language for Docker deployment. Go and Rust were both evaluated.

## Decision

Go was chosen over Rust for the migration. Key factors: trivial cross-compilation, single external dependency (paho.mqtt.golang), significantly faster development effort (~20-30h Go vs ~28-40h Rust), 10-15x image size reduction over Python.

## Consequences

~15 MB image (vs ~150 MB Python), milliseconds startup. All MQTT and HTTP logic ported to Go stdlib. Estimated ~400-500 lines across ~10 files.

