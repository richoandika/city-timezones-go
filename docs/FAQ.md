# ❓ Frequently Asked Questions (FAQ)

This document answers common questions about using City Timezones Go.

## Table of Contents

- [General Questions](#general-questions)
- [Performance Questions](#performance-questions)
- [Usage Questions](#usage-questions)
- [Troubleshooting](#troubleshooting)
- [Data Questions](#data-questions)

## General Questions

### Q: How often is the timezone data updated?

**A:** The timezone data is based on the original city-timezones dataset and includes over 7,000 cities worldwide. For updates, check the project repository for new releases. If you need to update the data, please open an issue or pull request.

### Q: What happens if a city isn't found?

**A:** If a city isn't found, the search functions return an empty slice (`[]CityData`) without an error. This allows you to distinguish between lookup failures (errors) and legitimate "no results found" cases.

```go
cities, err := citytimezones.LookupViaCity("NonExistentCity")
if err != nil {
    // Handle error (validation failure, etc.)
}
if len(cities) == 0 {
    // City not found - this is normal, not an error
}
```

### Q: Is this library safe for concurrent use?

**A:** Yes! The library is fully thread-safe. All search operations and cache access are protected with proper synchronization primitives (`sync.Once`, `sync.RWMutex`), making it safe to use from multiple goroutines simultaneously.

### Q: How much memory does the library use?

**A:** The library uses approximately:
- **~2MB** for the city data (loaded lazily on first use)
- **Up to ~1000 entries** in the search cache by default (configurable)
- Cache uses LRU eviction to prevent unbounded growth

Total memory footprint is typically under 5MB in production use.

### Q: Does the library require external dependencies?

**A:** No! The library has zero external dependencies, making it lightweight and reducing security concerns. Only the Go standard library is used.

### Q: Which Go versions are supported?

**A:** The library requires Go 1.21 or higher. It's tested on Go 1.21, 1.22, and 1.23 in CI.

## Performance Questions

### Q: How fast are the lookups?

**A:** Performance varies by operation:
- **Exact city name lookup**: ~0.5ms (first call), ~0.05ms (cached)
- **Partial match search**: ~1-2ms
- **ISO code lookup**: ~3-5ms

Results are cached automatically for faster subsequent lookups. With caching, most lookups complete in under 100 microseconds.

### Q: Will the library slow down my application?

**A:** No. The library uses lazy loading (data loads only on first use) and aggressive caching. After the initial load (~10ms), lookups are extremely fast.

### Q: Can I control the cache size?

**A:** The cache has a default maximum size of 1000 entries with LRU eviction. For custom cache management, you can:

```go
// Check current cache size
size := citytimezones.CacheSize()

// Check maximum cache size
maxSize := citytimezones.CacheMaxSize()

// Get detailed statistics
stats := citytimezones.GetCacheStats()
fmt.Printf("Hit rate: %.2f%%\n", stats.HitRate)

// Clear cache if needed
citytimezones.ClearCache()
```

### Q: How do I benchmark my usage?

**A:** Use Go's built-in benchmarking:

```go
func BenchmarkCityLookup(b *testing.B) {
    for i := 0; i < b.N; i++ {
        citytimezones.LookupViaCity("Chicago")
    }
}
```

## Usage Questions

### Q: How do I search for a city with a partial name?

**A:** Use `FindFromCityStateProvince()` which supports partial matching:

```go
cities, err := citytimezones.FindFromCityStateProvince("york")
// Finds "New York", "York", etc.
```

### Q: Can I search for cities in a specific country?

**A:** Yes, use `FindFromIsoCode()` with the country's ISO2 or ISO3 code:

```go
cities, err := citytimezones.FindFromIsoCode("US")  // or "USA"
```

### Q: How do I handle cities with the same name?

**A:** Search functions return a slice of cities. Check the `Province` and `Country` fields to distinguish between them:

```go
cities, _ := citytimezones.LookupViaCity("Springfield")
for _, city := range cities {
    fmt.Printf("%s, %s, %s\n", city.City, city.Province, city.Country)
}
```

### Q: Can I use custom search options?

**A:** Yes, use `SearchCities()` with `SearchOptions`:

```go
options := citytimezones.SearchOptions{
    CaseSensitive: true,
    ExactMatch:    false,
}
cities, err := citytimezones.SearchCities("Chicago", options)
```

### Q: How do I integrate this with a web API?

**A:** Here's a simple example:

```go
http.HandleFunc("/api/city", func(w http.ResponseWriter, r *http.Request) {
    cityName := r.URL.Query().Get("name")
    cities, err := citytimezones.LookupViaCity(cityName)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    json.NewEncoder(w).Encode(cities)
})
```

## Troubleshooting

### Q: Why am I getting "validation error" messages?

**A:** The library validates all inputs for security. Common causes:
- Input strings that are too long (>100 chars for city names, >200 for search strings)
- Invalid characters or suspected injection attempts
- Empty search strings

**Solution:** Ensure your inputs are reasonable lengths and don't contain suspicious patterns.

### Q: The library seems slow on the first call. Why?

**A:** The city data is loaded lazily on first use (lazy loading pattern). The ~2MB dataset loads in about 10ms. Subsequent calls will be much faster (microseconds with caching).

**Solution:** If you need immediate first-call performance, you can "warm up" the cache:

```go
// Warm up the library at startup
_, _ = citytimezones.GetCityMapping()
```

### Q: How can I clear the cache in a long-running application?

**A:** Call `citytimezones.ClearCache()` periodically if needed:

```go
// Example: Clear cache every hour
ticker := time.NewTicker(1 * time.Hour)
go func() {
    for range ticker.C {
        citytimezones.ClearCache()
    }
}()
```

### Q: Can I use this library in a web server?

**A:** Absolutely! The library is designed for concurrent use and performs well in web applications. The thread-safe cache and fast lookups make it ideal for API servers.

### Q: I'm getting race condition warnings. What should I do?

**A:** The library itself is fully thread-safe. If you see race condition warnings:
1. Ensure you're not modifying returned `CityData` structs from multiple goroutines
2. Run your tests with `-race` flag to identify the issue in your code
3. The library's own tests pass with `-race` flag

## Data Questions

### Q: Where does the timezone data come from?

**A:** The data is sourced from the original [city-timezones](https://github.com/kevinroberts/city-timezones) JavaScript library, which aggregates data from multiple sources including GeoNames.

### Q: Are the coordinates accurate?

**A:** The coordinates represent approximate city center locations. For precise geolocation, consider using a dedicated geolocation service.

### Q: What if I need to add custom city data?

**A:** Currently, the library uses a static dataset. For custom data, you can:
1. Fork the repository and modify `data/cityMap.json`
2. Submit a pull request to add missing cities
3. Open an issue for data corrections

We're considering adding custom data support in v2.0 (see [ROADMAP.md](ROADMAP.md)).

### Q: Does the library handle daylight saving time?

**A:** The library provides timezone identifiers (e.g., "America/New_York"). Use Go's `time` package to handle DST:

```go
loc, _ := time.LoadLocation(city.Timezone)
now := time.Now().In(loc)
```

The timezone identifier automatically handles DST transitions through Go's time package.

### Q: How accurate is the population data?

**A:** Population data is approximate and may not be current. It's provided for reference but shouldn't be relied upon for critical applications requiring precise population figures.

### Q: Can I get timezone offset information?

**A:** Yes, using Go's time package:

```go
cities, _ := citytimezones.LookupViaCity("Tokyo")
if len(cities) > 0 {
    loc, _ := time.LoadLocation(cities[0].Timezone)
    _, offset := time.Now().In(loc).Zone()
    fmt.Printf("Offset: %d seconds\n", offset)
}
```

## Still Have Questions?

If your question isn't answered here:

1. Check the [API Documentation](API.md)
2. Look through [GitHub Issues](https://github.com/richoandika/city-timezones-go/issues)
3. Ask in [GitHub Discussions](https://github.com/richoandika/city-timezones-go/discussions)
4. Open a new issue

---

[⬆ Back to README](../README.md)
