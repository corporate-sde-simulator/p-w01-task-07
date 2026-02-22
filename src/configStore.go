package config

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

// ConfigStore holds configuration values with thread-safe access.
//
// Author: Suresh Kumar
// Last Modified: 2026-02-07

type ConfigStore struct {
	data      map[string]interface{}
	defaults  map[string]interface{}
	listeners map[string][]func(oldVal, newVal interface{})
	mu        sync.RWMutex
}

func NewConfigStore() *ConfigStore {
	return &ConfigStore{
		data:      make(map[string]interface{}),
		defaults:  make(map[string]interface{}),
		listeners: make(map[string][]func(oldVal, newVal interface{})),
	}
}

// Get retrieves a config value using dot notation (e.g., "database.pool.max_size").
func (cs *ConfigStore) Get(key string) (interface{}, bool) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	parts := strings.Split(key, ".")
	// BUG: Only looks at top-level key, ignores nested traversal
	// Example: Get("database.host") returns the entire "database" map,
	// not the "host" value inside it
	val, ok := cs.data[parts[0]]
	if ok {
		return val, true
	}

	// Fallback to defaults
	val, ok = cs.defaults[key]
	return val, ok
}

// Set updates a config value and notifies listeners.
func (cs *ConfigStore) Set(key string, value interface{}) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	oldVal := cs.data[key]
	cs.data[key] = value

	// Notify listeners watching this key
	if listeners, exists := cs.listeners[key]; exists {
		for _, listener := range listeners {
			listener(oldVal, value)
		}
	}
}

// SetDefaults sets default values for configuration keys (used as fallback).
func (cs *ConfigStore) SetDefaults(defaults map[string]interface{}) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.defaults = defaults
}

// MergeConfig merges new config data into existing data.
func (cs *ConfigStore) MergeConfig(newData map[string]interface{}) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	for key, newVal := range newData {
		// BUG: Replaces entire subtree instead of deep merging
		// If base has {db: {host: "localhost", port: 5432}} and
		// newData has {db: {port: 5433}}, this LOSES the host key
		cs.data[key] = newVal
	}
}

// Watch registers a listener for changes to a specific key.
func (cs *ConfigStore) Watch(key string, callback func(oldVal, newVal interface{})) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	cs.listeners[key] = append(cs.listeners[key], callback)
}

// GetAll returns a copy of all configuration data.
func (cs *ConfigStore) GetAll() map[string]interface{} {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	result := make(map[string]interface{})
	for k, v := range cs.data {
		result[k] = v
	}
	return result
}

// ApplyEnvOverrides applies environment variable overrides.
// Format: CONFIG_DATABASE_HOST overrides "database.host"
func (cs *ConfigStore) ApplyEnvOverrides(prefix string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) != 2 {
			continue
		}
		envKey := parts[0]
		envVal := parts[1]

		if strings.HasPrefix(envKey, prefix+"_") {
			configKey := strings.ToLower(strings.Replace(
				strings.TrimPrefix(envKey, prefix+"_"), "_", ".", -1,
			))
			cs.data[configKey] = envVal
		}
	}
}

// String returns a debug representation.
func (cs *ConfigStore) String() string {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return fmt.Sprintf("ConfigStore{keys: %d}", len(cs.data))
}
