# VS Code Workspace Extension Configuration Skill - Creation Prompt

Use this prompt with a skill creator skill to generate a Claude Code skill that creates VS Code workspace extension configuration files.

---

## Context

VS Code supports workspace-level extension configuration through two mechanisms:

1. **`.vscode/extensions.json`** - Native, auto-loaded, version-controllable
2. **`.code-profile` files** - Comprehensive profiles that must be manually imported

**Important Limitation**: VS Code does NOT automatically enable/disable extensions based on configuration files. The `extensions.json` file only provides *recommendations* - users must manually act on them.

### Official Documentation

- [Extension Marketplace](https://code.visualstudio.com/docs/configure/extensions/extension-marketplace) - Workspace recommendations
- [Profiles in VS Code](https://code.visualstudio.com/docs/configure/profiles) - Profile export/import
- [Feature Request #40239](https://github.com/microsoft/vscode/issues/40239) - Auto enable/disable (not implemented)

## Goal

Create a Claude Code skill that generates VS Code workspace extension configuration files with the following capabilities:

1. Generate `.vscode/extensions.json` with recommendations and unwanted recommendations
2. Optionally generate a `.code-profile` export file for team sharing
3. Update `.gitignore` to include VS Code config files appropriately
4. Provide clear documentation about limitations

## Skill Specification

### Skill Name

`vscode-extensions-skill`

### Skill Location

`.claude/skills/vscode-extensions-skill/`

### Files to Create

```
.claude/skills/vscode-extensions-skill/
├── run.sh                    # Main skill entry point
├── README.md                 # Skill documentation
└── templates/
    ├── extensions.json.tmpl  # Template for extensions.json
    └── gitignore-snippet.txt # Recommended .gitignore entries
```

### Skill Arguments

The skill should accept the following arguments:

| Argument | Required | Description |
|----------|----------|-------------|
| `--recommend` | No | Comma-separated list of extension IDs to recommend |
| `--unwanted` | No | Comma-separated list of extension IDs to mark as unwanted |
| `--preset` | No | Use a preset configuration (python, go, typescript, rust, etc.) |
| `--profile` | No | Also generate a `.code-profile` file |
| `--gitignore` | No | Update `.gitignore` with recommended VS Code entries |
| `--output-dir` | No | Output directory (default: current working directory) |

### Preset Configurations

The skill should include presets for common development environments:

#### Python Preset
```json
{
  "recommendations": [
    "ms-python.python",
    "ms-python.vscode-pylance",
    "charliermarsh.ruff",
    "ms-python.debugpy"
  ],
  "unwantedRecommendations": [
    "ms-python.pylint"
  ]
}
```

#### Go Preset
```json
{
  "recommendations": [
    "golang.go"
  ],
  "unwantedRecommendations": []
}
```

#### TypeScript Preset
```json
{
  "recommendations": [
    "dbaeumer.vscode-eslint",
    "esbenp.prettier-vscode"
  ],
  "unwantedRecommendations": [
    "ms-vscode.vscode-typescript-tslint-plugin"
  ]
}
```

#### Rust Preset
```json
{
  "recommendations": [
    "rust-lang.rust-analyzer"
  ],
  "unwantedRecommendations": [
    "rust-lang.rust"
  ]
}
```

### Output: `.vscode/extensions.json`

```json
{
  "recommendations": [
    "ms-python.python",
    "charliermarsh.ruff"
  ],
  "unwantedRecommendations": [
    "ms-vscode.deprecated-extension"
  ]
}
```

### Output: `.gitignore` Snippet

When `--gitignore` is specified, append or create:

```gitignore
# VS Code workspace settings (selective)
.vscode/*
!.vscode/settings.json
!.vscode/tasks.json
!.vscode/launch.json
!.vscode/extensions.json
!.vscode/*.code-snippets
```

### Output: `.code-profile` (Optional)

When `--profile` is specified, generate a profile file that can be imported via:
`File > Preferences > Profiles > Import Profile...`

The profile should be a JSON file with this structure:

```json
{
  "name": "Project Profile",
  "extensions": [
    {
      "identifier": {
        "id": "ms-python.python"
      },
      "displayName": "Python"
    }
  ],
  "settings": {}
}
```

**Note**: The exact `.code-profile` format is not fully documented. The skill should generate a minimal working profile or note this limitation.

## Implementation Requirements

### 1. Shell Script (`run.sh`)

```bash
#!/usr/bin/env bash
set -euo pipefail

# Parse arguments
RECOMMEND=""
UNWANTED=""
PRESET=""
PROFILE=false
GITIGNORE=false
OUTPUT_DIR="."

while [[ $# -gt 0 ]]; do
  case $1 in
    --recommend) RECOMMEND="$2"; shift 2 ;;
    --unwanted) UNWANTED="$2"; shift 2 ;;
    --preset) PRESET="$2"; shift 2 ;;
    --profile) PROFILE=true; shift ;;
    --gitignore) GITIGNORE=true; shift ;;
    --output-dir) OUTPUT_DIR="$2"; shift 2 ;;
    *) echo "Unknown option: $1"; exit 1 ;;
  esac
done

# Implementation continues...
```

### 2. Extension ID Validation

Extension IDs follow the pattern: `publisher.extension-name`

The skill should validate that provided extension IDs match this pattern.

### 3. User Feedback

The skill should output:
- List of files created/modified
- Clear explanation of what each file does
- Warning about the recommendation-only limitation
- Instructions for users to act on recommendations

Example output:

```
Created: .vscode/extensions.json

Recommended extensions (VS Code will prompt to install):
  - ms-python.python
  - charliermarsh.ruff

Unwanted extensions (VS Code will NOT auto-disable):
  - ms-python.pylint

NOTE: VS Code only provides recommendations. Extensions are NOT
automatically enabled or disabled. Users must manually:
1. Install recommended extensions when prompted
2. Disable unwanted extensions via Extensions > Gear > Disable (Workspace)

Consider installing the 'garaio.vscode-unwanted-recommendations' extension
to get warnings when unwanted extensions are enabled.
```

## Behavior Specifications

### Merge vs Overwrite

When `.vscode/extensions.json` already exists:
- Read existing file
- Merge new recommendations (avoid duplicates)
- Warn user about the merge
- Allow `--force` flag to overwrite instead

### Idempotency

Running the skill multiple times with the same arguments should produce the same result.

### Error Handling

- Invalid extension ID format: Error with example of correct format
- Missing `.vscode` directory: Create it automatically
- Existing file conflicts: Prompt or use `--force`

## Testing the Skill

After creation, test with:

```bash
# Basic usage
./.claude/skills/vscode-extensions-skill/run.sh --preset python

# Custom extensions
./.claude/skills/vscode-extensions-skill/run.sh \
  --recommend "ms-python.python,charliermarsh.ruff" \
  --unwanted "ms-python.pylint"

# Full setup
./.claude/skills/vscode-extensions-skill/run.sh \
  --preset python \
  --gitignore \
  --profile
```

## Success Criteria

1. Skill generates valid `.vscode/extensions.json`
2. Generated JSON passes VS Code's schema validation
3. Presets cover common development environments
4. User feedback clearly explains limitations
5. Merge behavior preserves existing recommendations
6. `.gitignore` updates are non-destructive

## References

- [VS Code Extension Marketplace Docs](https://code.visualstudio.com/docs/configure/extensions/extension-marketplace)
- [VS Code Profiles Docs](https://code.visualstudio.com/docs/configure/profiles)
- [VS Code Workspaces](https://code.visualstudio.com/docs/editor/workspaces)
- [GARAIO Unwanted Recommendations Extension](https://github.com/garaio/vscode-unwanted-recommendations)
- [GitHub Issue #40239 - Auto Enable/Disable Feature Request](https://github.com/microsoft/vscode/issues/40239)

---

**After creating the skill, test it on this repository with:**

```bash
./.claude/skills/vscode-extensions-skill/run.sh \
  --preset python \
  --recommend "ms-python.mypy-type-checker,tamasfe.even-better-toml" \
  --gitignore
```
