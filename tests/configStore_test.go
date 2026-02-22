package config

import (
	"testing"
)

func TestConfigStore_SetAndGet(t *testing.T) {
	store := NewConfigStore()
	store.Set("database.host", "localhost")

	val, ok := store.Get("database.host")
	if !ok {
		t.Fatal("expected key to exist")
	}
	if val != "localhost" {
		t.Errorf("expected 'localhost', got '%v'", val)
	}
}

func TestConfigStore_GetNestedKey(t *testing.T) {
	store := NewConfigStore()
	store.MergeConfig(map[string]interface{}{
		"database": map[string]interface{}{
			"pool": map[string]interface{}{
				"max_size": 20,
			},
		},
	})

	val, ok := store.Get("database.pool.max_size")
	if !ok {
		t.Fatal("expected nested key to exist")
	}
	if val != 20 {
		t.Errorf("expected 20, got %v", val)
	}
}

func TestConfigStore_DefaultValues(t *testing.T) {
	store := NewConfigStore()
	store.SetDefaults(map[string]interface{}{
		"timeout": 30,
	})

	val, ok := store.Get("timeout")
	if !ok {
		t.Fatal("expected default to exist")
	}
	if val != 30 {
		t.Errorf("expected 30, got %v", val)
	}
}

func TestConfigStore_MergePreservesExistingKeys(t *testing.T) {
	store := NewConfigStore()
	store.MergeConfig(map[string]interface{}{
		"database": map[string]interface{}{
			"host": "localhost",
			"port": 5432,
		},
	})

	// Merge override for port only
	store.MergeConfig(map[string]interface{}{
		"database": map[string]interface{}{
			"port": 5433,
		},
	})

	// Should still have host
	all := store.GetAll()
	db, ok := all["database"].(map[string]interface{})
	if !ok {
		t.Fatal("database key should be a map")
	}
	if db["host"] != "localhost" {
		t.Errorf("merge should preserve existing keys, host is missing")
	}
	if db["port"] != 5433 {
		t.Errorf("merge should update changed keys, expected 5433 got %v", db["port"])
	}
}

func TestConfigStore_WatchNotifiesListener(t *testing.T) {
	store := NewConfigStore()
	notified := false

	store.Watch("api.timeout", func(oldVal, newVal interface{}) {
		notified = true
	})

	store.Set("api.timeout", 60)
	if !notified {
		t.Error("expected listener to be notified")
	}
}

func TestConfigStore_GetMissingKey(t *testing.T) {
	store := NewConfigStore()
	_, ok := store.Get("nonexistent.key")
	if ok {
		t.Error("expected missing key to return false")
	}
}

func TestConfigStore_ConcurrentAccess(t *testing.T) {
	store := NewConfigStore()
	done := make(chan bool, 2)

	go func() {
		for i := 0; i < 1000; i++ {
			store.Set("counter", i)
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			store.Get("counter")
		}
		done <- true
	}()

	<-done
	<-done
}
