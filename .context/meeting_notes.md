# Meeting Notes — Sprint 23 Standup

**Date:** Feb 12, 2026  
**Attendees:** Vikram (Platform Lead), Suresh, Neha, Intern

---

## Config Hot-Reload

- **Vikram:** We had three outages last month because of config changes requiring restarts. We need hot-reload ASAP. @Intern, take over PLATFORM-2844 from Suresh.

- **Suresh:** The file watcher is there but the modification time comparison might be wrong — I think I'm comparing with the wrong baseline. Also the config merge for nested keys is broken — it replaces entire subtrees instead of merging.

- **Neha:** For the listener pattern, make sure listeners are called with both old and new values so they can react appropriately.

## Action Items

- [ ] @Intern — Fix config hot-reload service (PLATFORM-2844)
- [ ] @Suresh — On-call but available for questions
- [ ] @Neha — Integrate with feature flag service after merge
