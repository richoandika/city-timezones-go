package city

import (
	"container/list"
	"sync"
)

const (
	// DefaultMaxCacheSize is the default maximum number of cache entries
	DefaultMaxCacheSize = 1000
)

// cacheEntry represents a single cache entry with its key
type cacheEntry struct {
	key   string
	value []CityData
}

// SearchCache provides thread-safe caching for search results with LRU eviction
type SearchCache struct {
	mu        sync.RWMutex
	cache     map[string]*list.Element
	lruList   *list.List
	maxSize   int
	hits      uint64
	misses    uint64
	evictions uint64
}

// NewSearchCache creates a new search cache with default max size
func NewSearchCache() *SearchCache {
	return NewSearchCacheWithSize(DefaultMaxCacheSize)
}

// NewSearchCacheWithSize creates a new search cache with specified max size
func NewSearchCacheWithSize(maxSize int) *SearchCache {
	if maxSize <= 0 {
		maxSize = DefaultMaxCacheSize
	}
	return &SearchCache{
		cache:   make(map[string]*list.Element),
		lruList: list.New(),
		maxSize: maxSize,
	}
}

// Get retrieves a cached result and updates LRU order
func (c *SearchCache) Get(key string) ([]CityData, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	element, exists := c.cache[key]
	if !exists {
		c.misses++
		return nil, false
	}

	// Move to front (most recently used)
	c.lruList.MoveToFront(element)
	c.hits++

	entry := element.Value.(*cacheEntry)
	return entry.value, true
}

// Set stores a result in the cache with LRU eviction
func (c *SearchCache) Set(key string, result []CityData) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Check if key already exists
	if element, exists := c.cache[key]; exists {
		// Update existing entry and move to front
		c.lruList.MoveToFront(element)
		entry := element.Value.(*cacheEntry)
		entry.value = result
		return
	}

	// Add new entry
	entry := &cacheEntry{
		key:   key,
		value: result,
	}
	element := c.lruList.PushFront(entry)
	c.cache[key] = element

	// Evict least recently used if over capacity
	if c.lruList.Len() > c.maxSize {
		c.evictOldest()
	}
}

// evictOldest removes the least recently used entry (must be called with lock held)
func (c *SearchCache) evictOldest() {
	oldest := c.lruList.Back()
	if oldest != nil {
		c.lruList.Remove(oldest)
		entry := oldest.Value.(*cacheEntry)
		delete(c.cache, entry.key)
		c.evictions++
	}
}

// Clear clears the cache
func (c *SearchCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache = make(map[string]*list.Element)
	c.lruList = list.New()
	// Note: We don't reset statistics on clear
}

// Size returns the number of cached entries
func (c *SearchCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.cache)
}

// MaxSize returns the maximum cache size
func (c *SearchCache) MaxSize() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.maxSize
}

// Stats returns cache statistics
func (c *SearchCache) Stats() CacheStats {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var hitRate float64
	total := c.hits + c.misses
	if total > 0 {
		hitRate = float64(c.hits) / float64(total) * 100
	}

	return CacheStats{
		Size:      len(c.cache),
		MaxSize:   c.maxSize,
		Hits:      c.hits,
		Misses:    c.misses,
		Evictions: c.evictions,
		HitRate:   hitRate,
	}
}

// CacheStats contains cache performance statistics
type CacheStats struct {
	Size      int     // Current number of entries
	MaxSize   int     // Maximum number of entries
	Hits      uint64  // Number of cache hits
	Misses    uint64  // Number of cache misses
	Evictions uint64  // Number of evictions due to size limit
	HitRate   float64 // Cache hit rate as percentage
}

// Global cache instance
var searchCache = NewSearchCache()

// GetCachedResult retrieves a cached search result
func GetCachedResult(key string) ([]CityData, bool) {
	return searchCache.Get(key)
}

// SetCachedResult stores a search result in cache
func SetCachedResult(key string, result []CityData) {
	searchCache.Set(key, result)
}

// ClearCache clears the global search cache
func ClearCache() {
	searchCache.Clear()
}

// CacheSize returns the size of the global cache
func CacheSize() int {
	return searchCache.Size()
}

// CacheMaxSize returns the maximum size of the global cache
func CacheMaxSize() int {
	return searchCache.MaxSize()
}

// CacheStatistics returns statistics about the global cache
func CacheStatistics() CacheStats {
	return searchCache.Stats()
}
