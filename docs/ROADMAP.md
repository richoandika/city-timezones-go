# 🗺️ Roadmap

This document outlines the development roadmap for City Timezones Go, including completed features and planned enhancements.

## ✅ Implemented (v1.0)

The following features are currently implemented and available:

- [x] **Exact City Lookup** - Fast case-insensitive city name matching
- [x] **Partial Matching** - Multi-field token-based search across city, state, province, and country
- [x] **ISO Code Search** - ISO2 and ISO3 country code lookups
- [x] **Advanced Search** - Configurable case-sensitive and exact match options
- [x] **LRU Caching** - Automatic least-recently-used eviction with 1000 entry default limit
- [x] **Performance Optimizations** - Lazy loading, caching, concurrent access
- [x] **Security Features** - Input validation, XSS prevention, injection protection
- [x] **CLI Tool** - Full-featured command-line interface with multiple output formats
- [x] **Cache Management** - Statistics API with hit/miss/eviction tracking
- [x] **Thread Safety** - Full concurrent access support with proper synchronization
- [x] **Comprehensive Testing** - 96.2% test coverage with race detection
- [x] **Cross-Platform Support** - Linux, macOS, Windows compatibility
- [x] **Zero Dependencies** - No external packages required

## 🚧 Planned (v2.0)

These features are planned for the next major release:

### Core Features
- [ ] **Fuzzy Search** - Levenshtein distance and phonetic matching algorithms
- [ ] **Timezone Conversion** - Convert times between different timezones
- [ ] **DST Calculations** - Automatic daylight saving time offset calculations
- [ ] **Custom Data Support** - Allow users to provide their own city data

### Infrastructure
- [ ] **Web API Wrapper** - RESTful API server with OpenAPI/Swagger documentation
- [ ] **Docker Support** - Official Docker images with Alpine and scratch variants
- [ ] **GraphQL API** - GraphQL interface for advanced queries

### Data Enhancements
- [ ] **Additional Data Sources** - Integrate more city data from multiple sources
- [ ] **Geolocation Integration** - Find nearest city by lat/lng coordinates
- [ ] **Enhanced Metadata** - Add population density, elevation, area codes

## 💡 Under Consideration

Features that are being evaluated for future releases:

- [ ] **Autocomplete/Typeahead** - Fast prefix search for autocomplete UIs
- [ ] **Timezone Boundary Data** - Polygon data for precise timezone boundaries
- [ ] **Historical Timezone Data** - Historical timezone changes and DST rules
- [ ] **City Aliases** - Alternative names and spellings for cities
- [ ] **Distance Calculations** - Calculate distances between cities
- [ ] **Batch Operations** - Optimize for bulk lookups
- [ ] **Streaming API** - Stream large result sets efficiently
- [ ] **Redis Cache Backend** - Optional Redis integration for distributed caching
- [ ] **Prometheus Metrics** - Built-in metrics for monitoring

## 📊 Release Timeline

| Version | Target Date | Status | Focus |
|---------|-------------|--------|-------|
| v1.0.0 | Q4 2024 | ✅ Released | Core functionality, caching, CLI |
| v1.1.0 | Q1 2025 | 📋 Planning | Bug fixes, minor enhancements |
| v2.0.0 | Q2 2025 | 🎯 Planned | Fuzzy search, timezone conversion, API server |
| v3.0.0 | Q4 2025 | 💭 Concept | Advanced features, integrations |

## 🎯 Contributing

We welcome contributions to help implement these features! If you're interested in working on any of these items:

1. Check the [Issues](https://github.com/richoandika/city-timezones-go/issues) for related discussions
2. Comment on an issue to express interest or ask questions
3. Review our [Contributing Guide](../CONTRIBUTING.md)
4. Submit a pull request

## 📝 Suggesting New Features

Have an idea not listed here? We'd love to hear it!

1. Check if a similar idea already exists in [Issues](https://github.com/richoandika/city-timezones-go/issues)
2. Create a new feature request with:
   - Clear description of the feature
   - Use cases and benefits
   - Potential implementation approach
   - Any examples from other projects

## 📅 Last Updated

**Last Updated:** January 2025

---

[⬆ Back to README](../README.md)
