#!/usr/bin/env python3
"""Fetch and analyze Quay.io/Clair vulnerability report for an image.

Usage:
  QUAY_API_TOKEN=... python3 clair-report.py <repo> <manifest_digest>
  python3 clair-report.py <clair-report.json>

If called with repo+digest, fetches from Quay API.
If called with a file path, reads local report.
"""
import json, sys, os, urllib.request

HOST = "https://quay.io"

def fetch(repo, digest):
    token = os.environ.get("QUAY_API_TOKEN", "")
    if not token:
        print("ERROR: QUAY_API_TOKEN not set", file=sys.stderr)
        sys.exit(1)
    url = f"{HOST}/api/v1/repository/{repo}/manifest/{digest}/security"
    req = urllib.request.Request(url, headers={
        "Authorization": f"Bearer {token}", "Accept": "application/json"
    })
    with urllib.request.urlopen(req) as resp:
        return json.loads(resp.read())

def load(path):
    with open(path) as f:
        return json.load(f)

def classify(pkg_name, ns_name):
    """Classify a package as controllable by us or baked into base image."""
    base_layers = {"golang:1.26-alpine", "ghcr.io/home-assistant/", "ha-base"}
    if not ns_name or ns_name.startswith("alpine-"):
        return "os_base"
    if "stdlib" in pkg_name and ns_name == "osv/go":
        return "go_stdlib_builder"
    if "golang.org/x/crypto" in pkg_name and "v0.50" in pkg_name:
        return "go_transitive_fixed"
    if "golang.org/" in pkg_name or "go_" in ns_name:
        return "go_transitive"
    return "other"

def report(data):
    features = data.get("data", {}).get("Layer", {}).get("Features", [])
    total = 0
    by_sev = {}
    buckets = {"go_stdlib_builder": [], "go_transitive": [],
               "go_transitive_fixed": [], "os_base": [], "other": []}

    for feat in features:
        pname = feat.get("Name", "")
        ns = feat.get("NamespaceName", "")
        ver = feat.get("Version", "")
        for v in feat.get("Vulnerabilities", []):
            total += 1
            sev = v.get("Severity", "Unknown")
            if isinstance(sev, dict): sev = sev.get("Severity", "Unknown")
            by_sev[sev] = by_sev.get(sev, 0) + 1
            c = classify(pname, ns)
            buckets.setdefault(c, []).append({
                "name": v.get("Name", ""), "severity": sev,
                "pkg": f"{pname}@{ver}", "fix": v.get("FixedBy", "N/A"),
                "desc": v.get("Description", "")[:80]
            })

    print(f"\n{'='*60}")
    print(f"  VULNERABILITY REPORT")
    print(f"{'='*60}")
    print(f"\nTotal CVEs: {total}")
    print(f"Severity:  {json.dumps(dict(sorted((k,v) for k,v in by_sev.items() if k != 'Unknown')))}")
    print(f"Unknown:   {by_sev.get('Unknown', 0)} (no CVSS score assigned)")
    print()

    sections = [
        ("GO BUILDER STDLIB (golang:1.26-alpine) — fixable via digest update",
         buckets["go_stdlib_builder"], True),
        ("GO TRANSITIVE DEPS (our code) — fixable via go get",
         buckets["go_transitive"], True),
        ("OS PACKAGES (Alpine base) — fixable via Alpine update",
         buckets["os_base"], True),
        ("GO TRANSITIVE (already at fixed version)",
         buckets["go_transitive_fixed"], False),
    ]

    for title, items, show_table in sections:
        if not items:
            continue
        print(f"\n--- {title} ({len(items)} CVEs) ---")
        if not show_table:
            continue
        for c in sorted(items, key=lambda x: x["severity"], reverse=True):
            print(f"  {c['name']:25s} {c['severity']:8s} in {c['pkg'][:40]:40s} → fix {c['fix']}")
            if c['desc']:
                print(f"  {'':25s} {c['desc'][:70]}")

    print(f"\n{'='*60}")
    print(f"  Actionable fixes:")
    if buckets["go_transitive"]:
        for c in buckets["go_transitive"]:
            print(f"    go get {c['pkg'].split('@')[0]}@{c['fix']}")
    if buckets["go_stdlib_builder"]:
        fixes = sorted(set(c['fix'] for c in buckets["go_stdlib_builder"]))
        print(f"    Update Go builder (pinned SHA): needs golang/tags {' or '.join(fixes)}")
        print(f"    Renovate will auto-update SHA pin when new Go patch is released")
    if buckets["os_base"]:
        fixes = sorted(set(c['fix'] for c in buckets["os_base"]))
        print(f"    Alpine packages need: {' '.join(fixes)}")
        print(f"    Nightly rebuild with --pull gets latest once Alpine releases fixes")

if __name__ == "__main__":
    if len(sys.argv) == 3:
        repo, digest = sys.argv[1], sys.argv[2]
        data = fetch(repo, digest)
    elif len(sys.argv) == 2:
        data = load(sys.argv[1])
    else:
        print(f"Usage: {sys.argv[0]} <repo> <digest>  |  {sys.argv[0]} <file.json>")
        sys.exit(1)
    report(data)
