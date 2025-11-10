# Dependency Management

This project uses **Renovate Bot** for automated dependency updates and vulnerability tracking.

## Overview

Renovate automatically:
- Checks for updates to Python packages (requirements.txt)
- Monitors Docker base image updates (Home Assistant base images)
- Tracks GitHub Actions version updates
- Creates pull requests for dependency updates
- Provides a centralized Dependency Dashboard
- Alerts on security vulnerabilities

## Configuration

### Renovate Configuration
Location: `.github/renovate.json`

Key features:
- **Schedule**: Runs every weekend (Sunday)
- **Timezone**: Europe/Berlin
- **Grouping**: Minor/patch Python updates are grouped, major updates get separate PRs
- **Labels**: All PRs tagged with `dependencies`, security issues tagged with `security`
- **Concurrent PRs**: Maximum 5 at a time
- **Auto-merge**: Disabled (requires manual review)

### Dependency Dashboard

Renovate creates and maintains an issue titled "Dependency Dashboard 🤖" that shows:
- Pending updates
- Rate-limited PRs
- Errored updates
- Ignored dependencies

## Workflows

### 1. Dependency Review (`dependency-review.yml`)
- **Triggers**: Pull requests to main branch with `dependencies` label or from renovate[bot]
- **Actions**:
  - Reviews dependency changes for security issues
  - Runs pip-audit for Python vulnerability scanning
  - Runs bandit for code security analysis
  - Comments summary in PR

### 2. Update Dependencies (`update-dependencies.yml`)
- **Triggers**:
  - Scheduled: Every Sunday at 02:00 UTC
  - Manual: Via workflow_dispatch
- **Actions**:
  - Lists outdated Python packages
  - Checks Docker base image status
  - Scans for security vulnerabilities
  - Generates summary report

## Package Types Managed

### Python Dependencies (requirements.txt)
- **Flask**: Web framework for weather endpoint
- **paho-mqtt**: MQTT client library
- **pytz**: Timezone handling
- **dnspython**: DNS resolution for Weather Underground
- **requests**: HTTP client

Update strategy:
- **Patch/Minor**: Grouped into single PR (e.g., Flask 3.0.0 -> 3.0.1)
- **Major**: Separate PRs for review (e.g., Flask 3.x -> 4.x)

### Docker Base Images
- **Home Assistant base images**: ghcr.io/home-assistant/*-base-python:3.12-alpine3.19
- Updates tracked for all 5 architectures (amd64, armv7, aarch64, armhf, i386)

### GitHub Actions
- **actions/checkout**
- **actions/setup-python**
- **docker/setup-qemu-action**
- **docker/setup-buildx-action**
- **docker/login-action**
- **docker/build-push-action**
- **docker/metadata-action**

## Reviewing Dependency Updates

### For Python Packages
1. Check the PR description for changelog links
2. Review `dependency-review` workflow results
3. Check for breaking changes in major updates
4. Verify pip-audit shows no new vulnerabilities
5. Test locally if major update

### For Docker Base Images
1. Review Home Assistant release notes
2. Check if Python version changed
3. Test multi-architecture builds
4. Verify addon still starts correctly

### For GitHub Actions
1. Review action changelog
2. Check for breaking changes in workflows
3. Verify all architectures build successfully

## Security Vulnerabilities

When Renovate detects a vulnerability:
1. Creates a PR with `security` label
2. Assigns to repository owner
3. Provides CVE details and severity
4. Suggests fixed version

**Action Required**:
- Review immediately for high/critical severity
- Test the update
- Merge and release new addon version

## Manual Dependency Updates

If you need to update dependencies manually:

```bash
# Check for outdated Python packages
pip list --outdated

# Update a specific package
pip install --upgrade <package-name>
pip freeze > vevor-weatherbridge/requirements.txt

# Run security scan
pip install pip-audit
pip-audit

# Update version and changelog
./.claude/skills/version-management-skill/run.sh patch
# Edit CHANGELOG.md with dependency updates

# Commit
git add vevor-weatherbridge/requirements.txt vevor-weatherbridge/config.yaml vevor-weatherbridge/CHANGELOG.md
git commit -m "chore(deps): update <package-name> to <version>"
git push origin main
```

## Ignoring Dependencies

To prevent Renovate from updating specific packages, add to `.github/renovate.json`:

```json
{
  "ignoreDeps": ["package-name"]
}
```

Or ignore specific versions:

```json
{
  "packageRules": [
    {
      "matchPackageNames": ["package-name"],
      "allowedVersions": "<=2.0"
    }
  ]
}
```

## Disabling Renovate

To temporarily disable Renovate (e.g., during major refactoring):

1. Add to `.github/renovate.json`:
   ```json
   {
     "enabled": false
   }
   ```

2. Or close the Dependency Dashboard issue

## Integration with CI/CD

Renovate PRs trigger:
- ✅ Build workflow (all architectures)
- ✅ Dependency review workflow
- ✅ Security scans (pip-audit, bandit)
- ❌ Release workflow (not triggered for dependency updates)

Dependency updates do **not** automatically create releases. After merging:
1. Test the addon
2. Use version management skill to bump version
3. Update CHANGELOG.md
4. Push to trigger release

## Best Practices

1. **Review Weekly**: Check Dependency Dashboard every weekend
2. **Group Updates**: Let Renovate group minor/patch updates
3. **Test Major Updates**: Always test major version updates locally
4. **Security First**: Prioritize security-labeled PRs
5. **Keep Updated**: Don't let dependencies drift too far behind
6. **Document Breaking Changes**: Note any required configuration changes in CHANGELOG.md

## Troubleshooting

### Renovate isn't creating PRs
- Check the Dependency Dashboard for errors
- Verify `.github/renovate.json` is valid JSON
- Check GitHub App installation permissions

### Too many PRs at once
- Reduce `prConcurrentLimit` in renovate.json
- Use `schedule` to control update timing
- Consider using `automerge` for patch updates

### Update breaks addon
1. Revert the PR
2. Add the package to `ignoreDeps` temporarily
3. Investigate the breaking change
4. Fix code or pin to previous version

## Resources

- [Renovate Documentation](https://docs.renovatebot.com/)
- [Renovate Configuration Options](https://docs.renovatebot.com/configuration-options/)
- [Home Assistant Add-on Dependencies](https://developers.home-assistant.io/docs/add-ons/configuration#add-on-dockerfile)
- [pip-audit Documentation](https://pypi.org/project/pip-audit/)
