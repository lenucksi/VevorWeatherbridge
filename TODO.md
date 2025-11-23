# TODO

## Open

- Test new Go version live -> muss der mensch machen, dafür muss aber das gh release durchgehen

## Completed

- [x] Fix Go warnings with golangci-lint (4 errcheck warnings fixed)

- [x] Add Go tests to CI workflow (added `go test -v ./...` step to ci.yml)

- [x] Clean up RELEASE_NOTES files from .github/
  - Deleted `RELEASE_NOTES_v0.1.6.md` and `RELEASE_NOTES_v0.1.7.md`
  - These were historical artifacts not used by automated release workflow

- [x] pyproject.toml location decision
  - **Result**: Keep at root (standard location per PEP 517/518/621)
  - Moving would break `uv sync` and require `--config` flags for all tools

- [x] pre-commit/action as GitHub Action
  - **Result**: Not needed - `pre-commit/action` is in maintenance mode
  - CI already runs the same linters (ruff, hadolint, golangci-lint)
  - Local pre-commit hooks are sufficient

- [x] Research HomeAssistant bronze/silver/gold integration quality requirements
  - **Result**: Quality Scale does NOT apply to this project (it's for native HA integrations, not MQTT bridges)
  - See [dev-docs/homeassistant-integration-notes.md](dev-docs/homeassistant-integration-notes.md) for details

- [x] Consider Go rewrite for smaller Docker images (~5-10 MB vs ~150 MB)
  - See [dev-docs/go-rust-migration-analysis.md](dev-docs/go-rust-migration-analysis.md) for analysis

## Future Considerations
