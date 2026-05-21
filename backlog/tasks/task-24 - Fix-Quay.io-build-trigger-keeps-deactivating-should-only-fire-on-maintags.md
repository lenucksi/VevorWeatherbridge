---
id: TASK-24
title: 'Fix Quay.io build trigger (keeps deactivating, should only fire on main+tags)'
status: To Do
assignee: []
created_date: '2026-05-21 01:28'
updated_date: '2026-05-21 01:28'
labels:
  - quay
  - ci
dependencies: []
priority: high
ordinal: 24000
---

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Research why Quay deactivates the build trigger
- [ ] #2 Create OAuth token with repo:admin scope for API access
- [ ] #3 Configure trigger to only build main→latest and v*→version tags
- [ ] #4 Verify trigger stays active after configuration
<!-- AC:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
## Research Findings

### Why trigger deactivates
- GitHub Webhook returns 4xx/5xx (repo temporarily unreachable)
- SSH deploy key was rotated/removed on GitHub side
- Repeated build failures trigger auto-disable in Quay
- Quay doesn't expose the exact reason via API to robot accounts

### Current state
- Trigger fires on EVERY push (default behavior)
- Robot token (QUAY_USERNAME/QUAY_TOKEN) only has Docker push/pull scope
- QUAY_API_TOKEN doesn't have repo:admin scope for management API
- No way to inspect/modify trigger config without proper OAuth token

### What needs to happen
1. Create OAuth token via Quay UI: Account → Applications → New App → Generate Token with 'repo:admin' + 'org:admin' scopes
2. Store as QUAY_OAUTH_TOKEN GitHub secret
3. Via API: GET /api/v1/repository/lenucksi-gh/vevor-weatherbridge-go-amd64/ → inspect trigger config
4. Via API: PUT trigger to set: branch=main, tag=v.* (only builds on main pushes and version tags)
5. Alternative: Configure via Quay web UI → Repository Settings → Build Trigger → edit trigger config

### API endpoints (need OAuth token with repo:admin)
- GET /api/v1/repository/{namespace}/{repo}/ — repo info including triggers
- PUT /api/v1/repository/{namespace}/{repo}/trigger/{id} — update trigger
- Robot accounts (Basic auth) work only for Docker registry v2 operations, not management API
- Full API docs: https://docs.quay.io/api/
<!-- SECTION:NOTES:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 Root cause identified and documented
- [ ] #2 Trigger stays active for >7 days
- [ ] #3 Only fires on main branch and version tags (v*)
<!-- DOD:END -->
