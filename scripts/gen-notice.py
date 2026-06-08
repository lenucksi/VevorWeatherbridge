#!/usr/bin/env python3
# SPDX-License-Identifier: GPL-3.0-or-later
# Generate AboutCode inventory JSON from Syft CycloneDX SBOM + Go module metadata.
#
# Filters file-level noise (zoneinfo, terminfo, CA certs), groups Alpine apk
# packages and Go module dependencies, and merges license info from go-licenses.

import json
import subprocess
import sys
import os
import urllib.parse

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
PROJECT_ROOT = os.path.abspath(os.path.join(SCRIPT_DIR, ".."))

SBOM_PATH = os.environ.get("SBOM_PATH") or os.path.join(PROJECT_ROOT, "sbom-syft.cdx.json")
GO_MOD_DIR = os.path.join(PROJECT_ROOT, "vevor-weatherbridge-go")
OUTPUT_PATH = "/tmp/about-inventory.json"

# File paths to exclude from the notice (system data, not licensed components)
EXCLUDE_PATH_PREFIXES = (
    "/etc/terminfo/",
    "/usr/share/terminfo/",
    "/usr/share/zoneinfo/",
    "/usr/share/ca-certificates/",
    "/etc/ssl/certs/",
    "/etc/apk/keys/",
    "/usr/share/apk/keys/",
    "/usr/lib/python",
    "/lib/apk/",
)

def load_sbom(path):
    with open(path) as f:
        return json.load(f)

def get_go_module_info():
    """Run go list -m -json all and return a dict of module path -> info."""
    result = subprocess.run(
        ["go", "list", "-m", "-json", "all"],
        capture_output=True, text=True,
        cwd=GO_MOD_DIR,
    )
    if result.returncode != 0:
        print(f"Warning: go list failed: {result.stderr}", file=sys.stderr)
        return {}
    modules = {}
    for line in result.stdout.strip().split("\n}\n{"):
        if not line.strip():
            continue
        entry = "{" + line + "}" if not line.startswith("{") else line
        if entry.endswith("}") or not entry.endswith("}"):
            try:
                mod = json.loads(entry)
                modules[mod.get("Path", "")] = mod
            except json.JSONDecodeError:
                pass
    return modules

def get_go_licenses():
    """Run go-licenses csv and return {module_path: {license, url}}."""
    result = subprocess.run(
        ["go-licenses", "csv", "./..."],
        capture_output=True, text=True,
        cwd=GO_MOD_DIR,
    )
    if result.returncode not in (0, 1):
        print(f"Warning: go-licenses csv failed: {result.stderr}", file=sys.stderr)
        return {}
    licenses = {}
    for line in result.stdout.strip().split("\n"):
        parts = line.split(",")
        if len(parts) >= 3:
            mod_path, url, lic = parts[0], parts[1], parts[2]
            licenses[mod_path] = {"url": url, "license": lic}
    return licenses

# Known copyright holders for Go dependencies (from project metadata)
GO_COPYRIGHT = {
    "github.com/eclipse/paho.mqtt.golang":     "Copyright (c) Eclipse Contributors",
    "github.com/gorilla/websocket":             "Copyright (c) 2013 The Gorilla WebSocket Authors",
    "golang.org/x/net":                         "Copyright (c) 2009 The Go Authors",
    "golang.org/x/sync":                        "Copyright (c) 2009 The Go Authors",
    "golang.org/x/crypto":                      "Copyright (c) 2009 The Go Authors",
    "dario.cat/mergo":                          "Copyright (c) 2022 Dario Castañé",
    "github.com/Masterminds/goutils":           "Copyright (c) 2014 Masterminds",
    "github.com/Masterminds/semver/v3":         "Copyright (c) 2018 Masterminds",
    "github.com/Masterminds/sprig/v3":          "Copyright (c) 2017 Masterminds",
    "github.com/google/uuid":                   "Copyright (c) 2021 Google LLC",
    "github.com/home-assistant/tempio":         "Copyright (c) Home Assistant Contributors",
    "github.com/huandu/xstrings":               "Copyright (c) 2021 Huan Du",
    "github.com/mitchellh/copystructure":       "Copyright (c) 2018 Mitchell Hashimoto",
    "github.com/mitchellh/reflectwalk":         "Copyright (c) 2013 Mitchell Hashimoto",
    "github.com/shopspring/decimal":            "Copyright (c) 2015 Spring, Inc.",
    "github.com/spf13/cast":                    "Copyright (c) 2021 Steve Francia",
    "stdlib":                                   "Copyright (c) 2009 The Go Authors",
}

# Same subpackages that go-licenses may report differently
GO_LICENSE_ALIASES = {
    "golang.org/x/sync":                   "golang.org/x/sync/semaphore",
}

GO_LICENSES = {
    "dario.cat/mergo":                    "BSD-3-Clause",
    "github.com/Masterminds/goutils":     "Apache-2.0",
    "github.com/Masterminds/semver/v3":   "Apache-2.0",
    "github.com/Masterminds/sprig/v3":    "MIT",
    "github.com/google/uuid":             "BSD-3-Clause",
    "github.com/home-assistant/tempio":   "MIT",
    "github.com/huandu/xstrings":         "MIT",
    "github.com/mitchellh/copystructure": "MIT",
    "github.com/mitchellh/reflectwalk":   "MIT",
    "github.com/shopspring/decimal":      "MIT",
    "github.com/spf13/cast":              "MIT",
    "golang.org/x/crypto":                "BSD-3-Clause",
    "golang.org/x/sync":                  "BSD-3-Clause",
}

# Alpine package publisher -> copyright mapping
ALPINE_COPYRIGHT = {
    "musl":      "Copyright (c) 2005-2025 Rich Felker, et al.",
    "busybox":   "Copyright (c) 1999-2025 Erik Andersen, Rob Landley, Denys Vlasenko, et al.",
    "openssl":   "Copyright (c) 1998-2025 The OpenSSL Project",
    "zlib":      "Copyright (c) 1995-2025 Jean-loup Gailly and Mark Adler",
    "libuv":     "Copyright (c) 2015-2025 libuv contributors",
    "curl":      "Copyright (c) 1996-2025 Daniel Stenberg",
    "bash":      "Copyright (c) 1989-2025 Free Software Foundation, Inc.",
    "alpine":    "Copyright (c) Alpine Linux Development Team",
}

# Alpine packages missing SPDX in SBOM -> manual license mapping
ALPINE_LICENSES = {
    "ca-certificates":        "MPL-2.0",
    "ca-certificates-bundle": "MPL-2.0",
    "keyutils-libs":          "GPL-2.0-or-later",
    "libcom_err":             "MIT",
    "libgcc":                 "GPL-3.0-or-later WITH GCC-exception-3.1",
    "libidn2":                "GPL-2.0-or-later",
    "libstdc++":              "GPL-3.0-or-later WITH GCC-exception-3.1",
    "libunistring":           "LGPL-3.0-or-later",
    "musl-utils":             "MIT",
    "tzdata":                 "Public-Domain",
    "zstd-libs":              "BSD-3-Clause",
    "libstdc++":              "GPL-3.0-or-later WITH GCC-exception-3.1",
}

# Project's own module path — filter out of notice
OWN_MODULE = "github.com/lenucksi/vevor-weatherbridge-go"

def is_noise_component(c):
    """Check if a component is noise (file-level or system data)."""
    ctype = c.get("type", "")
    name = c.get("name", "")
    purl = c.get("purl", "")
    if purl and (purl.startswith("pkg:apk") or purl.startswith("pkg:golang")):
        return False
    if ctype == "library" and not name.startswith("/"):
        return False
    if ctype == "operating-system":
        return True
    if name.startswith(EXCLUDE_PATH_PREFIXES):
        return True
    if name.startswith("/"):
        return True
    return False

def build_inventory(sbom, go_modules, go_licenses_map):
    inventory = []

    # Process library components from SBOM (APK packages + Go modules)
    seen_purls = set()
    for c in sbom.get("components", []):
        if is_noise_component(c):
            continue
        purl = c.get("purl", "")
        if purl in seen_purls:
            continue
        seen_purls.add(purl)

        name = c.get("name", "")
        version = c.get("version", "")
        sbom_licenses = []
        for l in c.get("licenses", []):
            lid = l.get("license", {}).get("id", "")
            if lid:
                sbom_licenses.append(lid)
        license_expr = " AND ".join(sbom_licenses) if sbom_licenses else ""

        copyright_str = ""

        if purl.startswith("pkg:apk"):
            pkg_name = urllib.parse.unquote(purl.split("/")[-1].split("@")[0])
            if not license_expr:
                for pfx, lic in ALPINE_LICENSES.items():
                    if pkg_name.startswith(pfx):
                        license_expr = lic
                        break
            for prefix, cr in ALPINE_COPYRIGHT.items():
                if pkg_name.startswith(prefix):
                    copyright_str = cr
                    break
            if not copyright_str:
                copyright_str = f"Copyright (c) Alpine Linux Contributors"

        elif purl.startswith("pkg:golang"):
            go_path = name
            if go_path in go_licenses_map:
                gl = go_licenses_map[go_path]
                if not license_expr and gl["license"] != "Unknown":
                    license_expr = gl["license"]
            if not license_expr and go_path in GO_LICENSES:
                license_expr = GO_LICENSES[go_path]
            if not license_expr:
                alias = GO_LICENSE_ALIASES.get(go_path)
                if alias and alias in go_licenses_map:
                    al = go_licenses_map[alias]
                    if al["license"] != "Unknown":
                        license_expr = al["license"]
            if go_path in GO_COPYRIGHT:
                copyright_str = GO_COPYRIGHT[go_path]
            if go_path == OWN_MODULE:
                continue

        entry = {
            "name": name,
            "version": version,
            "license_expression": license_expr,
            "copyright": copyright_str,
        }
        if purl:
            entry["purl"] = purl
        inventory.append(entry)

    return inventory

def main():
    sbom = load_sbom(SBOM_PATH)
    go_modules = get_go_module_info()
    go_licenses_map = get_go_licenses()

    inventory = build_inventory(sbom, go_modules, go_licenses_map)

    with open(OUTPUT_PATH, "w") as f:
        json.dump(inventory, f, indent=2)

    print(f"Generated {OUTPUT_PATH} with {len(inventory)} components")
    sys.exit(0)

if __name__ == "__main__":
    main()
