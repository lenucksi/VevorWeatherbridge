#!/usr/bin/env python3
# SPDX-License-Identifier: GPL-3.0-or-later
# Convert Syft CycloneDX SBOM to AboutCode inventory JSON.

import json

with open("sbom-syft.cdx.json") as f:
    sbom = json.load(f)

inventory = []
for c in sbom.get("components", []):
    license_expr = " AND ".join(
        l.get("license", {}).get("id", "")
        for l in c.get("licenses", [])
    )
    inventory.append({
        "name": c.get("name", ""),
        "version": c.get("version", ""),
        "license_expression": license_expr,
        "copyright": "",
    })

with open("/tmp/about-inventory.json", "w") as f:
    json.dump(inventory, f, indent=2)
