# Project Consolidation Plan

This document outlines the structural issues found in the VevorWeatherbridge repository and the plan to resolve them.

## Executive Summary

The project has evolved with duplicate files at the root level and in the `vevor-weatherbridge/` subdirectory. This creates maintenance overhead, version mismatches, and CI/CD confusion. This plan consolidates everything properly.

## Current Structure Issues

### Duplicate Files

| File | Root Location | Subdirectory | Status |
|------|---------------|--------------|--------|
| `Dockerfile` | `/Dockerfile` | `/vevor-weatherbridge/Dockerfile` | Identical - root unused |
| `requirements.txt` | `/requirements.txt` | `/vevor-weatherbridge/requirements.txt` | Different content |
| `weatherstation.py` | `/weatherstation.py` | `/vevor-weatherbridge/weatherstation.py` | Root is outdated (v0.1.6) |

### Test Structure Duplication

| Location | Files | Status |
|----------|-------|--------|
| `/tests/` | `conftest.py`, `test_conversions.py`, `test_endpoint.py` | Subset of functionality |
| `/vevor-weatherbridge/` | `test_weatherstation.py`, `test_run_sh.py`, `pytest.ini` | More comprehensive |

### Version Mismatch

| File | Version | Issue |
|------|---------|-------|
| `pyproject.toml` | 0.1.0 | **Outdated** |
| `vevor-weatherbridge/config.yaml` | 0.1.7 | Current |

### Package Manager Confusion

| File | Purpose | Status |
|------|---------|--------|
| `pyproject.toml` | Poetry format | To be converted to UV/PEP 621 |
| `poetry.lock` | Poetry lock file | To be deleted |
| `uv.lock` | UV lock file | Active, used by CI |
| `/requirements.txt` | Pip format | To be deleted |
| `/vevor-weatherbridge/requirements.txt` | Docker build | To be generated from pyproject.toml |

## Consolidation Actions

### Phase 1: Remove Root Duplicates

**Files to delete:**

```bash
# These are duplicates or outdated
rm /Dockerfile                    # Duplicate of vevor-weatherbridge/Dockerfile
rm /requirements.txt              # Superseded by pyproject.toml
rm /weatherstation.py             # Outdated v0.1.6 copy
```

**Rationale:**

- All application code lives in `vevor-weatherbridge/`
- Docker builds from `vevor-weatherbridge/` context
- Root files cause confusion about what's authoritative

### Phase 2: Consolidate Tests

**Target structure:**

```text
vevor-weatherbridge/
├── tests/
│   ├── __init__.py
│   ├── conftest.py           # Merged from root tests/
│   ├── test_conversions.py   # Merged: root + subdirectory tests
│   ├── test_endpoint.py      # Merged: root + subdirectory tests
│   ├── test_mqtt_callbacks.py  # From subdirectory (new)
│   ├── test_wu_forwarding.py   # From subdirectory (new)
│   └── test_run_sh.py        # Keep from subdirectory
├── weatherstation.py
├── ...
```

**Actions:**

1. Create `vevor-weatherbridge/tests/` directory
2. Move `vevor-weatherbridge/pytest.ini` to `vevor-weatherbridge/tests/` or project root
3. Merge test files:
   - Combine `test_weatherstation.py` comprehensive tests with root `tests/` structure
   - Keep the more comprehensive subdirectory tests as base
   - Add unique edge cases from root `tests/test_conversions.py` (string inputs)
4. Delete root `/tests/` directory

**Rationale:**

- Code is in `vevor-weatherbridge/`, tests should be adjacent
- Easier to maintain single test location
- CI can run `pytest vevor-weatherbridge/tests/`

### Phase 3: UV Migration (Full)

**Convert `pyproject.toml` from Poetry to UV/PEP 621:**

```toml
# Before (Poetry format)
[tool.poetry]
name = "vevor-weatherbridge"
version = "0.1.0"
packages = [{include = "weatherstation.py"}]

[tool.poetry.dependencies]
python = "^3.12"
flask = "^3.0.0"
...

[tool.poetry.group.dev.dependencies]
ruff = "^0.8.4"
...

[build-system]
requires = ["poetry-core"]
build-backend = "poetry.core.masonry.api"

# After (UV/PEP 621 format)
[project]
name = "vevor-weatherbridge"
version = "0.1.7"
description = "Weather Station to Home Assistant Relay"
readme = "README.md"
requires-python = ">=3.12"
dependencies = [
    "flask>=3.0.0",
    "pytz>=2024.1",
    "paho-mqtt>=2.1.0",
    "requests>=2.31.0",
    "dnspython>=2.6.0",
]

[project.optional-dependencies]
dev = [
    "ruff>=0.8.4",
    "mypy>=1.13.0",
    "bandit>=1.7.10",
    "pip-audit>=2.7.3",
    "yamllint>=1.35.1",
    "types-requests>=2.31.0",
    "types-pytz>=2024.1.0",
]
test = [
    "pytest>=8.3.0",
    "pytest-cov>=6.0.0",
    "pytest-mock>=3.12.0",
]

[build-system]
requires = ["hatchling"]
build-backend = "hatchling.build"

[tool.hatch.build.targets.wheel]
packages = ["vevor-weatherbridge"]

# Keep existing tool configurations
[tool.ruff]
...

[tool.mypy]
...

[tool.pytest.ini_options]
testpaths = ["vevor-weatherbridge/tests"]
...
```

**Actions:**

1. Rewrite `pyproject.toml` in PEP 621 format
2. Update version to 0.1.7 (match config.yaml)
3. Delete `poetry.lock`
4. Regenerate `uv.lock` with `uv lock`
5. Update `vevor-weatherbridge/requirements.txt` for Docker:

   ```bash
   uv pip compile pyproject.toml -o vevor-weatherbridge/requirements.txt
   ```

### Phase 4: Update CI/CD Paths

**Changes to `.github/workflows/ci.yml`:**

```yaml
# Before
- run: uv run pytest --cov=weatherstation --cov-report=xml

# After
- run: uv run pytest vevor-weatherbridge/tests/ --cov=vevor-weatherbridge/weatherstation --cov-report=xml
```

**Changes to hadolint validation:**

```yaml
# Before
- run: hadolint Dockerfile

# After
- run: hadolint vevor-weatherbridge/Dockerfile
```

### Phase 5: Dockerfile Multi-Stage Build

**Current (single-stage):**

```dockerfile
ARG BUILD_FROM
FROM ${BUILD_FROM}

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY weatherstation.py .
COPY run.sh /
RUN chmod a+x /run.sh
EXPOSE 80
CMD ["/run.sh"]
```

**Upgraded (multi-stage):**

```dockerfile
ARG BUILD_FROM

# Stage 1: Install dependencies
FROM ${BUILD_FROM} AS builder

WORKDIR /build
COPY requirements.txt .
RUN pip install --no-cache-dir --target=/install -r requirements.txt

# Stage 2: Runtime image
FROM ${BUILD_FROM}

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# Copy installed packages from builder
COPY --from=builder /install /usr/local/lib/python3.12/site-packages/

# Copy application
WORKDIR /app
COPY weatherstation.py .
COPY run.sh /
RUN chmod a+x /run.sh

EXPOSE 80
CMD ["/run.sh"]
```

**Benefits:**

- Cleaner separation of build and runtime
- Potential for smaller final image (if build tools excluded)
- Follows HA add-on best practices

### Phase 6: Add Missing Add-on Files

**Create `vevor-weatherbridge/apparmor.txt`:**

```text
#include <tunables/global>

profile vevor-weatherbridge flags=(attach_disconnected,mediate_deleted) {
  #include <abstractions/base>

  # Network access (HTTP server + MQTT client)
  network inet stream,
  network inet dgram,
  network inet6 stream,
  network inet6 dgram,

  # Python interpreter
  /usr/bin/python3* ix,
  /usr/local/bin/python3* ix,

  # Application files
  /app/** r,
  /run.sh rx,

  # Temporary files
  /tmp/** rw,

  # Deny everything else by default
  deny /proc/** w,
  deny /sys/** w,
}
```

**Create `vevor-weatherbridge/logo.png`:**

- Copy from `icon.png` or create a wider version
- Recommended: 256x256 or similar aspect ratio

## File Changes Summary

### Files to Delete

- [ ] `/Dockerfile`
- [ ] `/requirements.txt`
- [ ] `/weatherstation.py`
- [ ] `/poetry.lock`
- [ ] `/tests/` (entire directory, after merging)
- [ ] `/vevor-weatherbridge/test_weatherstation.py` (after merging into tests/)
- [ ] `/vevor-weatherbridge/pytest.ini` (move to root or tests/)

### Files to Create

- [ ] `/vevor-weatherbridge/tests/__init__.py`
- [ ] `/vevor-weatherbridge/tests/conftest.py`
- [ ] `/vevor-weatherbridge/tests/test_conversions.py`
- [ ] `/vevor-weatherbridge/tests/test_endpoint.py`
- [ ] `/vevor-weatherbridge/tests/test_mqtt_callbacks.py`
- [ ] `/vevor-weatherbridge/tests/test_wu_forwarding.py`
- [ ] `/vevor-weatherbridge/apparmor.txt`
- [ ] `/vevor-weatherbridge/logo.png`

### Files to Modify

- [ ] `/pyproject.toml` - Convert to UV/PEP 621, update version
- [ ] `/vevor-weatherbridge/Dockerfile` - Multi-stage build
- [ ] `/vevor-weatherbridge/requirements.txt` - Regenerate from pyproject.toml
- [ ] `/.github/workflows/ci.yml` - Update test paths

## Verification Checklist

After consolidation:

- [ ] `uv sync` works from project root
- [ ] `uv run pytest vevor-weatherbridge/tests/` passes
- [ ] `uv run ruff check .` passes
- [ ] `uv run mypy vevor-weatherbridge/weatherstation.py` passes
- [ ] Docker build succeeds: `docker build -t test vevor-weatherbridge/`
- [ ] CI workflow passes on push
- [ ] Version in pyproject.toml matches config.yaml

## Migration Order

Execute in this order to minimize breakage:

1. **Convert pyproject.toml** to UV format, delete poetry.lock
2. **Create test structure** in vevor-weatherbridge/tests/
3. **Merge tests** from both locations
4. **Delete root duplicates** (Dockerfile, requirements.txt, weatherstation.py)
5. **Delete root tests/** directory
6. **Update CI paths**
7. **Upgrade Dockerfile** to multi-stage
8. **Add apparmor.txt and logo.png**
9. **Regenerate uv.lock and requirements.txt**
10. **Run full verification**

## Rollback Plan

If issues arise:

1. All changes tracked in git
2. Can revert individual commits
3. Keep poetry.lock in a branch until UV confirmed working
4. CI will catch regressions before merge
