# Development Documentation

This directory contains development-specific documentation, session notes, and transitional materials used during the conversion of VevorWeatherbridge to a Home Assistant Add-on.

## Directory Structure

```
dev-docs/
├── README.md                    # This file
├── project-rules/               # Project standards and guidelines
│   └── CLAUDE.md               # Claude Code project rules and automation
├── harness/                     # Quality assurance harness documentation
│   └── HARNESS_SUMMARY.md      # Claude Code harness implementation overview
└── conversion/                  # HA Add-on conversion documentation
    ├── CONVERSION_COMPLETE.md  # Final conversion summary and checklist
    ├── NEXT_SESSION_START.md   # Session continuation guide (historical)
    ├── SESSION_SUMMARY.md      # Development session notes
    ├── STARTNOTIZEN.md         # Initial project notes (German)
    └── ICON_PLACEHOLDER.md     # Icon requirements and resources
```

## Documentation Categories

### Project Rules ([project-rules/](project-rules/))

Contains coding standards, quality requirements, and automation rules:

- **[CLAUDE.md](project-rules/CLAUDE.md)** - Complete project rules including:
  - Code quality standards (2025)
  - Automation hooks and triggers
  - Development workflow
  - Skills available
  - Project-specific rules (MQTT, unit conversion, error handling)

### Quality Harness ([harness/](harness/))

Documentation for the Claude Code quality assurance system:

- **[HARNESS_SUMMARY.md](harness/HARNESS_SUMMARY.md)** - Overview of the quality harness implementation with hooks, skills, and agents

For complete harness documentation, see [../.claude/MANIFEST.md](../.claude/MANIFEST.md).

### Add-on Conversion ([conversion/](conversion/))

Historical documentation from the Docker-to-HA-Addon conversion:

- **[CONVERSION_COMPLETE.md](conversion/CONVERSION_COMPLETE.md)** - Final summary of conversion work, checklist, and next steps
- **[NEXT_SESSION_START.md](conversion/NEXT_SESSION_START.md)** - Session continuation guide (kept for historical reference)
- **[SESSION_SUMMARY.md](conversion/SESSION_SUMMARY.md)** - Development session notes and decisions
- **[STARTNOTIZEN.md](conversion/STARTNOTIZEN.md)** - Initial project setup notes
- **[ICON_PLACEHOLDER.md](conversion/ICON_PLACEHOLDER.md)** - Icon requirements and resource links

## User-Facing Documentation

User documentation is in the project root:

- **[../README.md](../README.md)** - Main project README with installation instructions
- **[../DOCS.md](../DOCS.md)** - Comprehensive user documentation for the HA add-on
- **[../CHANGELOG.md](../CHANGELOG.md)** - Version history and release notes

## Quick Links

### For Developers

- **Project Setup**: See [../README.md](../README.md#option-2-standalone-docker-container)
- **Quality Standards**: [project-rules/CLAUDE.md](project-rules/CLAUDE.md)
- **Run Quality Checks**: `../.claude/skills/python-ci-skill/run.sh`
- **Validate Add-on**: `../.claude/skills/ha-addon-skill/run.sh`
- **Security Scan**: `../.claude/skills/security-scan-skill/run.sh`

### For Add-on Users

- **Installation Guide**: [../README.md](../README.md#option-1-home-assistant-add-on-recommended)
- **Configuration**: [../DOCS.md](../DOCS.md#configuration)
- **DNS Setup**: [../DOCS.md](../DOCS.md#setup-instructions)
- **Troubleshooting**: [../DOCS.md](../DOCS.md#troubleshooting)

## Development Workflow

1. **Read project rules**: [project-rules/CLAUDE.md](project-rules/CLAUDE.md)
2. **Make changes** to code
3. **Run quality checks**:
   ```bash
   poetry run ruff check --fix .
   poetry run ruff format .
   poetry run mypy weatherstation.py
   ```
4. **Validate with skills**:
   ```bash
   ../.claude/skills/python-ci-skill/run.sh
   ../.claude/skills/ha-addon-skill/run.sh
   ```
5. **Review and commit** changes

## Contributing

When contributing to this project:

1. Follow the coding standards in [project-rules/CLAUDE.md](project-rules/CLAUDE.md)
2. Ensure all quality checks pass
3. Update documentation if adding features
4. Test with actual Home Assistant instance
5. Update [../CHANGELOG.md](../CHANGELOG.md) for user-facing changes

## Historical Context

This directory preserves the development process and decisions made during the conversion from a standalone Docker container to a Home Assistant Add-on. These documents are kept for:

- Understanding design decisions
- Troubleshooting similar issues
- Learning from the development process
- Maintaining project continuity

For current project status, see [conversion/CONVERSION_COMPLETE.md](conversion/CONVERSION_COMPLETE.md).

## Archive Policy

Documents in this directory are generally not deleted but may be:

- Moved to subdirectories for better organization
- Marked as historical/archived in their content
- Referenced but not actively maintained

Active documentation lives in:
- Project root (README.md, DOCS.md, CHANGELOG.md)
- `.claude/` directory (harness and skills)
