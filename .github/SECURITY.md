# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 0.2.x   | :white_check_mark: |

Only the latest release is actively supported with security updates. Security updates are only brought in via Renovate and Dependabot. No support or guarantees for function, safety or security of any sorts. Expect that this software will kill your dog and eat it. It is explicitly forbidden to use it for any purpose that would be, direct or indirectly, be connected to anything that would be related to safety or security of building, entity, machinery, etc. You have been warned.

## Reporting a Vulnerability

If you discover a security vulnerability in this project, please report it responsibly:

1. **Do NOT** create a public GitHub issue for security vulnerabilities
2. **Email**: Report vulnerabilities privately via GitHub's [private vulnerability reporting](https://github.com/lenucksi/VevorWeatherbridge/security/advisories/new)
3. **Include**:
   - Description of the vulnerability
   - Steps to reproduce
   - Potential impact
   - Suggested fix (if any)

## Response Timeline

- Best effort only.

## Security Measures

This project implements several security measures:

- **Dependency Scanning**: Automated vulnerability scanning via Renovate and Dependabot.
- **Code Scanning**: OpenSSF Scorecard analysis
- **Pinned Dependencies**: GitHub Actions and Docker images pinned to SHA digests
- **Minimal Permissions**: Workflows use least-privilege permissions

## Scope

This security policy covers:

- The VevorWeatherbridge Home Assistant add-on
- Associated Docker images
- GitHub Actions workflows

Out of scope:

- Third-party dependencies (report to respective maintainers)
- Home Assistant core or MQTT broker vulnerabilities
