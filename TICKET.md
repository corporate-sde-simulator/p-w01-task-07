# PLATFORM-2844: Implement configuration hot-reload service

**Status:** In Progress Â· **Priority:** Medium
**Sprint:** Sprint 23 Â· **Story Points:** 5
**Reporter:** Vikram Patel (Platform Lead) Â· **Assignee:** You (Intern)
**Created:** Â· **Due:** End of sprint (Friday)
**Labels:** `backend`, `config`, `golang`, `infrastructure`
**Epic:** PLATFORM-2805 (Platform Configuration v2)
**Task Type:** ðŸ› Bug Fix

---

## Description

Our services require a restart every time we change config values (feature flags, timeouts, thresholds). We need a configuration manager that watches config files and reloads values at runtime without downtime.

Suresh (senior dev) started the file watcher and config store but got pulled into a production incident. His code has bugs in the config store's nested key lookup, the merge logic, and the file watcher's change detection.

## Requirements

- Watch a JSON config file for changes
- Reload configuration automatically when file changes
- Support nested config keys with dot notation (`database.pool.max_size`)
- Thread-safe reads (config can be read while being updated)
- Notify registered listeners when config values change
- Support default values and environment variable overrides

## Acceptance Criteria

- [ ] Bug #1 fixed: `Get()` only reads top-level key, ignores nested dot-notation traversal
- [ ] Bug #2 fixed: `MergeConfig()` replaces entire subtree instead of deep merging
- [ ] Bug #3 fixed: `pollLoop()` never updates `lastModTime` after reload, causing infinite re-detection
- [ ] All unit tests pass
- [ ] Config changes detected within 2 seconds

## Design Notes

See `docs/DESIGN.md` for the watcher architecture.
See `.context/pr_comments.md` for Suresh's PR feedback.
