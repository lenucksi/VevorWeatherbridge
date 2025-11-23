# Release Notes Customization Research

## Current Setup

The project uses [release-please](https://github.com/googleapis/release-please) for automated releases with configuration in `.github/release-please-config.json`.

Current configuration uses:
- `"release-type": "simple"` for both addons
- Custom `changelog-sections` mapping conventional commits to sections
- `"include-component-in-tag": true` and `"include-v-in-tag": true`

## Research Question

How to achieve nicer release descriptions like [v0.1.7](https://github.com/lenucksi/VevorWeatherbridge/releases/tag/v0.1.7) with contributor credits automatically?

## Findings

### v0.1.7 Release Format

The v0.1.7 release was **manually edited** and includes:
- Custom header with emoji ("Critical Fix: MQTT Discovery Now Works! 🎉")
- Organized sections (What's New, How It Works, Installation)
- Contributor credits with GitHub avatars and @mentions
- Full changelog link

### release-please Configuration Options

| Option | Description |
|--------|-------------|
| `changelog-type` | `"default"` (current) or `"github"` |
| `changelog-sections` | Custom commit type → section mapping |
| `pull-request-header` | Customize PR header text |
| `pull-request-footer` | Customize PR footer text |
| `draft` | Create releases as drafts |
| `prerelease` | Mark as prerelease |

### changelog-type Options

#### `"default"` (current)
- Uses custom `changelog-sections` configuration
- Groups commits by type with PR/commit links
- **No automatic contributor list**

#### `"github"`
- Uses GitHub's API to generate release notes
- **Automatically includes contributors with avatars**
- GitHub-native format matching [auto-generated release notes](https://docs.github.com/en/repositories/releasing-projects-on-github/automatically-generated-release-notes)
- **Ignores custom `changelog-sections`** - uses GitHub's own categorization

### Trade-offs

| Approach | Pros | Cons |
|----------|------|------|
| Keep `default` | Custom section grouping, detailed commit categorization | No auto contributors, manual editing for nice releases |
| Switch to `github` | Auto contributors with avatars, GitHub-native format | Loses custom `changelog-sections`, less control over grouping |
| Manual editing | Full control over release appearance | Manual work for each release |

### GitHub's release.yml

GitHub supports a separate `.github/release.yml` for configuring [auto-generated release notes](https://docs.github.com/en/repositories/releasing-projects-on-github/automatically-generated-release-notes). This can:
- Define categories for organizing PRs
- Exclude certain labels/authors
- Set custom labels for categories

However, this works with GitHub's native release UI, not directly with release-please.

## Decision

**Keep current configuration** for now:
- Custom `changelog-sections` provides useful categorization
- Manual enhancement of important releases (like v0.1.7) when needed
- Re-evaluate if release volume increases significantly

## Future Considerations

1. **Hybrid approach**: Use `draft: true` in release-please, then manually enhance before publishing
2. **Post-processing**: Add a GitHub Action step to enhance release notes after release-please creates them
3. **Switch to `github` type**: If contributor credits become more important than custom sections

## References

- [release-please repository](https://github.com/googleapis/release-please)
- [release-please customizing docs](https://github.com/googleapis/release-please/blob/main/docs/customizing.md)
- [release-please manifest docs](https://github.com/googleapis/release-please/blob/main/docs/manifest-releaser.md)
- [release-please config schema](https://raw.githubusercontent.com/googleapis/release-please/main/schemas/config.json)
- [GitHub auto-generated release notes](https://docs.github.com/en/repositories/releasing-projects-on-github/automatically-generated-release-notes)
- [Release Drafter](https://github.com/release-drafter/release-drafter) - alternative tool with more contributor options
