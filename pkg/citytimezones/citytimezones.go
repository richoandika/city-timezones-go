package citytimezones

import (
	"github.com/richoandika/city-timezones-go/internal/city"
)

// CityData represents a city with its timezone and geographical information
type CityData = city.CityData

// SearchOptions provides configuration for search operations
type SearchOptions = city.SearchOptions

// LookupViaCity searches for cities by exact city name match
func LookupViaCity(cityName string) ([]CityData, error) {
	return city.LookupViaCity(cityName)
}

// FindFromCityStateProvince searches for cities using partial matching
// across city, state, province, and country fields
func FindFromCityStateProvince(searchString string) ([]CityData, error) {
	return city.FindFromCityStateProvince(searchString)
}

// FindFromIsoCode searches for cities by ISO2 or ISO3 country codes
func FindFromIsoCode(isoCode string) ([]CityData, error) {
	return city.FindFromIsoCode(isoCode)
}

// SearchCities provides a flexible search function with options
func SearchCities(query string, options SearchOptions) ([]CityData, error) {
	return city.SearchCities(query, options)
}

// GetCityMapping returns all available cities
func GetCityMapping() ([]CityData, error) {
	return city.GetCityData()
}

// DefaultSearchOptions returns the default search configuration
func DefaultSearchOptions() SearchOptions {
	return city.DefaultSearchOptions()
}

// CacheStats contains cache performance statistics
type CacheStats = city.CacheStats

// ClearCache clears the global search cache
func ClearCache() {
	city.ClearCache()
}

// CacheSize returns the current number of entries in the cache
func CacheSize() int {
	return city.CacheSize()
}

// CacheMaxSize returns the maximum number of entries the cache can hold
func CacheMaxSize() int {
	return city.CacheMaxSize()
}

// GetCacheStats returns performance statistics about the cache
func GetCacheStats() CacheStats {
	return city.CacheStatistics()
}
