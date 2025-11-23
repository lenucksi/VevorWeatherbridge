# VS Code Workspace Extension Configuration Research Report

**Date**: 2025-11-23
**Status**: Research Complete
**Related Prompt**: [prompts/VSCODE_EXTENSIONS_SKILL_PROMPT.md](../prompts/VSCODE_EXTENSIONS_SKILL_PROMPT.md)

## Executive Summary

VS Code supports workspace-level extension configuration through `.vscode/extensions.json`, which can be version-controlled and automatically loaded. However, this mechanism only provides **recommendations** - it does NOT automatically enable or disable extensions. A feature request for auto-enable/disable has been open since 2017 with 786+ upvotes but remains unimplemented.

## Research Findings

### 1. Available Mechanisms

| Mechanism | Location | Auto-Loaded | Version Controllable | Effect |
|-----------|----------|-------------|---------------------|--------|
| `extensions.json` | `.vscode/extensions.json` | Yes | Yes | Prompts user to install/warns |
| Profiles | User data directory | No (manual import) | Via export | Full config bundle |
| Manual disable | VS Code settings | N/A | No | Per-workspace disable |

### 2. `.vscode/extensions.json` Format

This is the primary workspace extension configuration file:

```json
{
  "recommendations": [
    "ms-python.python",
    "charliermarsh.ruff",
    "esbenp.prettier-vscode"
  ],
  "unwantedRecommendations": [
    "ms-vscode.vscode-typescript-tslint-plugin",
    "deprecated.extension-id"
  ]
}
```

**Behavior:**

- When a user opens the workspace for the first time, VS Code prompts them to install recommended extensions
- The `@recommended` search filter in Extensions view shows these extensions
- `unwantedRecommendations` does NOT auto-disable - it only suppresses those extensions from appearing in recommendations

### 3. VS Code Profiles

Profiles provide a more comprehensive solution but require manual action:

**What profiles include:**

- Extensions (enabled/disabled state)
- Settings
- Keyboard shortcuts
- Snippets
- Tasks
- MCP servers
- UI layout

**Export/Import:**

1. Export: `File > Preferences > Profiles > Export...` → Creates `.code-profile` file
2. Import: `File > Preferences > Profiles > Import Profile...`
3. CLI: `code --profile "Profile Name" /path/to/folder`

**Limitation:** Profiles are stored in user data, not the workspace folder. To share with a team, you must export to a `.code-profile` file and have collaborators import it manually.

### 4. The Missing Feature

**GitHub Issue [#40239](https://github.com/microsoft/vscode/issues/40239)** requests the ability to enable/disable extensions from a workspace configuration file.

- Opened: December 2017
- Status: Open (assigned to "On Deck" milestone)
- Upvotes: 786+
- Comments: 414+

**VS Code team's concern:** Extension recommendations are designed for team sharing, while disabling extensions is considered user-specific. They've been hesitant to merge these concepts.

### 5. Workarounds

#### Option A: GARAIO Extension

The [vscode-unwanted-recommendations](https://github.com/garaio/vscode-unwanted-recommendations) extension adds enforcement:

- Monitors enabled extensions
- Shows warning when an extension in `unwantedRecommendations` is enabled
- Does NOT auto-disable, but alerts the user

#### Option B: Manual Workspace Disable

Users can manually disable extensions per workspace:

1. Open Extensions view
2. Click gear icon on extension
3. Select "Disable (Workspace)"

This setting is stored in VS Code's internal state, not in `.vscode/`.

#### Option C: VS Code Profiles + Documentation

1. Create a profile with desired extensions
2. Export to `.code-profile` file
3. Commit to repository
4. Document import instructions for team members

## Recommended Configuration for This Project

### `.vscode/extensions.json`

```json
{
  "recommendations": [
    "ms-python.python",
    "ms-python.vscode-pylance",
    "charliermarsh.ruff",
    "ms-python.debugpy",
    "ms-python.mypy-type-checker",
    "tamasfe.even-better-toml",
    "redhat.vscode-yaml",
    "timonwong.shellcheck"
  ],
  "unwantedRecommendations": [
    "ms-python.pylint"
  ]
}
```

### `.gitignore` Entries

```gitignore
# VS Code workspace settings (selective)
.vscode/*
!.vscode/settings.json
!.vscode/tasks.json
!.vscode/launch.json
!.vscode/extensions.json
!.vscode/*.code-snippets
```

## Implications for Claude Code Skills

A skill can generate `.vscode/extensions.json` files that will be:

- Automatically loaded by VS Code
- Version-controllable (committable to git)
- Cross-platform compatible

The skill **cannot**:

- Force extensions to be enabled or disabled
- Guarantee extensions are installed
- Override user preferences

The skill **should**:

- Clearly document the recommendation-only nature
- Suggest the GARAIO extension for enforcement
- Provide instructions for manual workspace disabling

## Sources

- [Extension Marketplace - VS Code](https://code.visualstudio.com/docs/configure/extensions/extension-marketplace) - Official documentation on workspace recommendations
- [Profiles in Visual Studio Code](https://code.visualstudio.com/docs/configure/profiles) - Profile export/import documentation
- [What is a VS Code workspace?](https://code.visualstudio.com/docs/editor/workspaces) - Workspace concepts
- [GitHub Issue #40239](https://github.com/microsoft/vscode/issues/40239) - Feature request for auto enable/disable
- [GARAIO vscode-unwanted-recommendations](https://github.com/garaio/vscode-unwanted-recommendations) - Enforcement extension
- [Stack Overflow - Workspace extension management](https://stackoverflow.com/questions/76919221/how-can-you-manage-set-not-recommend-what-extensions-are-enabled-disabled-in) - Community discussion

## Next Steps

1. Use the skill prompt at [prompts/VSCODE_EXTENSIONS_SKILL_PROMPT.md](../prompts/VSCODE_EXTENSIONS_SKILL_PROMPT.md) to create the skill
2. Generate `.vscode/extensions.json` for this project using the Python preset
3. Update `.gitignore` to include VS Code config files
4. Consider adding the GARAIO extension to recommendations for enforcement
