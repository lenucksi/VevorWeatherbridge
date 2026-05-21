#!/usr/bin/env python3
"""Fetch and analyze Quay.io security report for an image.

Usage (fetch from Quay API):
  QUAY_USER=... QUAY_TOKEN=... python3 clair-report.py <namespace/repo> <manifest_digest>

Usage (read local file):
  python3 clair-report.py <report.json>
"""
import json, sys, os, urllib.request, base64

HOST = "https://quay.io"

def fetch(repo, digest):
    user = os.environ.get("QUAY_USER", "")
    token = os.environ.get("QUAY_TOKEN", "")
    if not user or not token:
        print("ERROR: QUAY_USER and QUAY_TOKEN env vars required", file=sys.stderr)
        sys.exit(1)
    url = f"{HOST}/api/v1/repository/{repo}/manifest/{digest}/security"
    auth = base64.b64encode(f"{user}:{token}".encode()).decode()
    req = urllib.request.Request(url, headers={
        "Authorization": f"Basic {auth}", "Accept": "application/json"
    })
    with urllib.request.urlopen(req) as resp:
        return json.loads(resp.read())

def load(path):
    with open(path) as f:
        return json.load(f)

def classify(pname, ns, ver):
    if not ns or ns.startswith("alpine-"):
        return "os_alpine"
    if "stdlib" in pname and ns == "osv/go":
        # Check whether it's builder Go or HA base Go
        return "go_stdlib_builder" if "1.26" in ver else "go_stdlib_base"
    if "golang.org/" in pname:
        fix_ver = ver.lstrip("v")
        # x/crypto@v0.50.0 is already above all fix versions from the report
        if pname == "golang.org/x/crypto" and fix_ver >= "0.50":
            return "go_transitive_fixed"
        return "go_transitive"
    return "other"

def report(data):
    features = data.get("data", {}).get("Layer", {}).get("Features", [])
    total = 0
    by_sev = {}
    buckets = {}

    for feat in features:
        pname = feat.get("Name", "")
        ns = feat.get("NamespaceName", "")
        ver = feat.get("Version", "")
        for v in feat.get("Vulnerabilities", []):
            total += 1
            sev = v.get("Severity", "Unknown")
            if isinstance(sev, dict): sev = sev.get("Severity", "Unknown")
            by_sev[sev] = by_sev.get(sev, 0) + 1
            c = classify(pname, ns, ver)
            buckets.setdefault(c, []).append({
                "name": v.get("Name", ""), "severity": sev,
                "pkg": f"{pname}@{ver}", "fix": v.get("FixedBy", "N/A"),
                "desc": v.get("Description", "")[:80]
            })

    print(f"\n{'='*60}")
    print(f"  QUAY SECURITY REPORT")
    print(f"{'='*60}")
    print(f"\nTotal CVEs: {total}")
    sev_known = {k: v for k, v in sorted(by_sev.items()) if k != "Unknown"}
    if sev_known:
        print(f"Known severity: {json.dumps(sev_known)}")
    print(f"Unknown severity: {by_sev.get('Unknown', 0)} (no CVSS score)")
    print()

    sections = [
        ("GO BUILDER STDLIB (golang:1.26-alpine) — fix when Renovate updates SHA",
         "go_stdlib_builder", True),
        ("GO STDLIB IN HA BASE IMAGE — fix when HA rebuilds their base",
         "go_stdlib_base", True),
        ("GO TRANSITIVE DEPS (our go.mod) — fixable via go get",
         "go_transitive", True),
        ("ALPINE OS PACKAGES — fix via nightly rebuild with --pull",
         "os_alpine", True),
        ("ALREADY FIXED (version > fix)",
         "go_transitive_fixed", False),
    ]

    for title, key, show_table in sections:
        items = buckets.get(key, [])
        if not items:
            continue
        print(f"\n── {title} ({len(items)} CVEs) ──")
        if not show_table:
            continue
        for c in sorted(items, key=lambda x: x["severity"], reverse=True):
            print(f"  {c['name']:30s} {c['severity']:8s} in {c['pkg'][:45]:45s} → fix {c['fix']}")

    print(f"\n{'='*60}")
    print(f"  ACTIONABLE:")
    for c in buckets.get("go_transitive", []):
        print(f"    go get {c['pkg'].split('@')[0]}@{c['fix']}")
    if buckets.get("go_stdlib_builder"):
        fixes = sorted(set(c['fix'] for c in buckets["go_stdlib_builder"]))
        print(f"    Wait for Go {'/'.join(fixes)} release → Renovate updates SHA")
    if buckets.get("os_alpine"):
        print(f"    Nightly rebuild with --pull (when Alpine releases fixes)")
    if not any(buckets.get(k) for k in ("go_stdlib_builder","go_transitive","os_alpine")):
        print(f"    Nothing actionable — all CVEs are in HA base image")

if __name__ == "__main__":
    if len(sys.argv) == 3:
        repo, digest = sys.argv[1], sys.argv[2]
        data = fetch(repo, digest)
    elif len(sys.argv) == 2:
        data = load(sys.argv[1])
    else:
        print(f"Usage: {sys.argv[0]} <namespace/repo> <digest> | {sys.argv[0]} <file.json>")
        sys.exit(1)
    report(data)
