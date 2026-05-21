#!/usr/bin/env python3
import json, sys, os
report_file = sys.argv[1] if len(sys.argv) > 1 else "clair-report.json"
try:
    with open(report_file) as f:
        data = json.load(f)
    pkgs = data.get("packages", {}).get("packages", [])
    total = sum(len(p.get("Vulnerabilities", [])) for p in pkgs)
    print(f"Total CVEs: {total}")
    for p in pkgs:
        for v in p.get("Vulnerabilities", []):
            sev = v.get("Severity", "Unknown")
            if sev in ("Critical", "High"):
                print(f"  {v['Name']} [{sev}] in {p['Name']} fix: {v.get('FixedBy','N/A')}")
except Exception as e:
    print(f"CVE report not available: {e}")
