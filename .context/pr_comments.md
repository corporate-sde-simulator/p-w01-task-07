# PR #378 Review — Config Hot-Reload Service (by Suresh Kumar)

## Reviewer: Vikram Patel (Platform Lead) — Feb 7, 2026

---

**Overall:** Architecture is solid but the implementation has bugs.

### `configStore.go`

> **Line 45** — `Get` method with dot notation:  
> Your key splitting logic splits on `.` but then only returns the top-level key. You need to recursively traverse the nested map.

> **Line 78** — `mergeConfig` method:  
> When merging nested maps, you're replacing the entire subtree instead of merging key-by-key. If the base config has `{database: {host: "localhost", port: 5432}}` and the override has `{database: {port: 5433}}`, you lose the `host` key.

### `fileWatcher.go`

> **Line 32** — `checkForChanges`:  
> You're storing `lastModTime` but never updating it after a reload. So every poll cycle detects a "change" even when nothing changed, causing infinite reload loops.

---

**Suresh Kumar** — Feb 8, 2026

> The merge bug is the worst one — it can silently drop config keys. I'll note these for the intern.
