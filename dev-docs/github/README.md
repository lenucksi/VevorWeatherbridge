# GitHub Actions Workflows

## build-addon.yml

Builds and publishes Docker images for the Home Assistant add-on to GitHub Container Registry.

### Triggers

- Push to `main` branch
- Push of version tags (`v*`)
- Pull requests to `main`
- Manual workflow dispatch

### What it does

1. Builds Docker images for all supported architectures:
   - amd64 (x86_64)
   - armv7 (32-bit ARM)
   - aarch64 (64-bit ARM)
   - armhf (ARM hard float)
   - i386 (32-bit x86)

2. Publishes to GitHub Container Registry at:
   - `ghcr.io/lenucksi/vevor-weatherbridge-{arch}:0.1.0`
   - `ghcr.io/lenucksi/vevor-weatherbridge-{arch}:latest`

3. Uses GitHub Actions cache for faster builds

### Registry Permissions

The images are published to GitHub Container Registry. After the first build:

1. Go to <https://github.com/users/lenucksi/packages>
2. Find the `vevor-weatherbridge-*` packages
3. Click on each package
4. Go to "Package settings"
5. Under "Danger Zone", change visibility to **Public** (if desired)

This allows Home Assistant to pull the images without authentication.

### Manual Trigger

To manually trigger a build:

1. Go to Actions tab in GitHub
2. Select "Build and Publish Add-on"
3. Click "Run workflow"
4. Select branch and run

### Versioning

Images are tagged with:

- `0.1.0` - Current version
- `latest` - Latest build from main branch
- Git tags (when pushing `v*` tags)
- Branch names (for PR testing)
