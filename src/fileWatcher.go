package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// FileWatcher polls a config file for changes and triggers reload.
//
// Author: Suresh Kumar
// Last Modified: 2026-02-07

type FileWatcher struct {
	filePath     string
	store        *ConfigStore
	pollInterval time.Duration
	lastModTime  time.Time
	stopChan     chan struct{}
	running      bool
}

func NewFileWatcher(filePath string, store *ConfigStore, interval time.Duration) *FileWatcher {
	return &FileWatcher{
		filePath:     filePath,
		store:        store,
		pollInterval: interval,
		stopChan:     make(chan struct{}),
	}
}

// Start begins polling the config file for changes.
func (fw *FileWatcher) Start() error {
	if fw.running {
		return fmt.Errorf("watcher already running")
	}

	// Load initial config
	if err := fw.loadConfig(); err != nil {
		return fmt.Errorf("initial config load failed: %w", err)
	}

	fw.running = true
	go fw.pollLoop()
	return nil
}

// Stop halts the file watcher.
func (fw *FileWatcher) Stop() {
	if fw.running {
		close(fw.stopChan)
		fw.running = false
	}
}

func (fw *FileWatcher) pollLoop() {
	ticker := time.NewTicker(fw.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-fw.stopChan:
			return
		case <-ticker.C:
			if fw.checkForChanges() {
				fw.loadConfig()
				// BUG: Never updates lastModTime after reload
				// This means every subsequent poll detects a "change"
				// and triggers an unnecessary reload cycle
			}
		}
	}
}

func (fw *FileWatcher) checkForChanges() bool {
	info, err := os.Stat(fw.filePath)
	if err != nil {
		return false
	}
	return info.ModTime().After(fw.lastModTime)
}

func (fw *FileWatcher) loadConfig() error {
	data, err := os.ReadFile(fw.filePath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	var configData map[string]interface{}
	if err := json.Unmarshal(data, &configData); err != nil {
		return fmt.Errorf("failed to parse config JSON: %w", err)
	}

	fw.store.MergeConfig(configData)
	return nil
}

// IsRunning returns whether the watcher is currently active.
func (fw *FileWatcher) IsRunning() bool {
	return fw.running
}
