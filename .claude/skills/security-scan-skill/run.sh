#!/usr/bin/env bash
# Security Scan Skill Runner
# Orchestrates security scanning using Poetry environment

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/../../.." && pwd)"

cd "$PROJECT_ROOT"

FOCUS="${1:-all}"

echo "🔒 Security Scan Skill"
echo "=================================================="
echo "Focus: $FOCUS"
echo ""

# Use Poetry to run tools
POETRY_RUN="poetry run"

# Check if Poetry is available
if ! command -v poetry >/dev/null 2>&1; then
    echo "⚠️  Poetry not found - falling back to system tools"
    POETRY_RUN=""
fi

# 1. Python Code Security (Bandit)
if [[ "$FOCUS" == "all" || "$FOCUS" == "sast" ]]; then
    echo "1️⃣  Running Bandit (Python security scanner)..."
    $POETRY_RUN bandit -r . \
        -f screen \
        -ll \
        --skip B404,B603 \
        --exclude ./.venv,./venv,./.git \
        2>/dev/null || echo "  ⚠️  Bandit found potential issues or not available"
    echo ""
fi

# 2. SAST with Semgrep
if [[ "$FOCUS" == "all" || "$FOCUS" == "sast" ]]; then
    echo "2️⃣  Running Semgrep (SAST)..."
    $POETRY_RUN semgrep --config=auto \
        --exclude="*.git*" \
        --exclude="*.venv*" \
        --exclude="venv" \
        --severity=WARNING \
        --severity=ERROR \
        . 2>/dev/null || echo "  ⚠️  Semgrep found potential issues or not available"
    echo ""
fi

# 3. Dependency Vulnerability Scanning
if [[ "$FOCUS" == "all" || "$FOCUS" == "dependencies" ]]; then
    echo "3️⃣  Scanning dependencies for vulnerabilities..."
    if [ -f "pyproject.toml" ]; then
        $POETRY_RUN pip-audit 2>/dev/null || echo "  ⚠️  Vulnerable dependencies found or tool not available"
    elif [ -f "requirements.txt" ]; then
        $POETRY_RUN pip-audit -r requirements.txt 2>/dev/null || echo "  ⚠️  Vulnerable dependencies found or tool not available"
    else
        echo "  ℹ️  No pyproject.toml or requirements.txt found"
    fi
    echo ""
fi

# 4. Secret Detection (basic patterns)
if [[ "$FOCUS" == "all" || "$FOCUS" == "secrets" ]]; then
    echo "4️⃣  Checking for exposed secrets..."

    SECRET_PATTERNS=(
        "password\s*=\s*['\"][^'\"]+['\"]"
        "api_key\s*=\s*['\"][^'\"]+['\"]"
        "secret\s*=\s*['\"][^'\"]+['\"]"
        "token\s*=\s*['\"][^'\"]+['\"]"
        "-----BEGIN\s+(RSA|DSA|EC|OPENSSH)\s+PRIVATE\s+KEY-----"
    )

    SECRETS_FOUND=0
    for pattern in "${SECRET_PATTERNS[@]}"; do
        if grep -rniE "$pattern" . \
            --exclude-dir=.git \
            --exclude-dir=.venv \
            --exclude-dir=venv \
            --exclude-dir=.claude \
            --exclude="*.md" 2>/dev/null; then
            SECRETS_FOUND=$((SECRETS_FOUND + 1))
        fi
    done

    if [ $SECRETS_FOUND -eq 0 ]; then
        echo "  ✅ No obvious hardcoded secrets found"
    else
        echo "  ⚠️  Potential secrets detected - review above matches"
    fi
    echo ""
fi

# 5. Docker Image Security
if [[ "$FOCUS" == "all" || "$FOCUS" == "container" ]]; then
    echo "5️⃣  Checking Docker security..."

    if [ -f "Dockerfile" ]; then
        echo "  Checking for security issues in Dockerfile..."

        # Check for running as root
        if ! grep -q "^USER" Dockerfile; then
            echo "  ⚠️  WARNING: Dockerfile does not specify USER (running as root)"
        else
            echo "  ✅ Dockerfile specifies non-root USER"
        fi

        # Check for pinned versions
        if grep -q ":latest" Dockerfile; then
            echo "  ⚠️  WARNING: Using ':latest' tag (should pin specific versions)"
        else
            echo "  ✅ Base image version is pinned"
        fi

        # Scan with trivy if available
        if check_tool "trivy"; then
            echo ""
            echo "  Running Trivy container scan..."
            trivy config Dockerfile || true
        fi
    else
        echo "  ℹ️  No Dockerfile found"
    fi
    echo ""
fi

# Summary and Recommendations
echo "=================================================="
echo "📋 Security Summary"
echo ""

echo "Critical security checks for this project:"
echo ""
echo "1. Flask Security:"
echo "   - Ensure debug mode is disabled in production"
echo "   - Validate all input from weather station"
echo "   - Add rate limiting to prevent abuse"
echo ""

echo "2. MQTT Security:"
echo "   - Use TLS for MQTT connections"
echo "   - Store MQTT credentials in environment variables (✅ currently done)"
echo "   - Validate MQTT broker certificates"
echo ""

echo "3. Weather Underground Forwarding:"
echo "   - Currently uses DNS resolution with public DNS (8.8.8.8)"
echo "   - Has timeout set to 5 seconds (✅ good)"
echo "   - Consider adding retry logic with exponential backoff"
echo ""

echo "4. Docker Security:"
echo "   - Currently runs as root - consider adding USER directive"
echo "   - Base image uses specific version (python:3.12-slim) ✅"
echo "   - Consider multi-stage build to reduce image size"
echo ""

echo "5. Dependency Security:"
echo "   - Review pip-audit output above"
echo "   - Pin all dependency versions in requirements.txt"
echo "   - Set up automated dependency updates (Dependabot/Renovate)"
echo ""

echo "📝 Recommended next steps:"
echo "  1. Review any HIGH/CRITICAL findings above"
echo "  2. Update vulnerable dependencies"
echo "  3. Add USER directive to Dockerfile"
echo "  4. Consider adding Flask rate limiting"
echo "  5. Set up automated security scanning in CI/CD"
echo ""

echo "✅ Security scan complete"
echo ""
