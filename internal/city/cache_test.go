package city

import (
	"testing"
)

func TestSearchCache(t *testing.T) {
	t.Run("NewSearchCache", func(t *testing.T) {
		cache := NewSearchCache()
		if cache == nil {
			t.Error("cache should not be nil")
		}
		if cache.Size() != 0 {
			t.Errorf("initial cache size should be 0, got %d", cache.Size())
		}
	})

	t.Run("Set and Get", func(t *testing.T) {
		cache := NewSearchCache()
		testData := []CityData{
			{City: "Chicago", ISO2: "US", Timezone: "America/Chicago"},
		}

		cache.Set("chicago", testData)
		if cache.Size() != 1 {
			t.Errorf("cache size should be 1, got %d", cache.Size())
		}

		result, exists := cache.Get("chicago")
		if !exists {
			t.Error("key should exist")
		}
		if len(result) != 1 {
			t.Errorf("result length should be 1, got %d", len(result))
		}
		if result[0].City != "Chicago" {
			t.Errorf("expected Chicago, got %s", result[0].City)
		}
	})

	t.Run("Clear", func(t *testing.T) {
		cache := NewSearchCache()
		testData := []CityData{{City: "Chicago"}}

		cache.Set("chicago", testData)
		cache.Clear()

		if cache.Size() != 0 {
			t.Errorf("cache should be empty, got %d", cache.Size())
		}
	})
}

func TestGlobalCache(t *testing.T) {
	t.Run("Global cache operations", func(t *testing.T) {
		// Clear cache first
		ClearCache()
		if CacheSize() != 0 {
			t.Errorf("cache should be empty, got %d", CacheSize())
		}

		// Test data
		testData := []CityData{
			{City: "Chicago", ISO2: "US", Timezone: "America/Chicago"},
		}

		// Set cached result
		SetCachedResult("chicago", testData)
		if CacheSize() != 1 {
			t.Errorf("cache should have 1 item, got %d", CacheSize())
		}

		// Get cached result
		result, exists := GetCachedResult("chicago")
		if !exists {
			t.Error("key should exist")
		}
		if len(result) != 1 {
			t.Errorf("result length should be 1, got %d", len(result))
		}
		if result[0].City != "Chicago" {
			t.Errorf("expected Chicago, got %s", result[0].City)
		}
	})
}

func TestLRUEviction(t *testing.T) {
	t.Run("Eviction when cache is full", func(t *testing.T) {
		// Create a small cache for testing
		cache := NewSearchCacheWithSize(3)

		testData := []CityData{{City: "Test"}}

		// Fill cache to max capacity
		cache.Set("key1", testData)
		cache.Set("key2", testData)
		cache.Set("key3", testData)

		if cache.Size() != 3 {
			t.Errorf("cache size should be 3, got %d", cache.Size())
		}

		// Add one more, should evict oldest (key1)
		cache.Set("key4", testData)

		if cache.Size() != 3 {
			t.Errorf("cache size should still be 3, got %d", cache.Size())
		}

		// key1 should be evicted
		_, exists := cache.Get("key1")
		if exists {
			t.Error("key1 should have been evicted")
		}

		// Others should still exist
		_, exists = cache.Get("key2")
		if !exists {
			t.Error("key2 should exist")
		}
	})

	t.Run("LRU order maintained", func(t *testing.T) {
		cache := NewSearchCacheWithSize(3)
		testData := []CityData{{City: "Test"}}

		cache.Set("key1", testData)
		cache.Set("key2", testData)
		cache.Set("key3", testData)

		// Access key1 to make it recently used
		cache.Get("key1")

		// Add key4, should evict key2 (least recently used)
		cache.Set("key4", testData)

		_, exists := cache.Get("key2")
		if exists {
			t.Error("key2 should have been evicted")
		}

		_, exists = cache.Get("key1")
		if !exists {
			t.Error("key1 should still exist (was recently used)")
		}
	})

	t.Run("Update existing entry", func(t *testing.T) {
		cache := NewSearchCacheWithSize(3)
		testData1 := []CityData{{City: "Test1"}}
		testData2 := []CityData{{City: "Test2"}}

		cache.Set("key1", testData1)
		initialSize := cache.Size()

		// Update existing key
		cache.Set("key1", testData2)

		if cache.Size() != initialSize {
			t.Error("cache size should not change when updating existing key")
		}

		result, exists := cache.Get("key1")
		if !exists {
			t.Error("key1 should exist")
		}
		if result[0].City != "Test2" {
			t.Errorf("expected Test2, got %s", result[0].City)
		}
	})
}

func TestCacheStats(t *testing.T) {
	t.Run("Statistics tracking", func(t *testing.T) {
		cache := NewSearchCacheWithSize(5)
		testData := []CityData{{City: "Test"}}

		// Initial stats
		stats := cache.Stats()
		if stats.Size != 0 {
			t.Errorf("initial size should be 0, got %d", stats.Size)
		}
		if stats.MaxSize != 5 {
			t.Errorf("max size should be 5, got %d", stats.MaxSize)
		}

		// Add entries
		cache.Set("key1", testData)
		cache.Set("key2", testData)

		// Hit
		cache.Get("key1")
		// Miss
		cache.Get("key3")

		stats = cache.Stats()
		if stats.Size != 2 {
			t.Errorf("size should be 2, got %d", stats.Size)
		}
		if stats.Hits != 1 {
			t.Errorf("hits should be 1, got %d", stats.Hits)
		}
		if stats.Misses != 1 {
			t.Errorf("misses should be 1, got %d", stats.Misses)
		}
		if stats.HitRate != 50.0 {
			t.Errorf("hit rate should be 50%%, got %.2f%%", stats.HitRate)
		}
	})

	t.Run("Eviction tracking", func(t *testing.T) {
		cache := NewSearchCacheWithSize(2)
		testData := []CityData{{City: "Test"}}

		cache.Set("key1", testData)
		cache.Set("key2", testData)
		cache.Set("key3", testData) // Triggers eviction

		stats := cache.Stats()
		if stats.Evictions != 1 {
			t.Errorf("evictions should be 1, got %d", stats.Evictions)
		}
	})

	t.Run("Global cache statistics", func(t *testing.T) {
		ClearCache()
		testData := []CityData{{City: "Test"}}

		SetCachedResult("test", testData)
		GetCachedResult("test")
		GetCachedResult("nonexistent")

		stats := CacheStatistics()
		if stats.Size != 1 {
			t.Errorf("size should be 1, got %d", stats.Size)
		}
		if stats.Hits < 1 {
			t.Errorf("hits should be at least 1, got %d", stats.Hits)
		}
	})
}

func TestCacheMaxSize(t *testing.T) {
	t.Run("Default max size", func(t *testing.T) {
		cache := NewSearchCache()
		if cache.MaxSize() != DefaultMaxCacheSize {
			t.Errorf("default max size should be %d, got %d", DefaultMaxCacheSize, cache.MaxSize())
		}
	})

	t.Run("Custom max size", func(t *testing.T) {
		cache := NewSearchCacheWithSize(100)
		if cache.MaxSize() != 100 {
			t.Errorf("max size should be 100, got %d", cache.MaxSize())
		}
	})

	t.Run("Invalid max size defaults to default", func(t *testing.T) {
		cache := NewSearchCacheWithSize(0)
		if cache.MaxSize() != DefaultMaxCacheSize {
			t.Errorf("should default to %d, got %d", DefaultMaxCacheSize, cache.MaxSize())
		}

		cache = NewSearchCacheWithSize(-10)
		if cache.MaxSize() != DefaultMaxCacheSize {
			t.Errorf("should default to %d, got %d", DefaultMaxCacheSize, cache.MaxSize())
		}
	})

	t.Run("Global cache max size", func(t *testing.T) {
		maxSize := CacheMaxSize()
		if maxSize != DefaultMaxCacheSize {
			t.Errorf("global cache max size should be %d, got %d", DefaultMaxCacheSize, maxSize)
		}
	})
}
