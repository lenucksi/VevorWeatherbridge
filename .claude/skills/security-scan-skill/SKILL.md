# security-scan-skill

## Role
Centralized SAST and security scanning orchestration for the VevorWeatherbridge project.

## Triggers
- Manual: `/skills run security-scan-skill`
- Delegated: Called by `python-ci-skill` for security focus
- Hook-based: Suggested when `requirements.txt` is modified

## Tools Required

### Primary SAST
- `semgrep` - Pattern-based security scanning (supports custom rules)
- `bandit` - Python AST-based security issue detection

### Dependency Security
- `pip-audit` - Python package vulnerability scanner
- `safety` - Alternative dependency checker (optional)

### Secret Detection
- `trufflehog` - Git history secret scanner (optional)
- Built-in regex patterns for common secrets

### Container Security
- `trivy` - Container vulnerability scanner (for Docker images)
- `grype` - Alternative container scanner (optional)

## Responsibilities

### 1. Python Security Scanning
- Run `bandit` with focus on high/critical severity issues
- Check for:
  - SQL injection vulnerabilities
  - Command injection (shell=True usage)
  - Hardcoded passwords/secrets
  - Insecure deserialization
  - Weak cryptography
  - Path traversal
  - XSS vulnerabilities (Flask templates)

### 2. SAST with Semgrep
- Use Semgrep rules for Python security patterns
- Check against OWASP Top 10 patterns
- Custom rules for Flask security:
  - Debug mode in production
  - Missing CSRF protection
  - Unsafe redirect usage
  - SQL query construction

### 3. Dependency Vulnerability Scanning
- Run `pip-audit` against `requirements.txt`
- Identify known CVEs in dependencies
- Suggest version upgrades or patches
- Check for outdated packages with security issues

### 4. Secret Detection
- Scan for exposed credentials:
  - API keys
  - Passwords
  - Tokens
  - Private keys
  - Database connection strings
- Check environment variable usage vs hardcoded secrets

### 5. Docker Security
- Scan base images for vulnerabilities
- Check for:
  - Running as root user
  - Exposed sensitive ports
  - Unnecessary capabilities
  - Outdated base images

## Usage Examples

### Full security scan
```bash
/skills run security-scan-skill
```

### Focus on dependencies only
```bash
/skills run security-scan-skill --focus dependencies
```

### Focus on code SAST
```bash
/skills run security-scan-skill --focus sast
```

### Scan Docker image
```bash
/skills run security-scan-skill --scan-image weatherstation:latest
```

### Check for secrets
```bash
/skills run security-scan-skill --focus secrets
```

## Output Format

JSON structure:
```json
{
  "summary": {
    "critical": 0,
    "high": 2,
    "medium": 5,
    "low": 8,
    "info": 12
  },
  "python_security": {
    "bandit": {
      "issues_found": 7,
      "high_severity": [
        {
          "file": "weatherstation.py",
          "line": 148,
          "issue": "Use of requests without timeout",
          "severity": "HIGH",
          "cwe": "CWE-400",
          "fix_suggestion": "Add timeout parameter to requests.get()"
        }
      ]
    },
    "semgrep": {
      "rules_matched": 3,
      "findings": [...]
    }
  },
  "dependencies": {
    "pip_audit": {
      "vulnerable_packages": 1,
      "details": [
        {
          "package": "requests",
          "version": "2.28.0",
          "vulnerability": "CVE-2023-xxxxx",
          "fixed_in": "2.31.0",
          "severity": "HIGH"
        }
      ]
    }
  },
  "secrets": {
    "found": 0,
    "potential_issues": []
  },
  "container": {
    "trivy": {
      "vulnerabilities": {
        "critical": 0,
        "high": 3,
        "medium": 12
      }
    }
  },
  "recommendations": [
    "Upgrade requests to 2.31.0 to fix CVE-2023-xxxxx",
    "Add timeout to all HTTP requests to prevent hanging",
    "Consider using connection pooling for MQTT client",
    "Review error handling to avoid information disclosure"
  ]
}
```

## Security Best Practices for This Project

### Flask Security
- Never run with `debug=True` in production
- Validate all input from weather station
- Sanitize data before MQTT publishing
- Use environment variables for sensitive config
- Implement rate limiting for endpoint

### MQTT Security
- Use TLS for MQTT connections in production
- Validate MQTT broker certificates
- Use strong authentication credentials
- Restrict MQTT topic permissions
- Avoid publishing sensitive data in retained messages

### Docker Security
- Use specific base image versions (not 'latest')
- Run container as non-root user
- Minimize installed packages
- Use multi-stage builds
- Scan images regularly
- Keep base images updated

### Dependency Management
- Pin dependency versions in requirements.txt
- Regularly update dependencies
- Review security advisories
- Use virtual environments
- Audit transitive dependencies

### Data Handling
- Validate weather station data format
- Sanitize before forwarding to Weather Underground
- Don't log sensitive credentials
- Handle DNS resolution securely
- Timeout external requests

## OWASP Top 10 Coverage

1. **A01:2021 - Broken Access Control**: Check MQTT topic permissions
2. **A02:2021 - Cryptographic Failures**: Validate MQTT TLS usage
3. **A03:2021 - Injection**: Check SQL/command injection in Flask
4. **A04:2021 - Insecure Design**: Review architecture for security flaws
5. **A05:2021 - Security Misconfiguration**: Check Flask debug mode, Docker config
6. **A06:2021 - Vulnerable Components**: Scan dependencies with pip-audit
7. **A07:2021 - Auth Failures**: Review MQTT authentication
8. **A08:2021 - Software/Data Integrity**: Check package integrity
9. **A09:2021 - Security Logging**: Ensure proper logging without secrets
10. **A10:2021 - SSRF**: Validate Weather Underground forwarding

## Reference Documentation

- **Semgrep Rules:** https://semgrep.dev/explore
- **Bandit Docs:** https://bandit.readthedocs.io/
- **pip-audit:** https://github.com/pypa/pip-audit
- **OWASP Top 10:** https://owasp.org/Top10/
- **Docker Security:** https://docs.docker.com/engine/security/
- **Flask Security:** https://flask.palletsprojects.com/en/latest/security/

## Model Recommendation
Use **Haiku** for basic vulnerability scanning, **Sonnet** for security remediation and architectural recommendations.
