# DevContainer Research Start Prompt

Use this prompt to start a new Claude Code session for researching and improving the devcontainer configuration.

---

## Context

This project has a `.devcontainer/devcontainer.json` file that needs to be reviewed and potentially improved. The project is a Python-based Home Assistant add-on with modern tooling (UV, ruff, mypy, pytest).

## Research Tasks

### 1. DevContainer Fundamentals

Research the official devcontainer specification at https://containers.dev/ and answer:

- What is the purpose of devcontainers?
- What is the devcontainer specification (vs VS Code implementation)?
- What are the core capabilities and limitations?
- What is the relationship between devcontainer.json and Dockerfile?

### 2. Editor/Environment Compatibility

Research and document compatibility with:

| Environment | Support Level | Notes |
|-------------|---------------|-------|
| VS Code | ? | Native support via extension |
| VS Code Web | ? | GitHub Codespaces |
| GitHub Codespaces | ? | Cloud-hosted devcontainers |
| JetBrains (IntelliJ, PyCharm) | ? | Gateway support? |
| Neovim | ? | Any devcontainer plugins? |
| Claude Code | ? | How does Claude Code use devcontainers? |
| GitPod | ? | devcontainer.json support? |
| DevPod | ? | Alternative to Codespaces |

### 3. Docker vs Podman

Research and document:

- Does the devcontainer spec support Podman?
- What configuration is needed for Podman compatibility?
- Are there limitations with Podman vs Docker?
- How to configure VS Code/other tools to use Podman?
- Rootless container considerations

### 4. Best Practices 2025

Research current best practices:

- Base image selection (Microsoft images vs custom)
- Feature vs manual installation
- devcontainer features ecosystem
- Multi-container setups (docker-compose)
- Performance optimization (caching, volumes)
- Security considerations
- CI/CD integration (GitHub Actions with devcontainers)
- Reproducibility guarantees

### 5. Project-Specific Recommendations

After research, analyze what should and should NOT be in this project's devcontainer:

**Should include:**
- Python 3.12 environment
- UV package manager
- Development tools (ruff, mypy, pytest)
- Git configuration
- VS Code extensions for Python

**Consider:**
- Docker-in-Docker (for building HA add-on images)
- MQTT broker for local testing
- Pre-commit hooks

**Should NOT include:**
- Production dependencies only
- Large unnecessary tools
- Secrets or credentials
- User-specific configurations

## Deliverables

Create a report in `dev-docs/devcontainer-analysis.md` containing:

1. **Executive Summary** - What devcontainers are, key benefits
2. **Specification Overview** - Core concepts, file structure
3. **Compatibility Matrix** - Which editors/environments support devcontainers
4. **Podman Support** - How to use with Podman instead of Docker
5. **Best Practices 2025** - Current recommendations
6. **Project Assessment** - Review of existing `.devcontainer/devcontainer.json`
7. **Recommendations** - Specific changes for this project
8. **Example Configuration** - Updated devcontainer.json if needed

## Reference Files

Read these files to understand the project:
- `.devcontainer/devcontainer.json` - Current configuration
- `pyproject.toml` - Project dependencies and tools
- `CLAUDE.md` - Project conventions

## Research Sources

Use web search and fetch to research from:
- https://containers.dev/ - Official specification
- https://code.visualstudio.com/docs/devcontainers/containers - VS Code docs
- https://github.com/devcontainers - Official GitHub org
- https://github.com/devcontainers/features - Available features
- JetBrains Gateway documentation
- Podman documentation for devcontainer support

---

**Start by reading the existing devcontainer.json, then research each topic systematically.**
