# ha-addon-skill

## Role
Validates and assists with Home Assistant addon development for the VevorWeatherbridge project.

## Triggers
- Manual: `/skills run ha-addon-skill`
- Hook-based: Automatically suggested when addon config files are edited
- Validation mode: `/skills run ha-addon-skill --validate-config`

## Tools Required
- `yamllint` - YAML syntax validator
- `jq` - JSON processor for build.json validation
- Home Assistant addon schema knowledge (embedded)

## Responsibilities

### 1. Config Validation (`config.yaml`)
Required fields for HA addon:
- `name`: Addon name
- `version`: Semantic version (x.y.z)
- `slug`: Unique identifier (lowercase, hyphens)
- `description`: Short description
- `arch`: Supported architectures (amd64, armhf, armv7, aarch64, i386)
- `startup`: Startup type (application, services, system, once, before, after)
- `boot`: Boot option (auto, manual)
- `ports`: Port mappings (if any)
- `options`: Configuration options with schema
- `schema`: JSON schema for options validation

Optional but recommended:
- `image`: Custom docker image
- `map`: Volume mappings
- `environment`: Environment variables
- `privileged`: Network/device access requirements
- `homeassistant_api`: If requires HA API access
- `mqtt`: MQTT broker requirements

### 2. Build Configuration (`build.json`)
Required structure:
```json
{
  "build_from": {
    "amd64": "ghcr.io/home-assistant/amd64-base:latest",
    "armhf": "ghcr.io/home-assistant/armhf-base:latest",
    "armv7": "ghcr.io/home-assistant/armv7-base:latest",
    "aarch64": "ghcr.io/home-assistant/aarch64-base:latest",
    "i386": "ghcr.io/home-assistant/i386-base:latest"
  },
  "args": {}
}
```

### 3. Documentation (`DOCS.md`)
Required sections:
- Configuration options explanation
- Installation instructions
- Usage guide
- Troubleshooting
- Examples

### 4. Addon Structure
Expected directory layout:
```
addon/
├── config.yaml       # Addon configuration
├── build.json        # Build configuration
├── Dockerfile        # Build instructions
├── DOCS.md          # User documentation
├── CHANGELOG.md     # Version history
├── README.md        # Developer documentation
├── icon.png         # Addon icon (256x256)
├── logo.png         # Addon logo (optional)
├── run.sh           # Startup script
└── ...              # Application files
```

### 5. Common Validations
- Ports don't conflict with HA core services
- Volume mappings are safe and necessary
- Environment variables follow naming conventions
- Startup dependencies are correct
- Architecture support matches Dockerfile
- Version follows semantic versioning

## Usage Examples

### Validate existing config
```bash
/skills run ha-addon-skill --validate-config
```

### Generate config template
```bash
/skills run ha-addon-skill --generate-template
```

### Check addon structure
```bash
/skills run ha-addon-skill --check-structure
```

### Full addon validation
```bash
/skills run ha-addon-skill --full-check
```

## Output Format

JSON structure:
```json
{
  "config_yaml": {
    "status": "valid|invalid",
    "errors": [],
    "warnings": [
      "Consider adding 'homeassistant_api: false' if not using HA API"
    ],
    "required_fields_missing": []
  },
  "build_json": {
    "status": "valid|invalid",
    "errors": [],
    "base_images_match_arch": true
  },
  "documentation": {
    "docs_md_exists": true,
    "readme_exists": true,
    "changelog_exists": false
  },
  "structure": {
    "icon_exists": false,
    "run_script_exists": true,
    "dockerfile_exists": true
  },
  "recommendations": [
    "Add icon.png (256x256) for better visibility in addon store",
    "Create CHANGELOG.md to track version history",
    "Consider adding logo.png for branding"
  ]
}
```

## Home Assistant Addon Best Practices

### Security
- Run as non-root user when possible
- Minimize privileged access
- Use specific base images, not 'latest'
- Validate all user inputs from options
- Don't expose sensitive data in logs

### Performance
- Use multi-stage builds to reduce image size
- Clean up package manager caches
- Combine RUN commands to reduce layers
- Use .dockerignore to exclude unnecessary files

### User Experience
- Provide clear, helpful documentation
- Use meaningful default values
- Validate configuration on startup
- Provide helpful error messages
- Support multiple architectures

### MQTT Integration
For addons using MQTT (like VevorWeatherbridge):
- Use HA's built-in MQTT broker when possible
- Support MQTT auto-discovery
- Group related sensors under single device
- Use retain flags appropriately
- Handle connection failures gracefully

## Reference Documentation

- **HA Addon Development:** https://developers.home-assistant.io/docs/add-ons
- **Config Reference:** https://developers.home-assistant.io/docs/add-ons/configuration
- **MQTT Discovery:** https://www.home-assistant.io/integrations/mqtt/#mqtt-discovery
- **Base Images:** https://github.com/home-assistant/docker-base

## Model Recommendation
Use **Haiku** for config validation, **Sonnet** for structural recommendations and template generation.
