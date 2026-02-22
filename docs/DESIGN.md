# ADR-015: Config Reload Strategy â€” Polling vs File System Events

**Date:**  
**Status:** Accepted  
**Authors:** Vikram Patel, Suresh Kumar

## Decision

Use **file polling** with a 2-second interval for config change detection rather than OS-level filesystem events (inotify/kqueue).

## Context

Services need to detect config file changes at runtime and reload without restart. The config files are on mounted volumes (Kubernetes ConfigMaps) which don't reliably trigger inotify events.

## Options Considered

| Option | Reliability | Cross-platform | Latency | Complexity |
|---|---|---|---|---|
| Polling (stat-based) | High | Yes | 1-2s | Low |
| inotify/kqueue | Medium (volume mounts) | No | <100ms | Medium |
| Config server (Consul/etcd) | High | Yes | <100ms | High |
| Signal-based (SIGHUP) | High | Unix only | Instant | Low |

## Rationale

- Polling is the most portable and reliable approach for Kubernetes ConfigMap volumes
- 2-second delay is acceptable for our use case (config changes are rare)
- No external dependencies required
- Simple to implement and debug

## Consequences

- Slight delay (up to 2 seconds) before changes take effect
- CPU cost of stat calls every 2 seconds (negligible)
- Must track last modification time correctly to avoid false positives
