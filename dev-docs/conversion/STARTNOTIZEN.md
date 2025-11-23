Heute bist du ein Experte in AI LLM based coding mit Claude sowie dem folgenden Techstack:

- Claude Code, insbesondere die Erweiterungsmechanismen wie Hooks, Skills, Subagents und MCPs. Ebenso Promptwriting für Claude Code dass sicherstellt dass die Fähikeiten die den genannten Funktionen eingebaut wurden benutzt werden.
- Homeassistant und das dortige Framework für Addons
- Python, Docker und aktuelle Qualitätstechniken, Linter, SAST/DAST Werkzeuge usw.

Du recherchierst für alles lang und tief im Web in 2025er, maximal 2024er Quellen zu den jeweiligen Themen für Best Practices. Du gibtst für ausnahmslos alles Deine Quellen an, immer mindestens 2. Du ziehst absolut immer offizielle Dokumentation der Hersteller vor.

Werte auch die folgenden menschlichen Rechercheresultate durch herunterladen und durchlesen aus in Hinsicht auf Deine folgende Aufgabe:

- Bericht über die Erweiterungsmöglichkeiten und Claude Code <https://alexop.dev/posts/understanding-claude-code-full-stack/>
- Projekt das es erlaubt aus Prompts Hooks zu bauen: <https://www.sourcepulse.org/projects/11549096> und <https://github.com/zxdxjtu/claudecode-rule2hook>
- <https://code.claude.com/docs/en/hooks> + <https://code.claude.com/docs/en/slash-commands> + <https://code.claude.com/docs/en/skills#agent-skills>

Es ist jetzt Deine Aufgabe für den o.g. TechStack alle Linter und Bestpractices aus 2025 zu finden sowie
einen Prompt zu schreiben der Claude Code erlaubt einen möglichst tokeneffizienten und Software+Engineeringqualität sichernden Weg umzusetzen unter bester Nutzung der Erweiterungsmöglichkeiten wie Hooks, Skills, MCPs usw.

# Prompt-Vorlage für Claude Code Harness mit Hooks/Skills/MCP

> Ziel: Claude Code soll **einen Coding-Harness erzeugen**, der automatisch Quality-Checks (Lint, Tests, Security) orchestriert, unter Nutzung der Claude Code-Mechanismen (Hooks, Skills, MCP), statt die Checks selbst auszuführen.

---

**System-Instruktion:**
Du bist ein Entwickler-Agent mit Zugriff auf Claude Code Tools. Deine Aufgabe: Erzeuge **only code, config und Markup** (keine ausführende Checks) – nämlich ein Harness-Gerüst, das künftig automatisiert Qualitätssicherung übernimmt. Verwende dabei:

- Skills (z. B. `python-ci-skill`, `ha-addon-skill`, `security-scan-skill`)
- Hooks (z. B. `.claude/settings.toml` / `.claude/settings.json`)
- MCP-Verweise (`@repo:…`, Tools via MCP)
- Slash-Commands (z. B. `/hooks`, `/skills list`, `/mcp …`)
- Token-effiziente Struktur: keine langen Prosa-Absätze, nur klar strukturierte Artefakte.

---

**Aufgabenstellung:**

1. Lege eine **Ordner- und Dateistruktur** für das Harness an (z. B. `/.claude/hooks/`, `/skills/`, `manifest.md`).
2. Definiere **mindestens drei Skills** (z. B. „python-ci-skill“, „ha-addon-skill“, „security-scan-skill“) mit je SKILL.md-Inhalt (Rolle, Trigger, Tools).
3. Formuliere eine `.claude/settings.toml` (oder JSON) mit **Hooks**, die folgende Automatisierung abbilden:
   - Bei Bearbeitung einer Python-Datei (`*.py`): Nach der `edit_file` Tool-Benutzung → Hook führt automatisiert `ruff --fix` und `pytest` aus (oder markiert Trigger)
   - Bei Bearbeitung eines Dockerfiles (`Dockerfile`): Nach `edit_file` → Hook führt `hadolint` aus
   - Bei Änderungen im `addons/` Verzeichnis (Home Assistant Addon): Trigger Skill „ha-addon-skill“.
4. Füge einen `manifest.md` (oder README.md) hinzu, der beschreibt, wie der Harness genutzt wird: wie Skills aufgerufen werden, wie MCP genutzt wird (z. B. `@repo:/…`), wie Slash-Commands ausgelöst werden.
5. Gebe kurze Code-Snippets zur Verwendung: z. B. wie ein Entwickler im Chat `/skills run python-ci-skill --target branchX` aufruft.
6. Verwende Token-sparsam: keine langen Fließtexte, nur klare Konfigurations– und Codeblöcke.

---

**Output-Schema:**

- Ordner-/Dateiliste
- Inhalte der Dateien (SKILL.md, settings.toml, manifest.md)
- Beispiel-Slash-Command-Aufruf
- Hinweis zur Integration mit MCP (nur Referenz)

Beginne jetzt mit der Ausgabe – nur Markdown, direkt kopierbar.

---
.claude/
  settings.toml
  hooks/
    rule2hook_wrapper.sh
skills/
  python-ci-skill/
    SKILL.md
    run.sh
  ha-addon-skill/
    SKILL.md
    run.sh
  security-scan-skill/
    SKILL.md
    run.sh
manifest.md
prompts/
  rule2hook-prompt.txt
  text2hook-prompt.txt

---

.claude/settings.toml

## Claude Code project hooks config (TOML). Nicht ausführend — nur Konfiguration

[hooks]

## Trigger: nach edit_file mit Python-Datei → markiere work und rufe Hook-Command (Wrapper)

"on.edit_file:*.py" = { command = ".claude/hooks/rule2hook_wrapper.sh --rule \"on edit_file *.py -> run ruff --fix && run pytest\" --label python-edit" }

## Trigger: nach edit_file an Dockerfile

"on.edit_file:Dockerfile" = { command = ".claude/hooks/rule2hook_wrapper.sh --rule \"on edit_file Dockerfile -> run hadolint\" --label dockerfile-edit" }

## Trigger: Änderungen im addons/ Verzeichnis → delegiere an HA skill

"on.edit_file:addons/**" = { command = ".claude/hooks/rule2hook_wrapper.sh --rule \"on edit_file addons/** -> trigger skill ha-addon-skill\" --label ha-addon" }

## Generic "on.write" example to catch Write tool edits (maps to edit_file semantics)

"on.write:*.py" = { command = ".claude/hooks/rule2hook_wrapper.sh --rule \"on write *.py -> run ruff --fix && run pytest\" --label write-python" }

---
.claude/hooks/rule2hook_wrapper.sh

## !/usr/bin/env bash

## Wrapper used by hooks config. This file is part of the harness scaffolding

## It does NOT execute the linters itself; it translates the rule into a Claude Hook stub

## or calls the local claudecode-rule2hook tool to produce final hook artifacts

##

## Usage

## ./rule2hook_wrapper.sh --rule "on edit_file *.py -> run ruff --fix && run pytest" --label python-edit

set -euo pipefail
RULE=""
LABEL=""
while [Link text]( $# -gt 0 )]; do
  case "$1" in
    --rule) RULE="$2"; shift 2;;
    --label) LABEL="$2"; shift 2;;
    *) echo "Unknown arg $1"; exit 2;;
  esac
done

## If claudecode-rule2hook is installed in project, invoke it to generate a hook file

## Otherwise emit a stub file under .claude/generated-hooks/${LABEL}.hook.sh (for manual inspection)

GEN_DIR=".claude/generated-hooks"
mkdir -p "$GEN_DIR"

if command -v claudecode-rule2hook >/dev/null 2>&1; then

## Transform natural language rule into a concrete hook using the community tool

  claudecode-rule2hook --rule "$RULE" --out "$GEN_DIR/${LABEL}.hook.sh"
else

## Emit a conservative stub that documents the intended action (no execution)

  cat > "$GEN_DIR/${LABEL}.hook.sh" <<EOF

## !/usr/bin/env bash

## Generated stub for label: ${LABEL}

## Intended action (do NOT execute automatically)

## ${RULE}

## Recommended actual implementation

## - For ruff: run `ruff --fix <changed-files>`

## - For pytest: run `pytest -q` limited to changed tests

## - For hadolint: run `hadolint Dockerfile`

EOF
  chmod +x "$GEN_DIR/${LABEL}.hook.sh"
fi

echo "Generated hook artifact for label=${LABEL} at ${GEN_DIR}/${LABEL}.hook.sh"

---

skills/python-ci-skill/SKILL.md

## python-ci-skill

Role: Provide Python project quality rules, canonical commands, and patch-generation logic.
Triggers: invoked by Hooks or by `/skills run python-ci-skill --target <ref>`
Tools (expected via MCP or local environment): ruff, pytest, mypy (optional), bandit (optional)
Responsibilities:

- Given a set of changed files, produce:
  1. Suggested unified-diff fix patches for formatting/lint issues (ruff --fix diff).
  2. Test selection plan (which pytest tests to run) and a command stub.
  3. Minimal security scan suggestions (bandit findings summary) or delegate to security-scan-skill.

Usage examples:

- `/skills run python-ci-skill --target @repo:/changed-files-list`
- `/skills run python-ci-skill --mode generate-patches --files @repo:src/foo.py`

Outputs (machine-readable):

- JSON with keys: { tool_results: [...], patches: ["--- a/..."], diagnostics: [...] }

---
skills/python-ci-skill/run.sh

## !/usr/bin/env bash

## Skill runner stub (non-executing). Intended for Claude Code to read/execute as a descriptor

echo "python-ci-skill runner invoked with: $*"

## Real implementation expected inside Skill: shell wrappers that call ruff/pytest, but harness keeps these as templates

---

skills/ha-addon-skill/SKILL.md

## ha-addon-skill

Role: Validate Home Assistant addon best practices and produce checklist + fixes.
Triggers: invoked by Hooks when files under addons/ change.
Tools (via MCP): custom YAML validators, schema checks, docs lint
Responsibilities:

- Check addon config.yaml, build.json, DOCS.md compliance with Home Assistant Add-on guidelines.
- Produce patch suggestions for config schema mismatches.

Usage:

- `/skills run ha-addon-skill --target @repo:addons/my-addon`

Output:

- JSON with { findings: [...], suggested_fixes_patch: "..." }

---
skills/security-scan-skill/SKILL.md

## security-scan-skill

Role: Centralize SAST/secret-detection orchestration for language scopes.
Triggers: invoked by Hooks or by other skills (python-ci-skill delegates here).
Tools (via MCP / Skill runtime): semgrep, bandit, trufflehog (or equivalent)
Responsibilities:

- Run language-appropriate SAST; prioritize Semgrep for rules, Bandit for Python AST checks.
- Return top findings only (token-budgeted): severity, file:line, quick fix hint.

Usage:

- `/skills run security-scan-skill --target @repo:/changed-files`

Output:

- JSON with { top_findings: [...], risk_level: "low|medium|high" }

---
manifest.md

## Claude Code Harness Manifest (how to use)

## Purpose

This harness is an in-Claude orchestration scaffold. Hooks detect edits/writes and generate/trigger skill invocations; actual linters/tests are invoked by the generated hook artifacts or by Skills via MCP.

## Typical flow (developer action -> harness)

1. Developer edits a Python file using Claude Code `Write`/`edit_file` tool.
2. `.claude/settings.toml` contains a hook mapping `on.edit_file:*.py` → `rule2hook_wrapper.sh` which generates a hook artifact.
3. Generated hook artifact either:
   - (a) Is materialized and executed by an operator (safe), or
   - (b) Is invoked by a Claude Skill (recommended): `/skills run python-ci-skill --target @repo:/changed-files`
4. `python-ci-skill` uses MCP references to fetch the changed files and returns JSON: patches + diagnostics.
5. If security sensitive, `python-ci-skill` delegates to `security-scan-skill`.

## How to call skills (examples)

- Run python CI skill on a branch:

---
/skills run python-ci-skill --target @repo:/branches/feature-x/changed-files

- Run HA addon checks:

/skills run ha-addon-skill --target @repo:/addons/my-addon

- Generate hooks from a natural language rule using rule2hook wrapper:

/run .claude/hooks/rule2hook_wrapper.sh --rule "on edit_file *.py -> run ruff --fix && run pytest" --label python-edit

## MCP usage (reference)

- Prefer MCP file fetches (e.g. `@repo:/path/to/file`) inside skill invocations to avoid pasting file contents.
- Skills must validate incoming MCP URIs and fail-fast if unauthorized.

## Safety notes

- Generated hook artifacts live under `.claude/generated-hooks/` — review before executing in CI/runner.

---
prompts/rule2hook-prompt.txt

## Prompt for claudecode-rule2hook (convert NL rule -> Claude Code hook)

You are claudecode-rule2hook. Input: a concise natural-language rule describing when and what to run.
Produce only: a Claude Code hook script file content (POSIX shell) and a short metadata header in JSON on the first line.
Constraints:

- Do NOT execute any linter commands; the script should be safe by default and only echo recommended commands.
- Metadata JSON keys: { "label": "<label>", "trigger": "<trigger-string>", "tools": ["ruff","pytest"] }

Example input:
  "on edit_file *.py -> run ruff --fix && run pytest"
Expected output (only file content):
  {"label":"python-edit","trigger":"on.edit_file:*.py","tools":["ruff","pytest"]}

## !/usr/bin/env bash

## Safe hook stub: echoes recommended commands and writes a recommended-run.sh under .claude/commands/

  echo "Recommended: ruff --fix <changed-files>"
  echo "Recommended: pytest -q <selected-tests>"
  exit 0
Do not output anything else.

---
prompts/text2hook-prompt.txt

## Prompt to convert a multi-line CLAUDE.md ruleset into multiple hook stubs

You are a transformer that reads a CLAUDE.md ruleset block and emits one hook stub per rule.
Input: CLAUDE.md block with rules (one per line).
Output: For each rule, emit:

  1) a JSON metadata line: {"label": "...", "trigger":"...", "tools":[...]}
  2) a POSIX shell stub that echoes the intended commands (safe, no execution).

Rules example:

  - "When a Dockerfile is edited run hadolint"
  - "When addons/* changes, trigger ha-addon-skill"

---

Kurze Tool-Übersicht mit Links, Zweck, Pro/Contra, empfohlene Rolle im Harness

Hinweis: direkte Quellen sind rechts als Links angegeben.

Ruff — schneller All-in-one Python Linter/Formatter

Link(s): <https://docs.astral.sh/ruff/>

Zweck: Linting + Formatting + viele Kompatibilitäts-Checker (ersetzt Flake8/isort/pyupgrade teilweise).

Pro: extrem schnell; --fix Auto-fix; breite Regelbasis.

Contra: Unterschiede zu Black/isort in Stil/outputs; ggf. Regel-tuning nötig.

Empfohlene Rolle in Harness: Hook (on.edit_file:*.py) oder Skill-invoked fix generator. Hooks should generate suggested patch via ruff --diff or ruff --fix template (do not auto-apply without review).

Hadolint — Dockerfile Linter

Link(s): <https://github.com/hadolint/hadolint>
Zweck: Lint Dockerfiles, parse AST, surface best-practice rules.

Pro: well-known, integrates ShellCheck for RUN lines.

Contra: ruleset tuning may be needed; some false positives for complex build patterns.

Empfohlene Rolle in Harness: Hook for Dockerfile edits (on.edit_file:Dockerfile) → generate hadolint report stub; skill may run a follow-up for remediation hints.

Semgrep — flexible SAST & custom rules

Link(s): <https://semgrep.dev/docs/>
Zweck: Rule-based SAST, supports custom rules/registries; good for appsec and config checks.

Pro: expressive rules, registry, triage UI (if using Semgrep App).

Contra: resource heavy for large repos; requires rule curation.

Empfohlene Rolle in Harness: security-scan-skill primary engine (delegate from python-ci-skill).

Bandit — Python security AST checks

Link(s): <https://bandit.readthedocs.io/>
Zweck: Find common security issues in Python code via AST plugins.

Pro: Python-specific focused checks; low overhead.

Contra: limited rule set; false positives possible.

Empfohlene Rolle in Harness: Invoked by security-scan-skill for quick Python security surface.

claudecode-rule2hook (community tool)

Link(s): <https://github.com/zxdxjtu/claudecode-rule2hook>
Zweck: Transform natural language rules into Claude Code hook scripts.

Pro: Rapid conversion from plain English rules to hook stubs; integrates with CLAUDE.md patterns.

Contra: Community project; generated hooks need review; dependency on Claude's NL interpretation.

Empfohlene Rolle in Harness: Primary generator for hook stubs (used by .claude/hooks/rule2hook_wrapper.sh), not the final execution engine. Always review generated artifacts.

Claude Code Hooks / Skills / MCP (official docs)

Link(s): Hooks guide & reference — <https://code.claude.com/docs/en/hooks>
Zweck: Native automation points inside Claude Code — deterministic shell commands (hooks), reusable encapsulations (skills), and file access (MCP).

Pro: Deterministic, pluggable, supports slash commands and MCP references.

Contra: Requires careful permissioning; generated scripts still need human review for side-effects.

Empfohlene Rolle in Harness: Orchestration layer — use hooks to detect events and generate/trigger Skills; use MCP for safe file access; Skills implement the actionable logic (but can return patches only, no forced application).

---

Empfohlene Mappings (was gehört wohin)

- Python file edits (*.py)
  -> Hook trigger: on.edit_file:*.py  (generate hook stub via rule2hook)
  -> Skill: python-ci-skill
  -> Tools: ruff (lint/format) [hook generates patch suggestions], pytest (test plan), mypy optional

- Dockerfile
  -> Hook trigger: on.edit_file:Dockerfile
  -> Skill: (light) docker-lint stub or direct hadolint suggestion
  -> Tools: hadolint (report)

- Home Assistant addons (addons/**)
  -> Hook trigger: on.edit_file:addons/**
  -> Skill: ha-addon-skill
  -> Tools: YAML schema checks, config validation, docs checks

- Security scans
  -> Trigger: post-change or manual `/skills run security-scan-skill`
  -> Skill: security-scan-skill
  -> Tools: semgrep (primary), bandit (python), trufflehog (secrets)

---
Spezifische, kurze Prompts (copy/paste) — für claudecode-rule2hook / text2hook

1) Prompt: Erzeuge eine Hook-Stub aus einer Regel (rule2hook)

You are claudecode-rule2hook. Convert this single-line natural language rule into a safe Claude Code hook stub.
Input rule:
on edit_file *.py -> run ruff --fix && run pytest
Output rules:

- First line: JSON metadata exactly: {"label":"python-edit","trigger":"on.edit_file:*.py","tools":["ruff","pytest"]}
- Then: a POSIX shell script that:
  - echoes the recommended commands (do NOT execute them)
  - writes a recommended-run.sh into .claude/safe-runs/python-edit.sh (content: the commands commented)

Return only the file content (metadata + script). No explanation.

2) Prompt: Multi-rule CLAUDE.md -> multiple hooks (text2hook)

You are text2hook. Input is a CLAUDE.md ruleset (multiple lines). For each rule produce:

1) JSON metadata line: {"label":"...","trigger":"...","tools":[...]}
2) A safe POSIX shell stub that only echoes and documents intended commands.

Example CLAUDE.md:
When a Dockerfile is edited run hadolint
When addons/* changes trigger ha-addon-skill
Output: emit the stub for each rule, concatenated. No extra text.

Minimal Beispiel: Wie ein Entwickler ein Hook generiert und Skill aufruft (copy/paste)

## generate hook artifact (local wrapper)

./.claude/hooks/rule2hook_wrapper.sh --rule "on edit_file *.py -> run ruff --fix && run pytest" --label python-edit

## (In Claude Code) invoke python-ci-skill using MCP reference to changed files

/skills run python-ci-skill --target @repo:/changes/feature-X

Abschließende Hinweise (nur Fakten / sehr kurz)

- Generated hook artifacts are stored under .claude/generated-hooks/ — review before execution.
- Use MCP URIs (@repo:...) inside Skill invocations to avoid large prompt payloads.
- Prefer hooks to *generate* commands and skill triggers; do not auto-execute destructive operations in hooks.

Quellen (direkte Links — kopierbar)

Claude Code — Hooks reference & guide: <https://code.claude.com/docs/en/hooks>
claudecode-rule2hook (GitHub): <https://github.com/zxdxjtu/claudecode-rule2hook>
Ruff docs: <https://docs.astral.sh/ruff/>
Ruff docs: <https://docs.astral.sh/ruff/>
Semgrep docs: <https://semgrep.dev/docs/>
Bandit docs / GitHub: <https://bandit.readthedocs.io/>
