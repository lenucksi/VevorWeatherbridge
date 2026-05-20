---
id: doc-1
title: Testing Guide
type: guide
created_date: '2026-05-20 14:36'
updated_date: '2026-05-20 14:36'
---
# Testing Guide

## Test Structure
- Go tests use standard `go test` with table-driven tests
- Tests cover: conversion functions, compass, MQTT, HTTP endpoint, time parsing

## Running Tests
- `go test -v ./...` (all tests)
- `go test -fuzz=. -fuzztime=30s` (fuzz tests)
- `go test -coverprofile=coverage.out ./...` (coverage)

## CI Integration
Tests run automatically on PR via ci.yml (golangci-lint + go test).

**Source:** dev-docs/guides/TESTING.md
