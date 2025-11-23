# Renovate Bot Setup Guide

This guide explains how to enable Renovate Bot for automated dependency management on this repository.

## Option 1: GitHub App (Recommended)

### Steps

1. **Install the Renovate GitHub App**
   - Visit: <https://github.com/apps/renovate>
   - Click "Install" or "Configure"
   - Select the repository: `lenucksi/VevorWeatherbridge`
   - Grant permissions (read/write to code, PRs, issues)

2. **Verify Installation**
   - Renovate will create an onboarding PR within minutes
   - Review and merge the onboarding PR
   - Renovate will then create the Dependency Dashboard issue

3. **Configure (Optional)**
   - Configuration is already in `.github/renovate.json`
   - You can modify schedule, grouping, automerge settings

### Permissions Required

- ✅ Read access to code
- ✅ Write access to pull requests
- ✅ Write access to issues
- ✅ Read access to workflows

## Option 2: Self-Hosted Renovate

If you prefer to run Renovate yourself:

### Using GitHub Actions

```yaml
# .github/workflows/renovate.yml
name: Renovate
on:
  schedule:
    - cron: '0 0 * * 0'  # Weekly on Sunday
  workflow_dispatch:

jobs:
  renovate:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Self-hosted Renovate
        uses: renovatebot/github-action@v40
        with:
          configurationFile: .github/renovate.json
          token: ${{ secrets.RENOVATE_TOKEN }}
```

### Required Secrets

- `RENOVATE_TOKEN`: GitHub Personal Access Token with repo permissions

## Configuration Overview

The repository is pre-configured with `.github/renovate.json`:

### Key Settings

- **Schedule**: Every weekend (Sunday)
- **Timezone**: Europe/Berlin
- **Concurrent PRs**: Max 5 at a time
- **Auto-merge**: Disabled (manual review required)
- **Dependency Dashboard**: Enabled
- **Security Alerts**: Enabled with `security` label

### Package Management

- **Python (requirements.txt)**: Grouped minor/patch updates, separate major updates
- **Docker (Dockerfile)**: Tracks Home Assistant base images
- **GitHub Actions**: Automatic version updates

## After Installation

### What Happens Next

1. **Onboarding PR** (first time only)
   - Reviews current dependencies
   - Proposes configuration
   - Merge to activate Renovate

2. **Dependency Dashboard**
   - Created as a GitHub issue
   - Shows pending updates
   - Tracks rate limits and errors

3. **Update PRs**
   - Created according to schedule (weekly)
   - Grouped by update type
   - Include changelog links

### First Actions

1. ✅ Review and merge onboarding PR
2. ✅ Check Dependency Dashboard
3. ✅ Review any immediate security updates
4. ✅ Configure labels if needed
5. ✅ Set up PR auto-assignment (optional)

## Testing Renovate

To test immediately after installation:

1. **Trigger Manual Check**:
   - Go to Dependency Dashboard issue
   - Check the box next to any dependency
   - Renovate will create PR within minutes

2. **Force Full Scan**:
   - Close and reopen the Dependency Dashboard
   - Or wait for next scheduled run

## Troubleshooting

### No PRs Created

- ✅ Check repository permissions
- ✅ Verify `.github/renovate.json` is valid JSON
- ✅ Look for errors in Dependency Dashboard
- ✅ Check GitHub Actions logs (if self-hosted)

### Too Many PRs

Adjust in `.github/renovate.json`:

```json
{
  "prConcurrentLimit": 2,
  "prHourlyLimit": 1
}
```

### Renovate Stops Working

- ✅ Check if GitHub App is still installed
- ✅ Verify repository permissions
- ✅ Look for rate limiting in Dependency Dashboard

## Monitoring

### Weekly Checklist

- [ ] Check Dependency Dashboard for updates
- [ ] Review security-labeled PRs (high priority)
- [ ] Merge dependency update PRs
- [ ] Test major version updates locally
- [ ] Update addon version if dependencies changed

### Dashboard Sections

- **Awaiting Schedule**: Updates waiting for next run
- **Rate Limited**: PRs delayed due to limits
- **Errored**: Failed updates (needs investigation)
- **Ignored**: Dependencies excluded from updates

## Security Best Practices

1. **Review Security PRs Immediately**
   - PRs with `security` label indicate vulnerabilities
   - Check CVE details in PR description
   - Test and merge ASAP

2. **Keep Renovate Updated**
   - GitHub App: Auto-updated by GitHub
   - Self-hosted: Update renovatebot/github-action regularly

3. **Monitor Permissions**
   - Renovate only needs what's configured
   - Review GitHub App permissions periodically

## Customization Examples

### Enable Auto-merge for Patch Updates

```json
{
  "packageRules": [
    {
      "matchUpdateTypes": ["patch"],
      "automerge": true
    }
  ]
}
```

### Ignore Specific Packages

```json
{
  "ignoreDeps": ["Flask", "paho-mqtt"]
}
```

### Change Schedule

```json
{
  "schedule": ["every 2 weeks on Monday"]
}
```

## Support

- **Renovate Docs**: <https://docs.renovatebot.com/>
- **GitHub App**: <https://github.com/apps/renovate>
- **Community**: <https://github.com/renovatebot/renovate/discussions>
- **Project Docs**: [DEPENDENCY_MANAGEMENT.md](DEPENDENCY_MANAGEMENT.md)

## Next Steps

After setting up Renovate:

1. Read [DEPENDENCY_MANAGEMENT.md](DEPENDENCY_MANAGEMENT.md) for workflow details
2. Review [CLAUDE.md](../dev-docs/project-rules/CLAUDE.md) for project guidelines
3. Set up notifications for `dependencies` and `security` labels
4. Consider enabling Dependabot security alerts as backup
