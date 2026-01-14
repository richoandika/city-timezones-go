# City Timezones Go

[![Go Reference](https://pkg.go.dev/badge/github.com/richoandika/city-timezones-go.svg)](https://pkg.go.dev/github.com/richoandika/city-timezones-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/richoandika/city-timezones-go)](https://goreportcard.com/report/github.com/richoandika/city-timezones-go)
[![CI](https://github.com/richoandika/city-timezones-go/workflows/CI/badge.svg)](https://github.com/richoandika/city-timezones-go/actions)
[![Coverage](https://codecov.io/gh/richoandika/city-timezones-go/branch/main/graph/badge.svg)](https://codecov.io/gh/richoandika/city-timezones-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue.svg)](https://golang.org/)

A high-performance, thread-safe Go library for city timezone lookups with comprehensive coverage of 7,326 cities worldwide.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Usage](#usage)
  - [Basic Lookups](#basic-lookups)
  - [Advanced Searches](#advanced-searches)
  - [Cache Management](#cache-management)
- [CLI Tool](#cli-tool)
- [Performance](#performance)
- [Documentation](#documentation)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Overview

City Timezones Go provides fast and efficient timezone information lookups by city name, supporting multiple search strategies including exact matching, partial matching, and ISO code queries. Built with performance and reliability in mind, it features zero external dependencies, comprehensive test coverage, and full thread safety.

This library is a Go port of the [city-timezones](https://github.com/kevinroberts/city-timezones) JavaScript library by Kevin Roberts, offering significant performance improvements and compile-time type safety.

## Features

- **Multiple Search Strategies**: Exact match, partial matching, and ISO code lookups
- **High Performance**: LRU caching with lazy loading, optimized for speed
- **Thread-Safe**: Full concurrent access support with proper synchronization
- **Security Hardened**: Input validation, XSS prevention, and injection protection
- **Zero Dependencies**: No external packages required
- **Well Tested**: 96.2% test coverage including race condition testing
- **Comprehensive Dataset**: 7,326 cities with accurate timezone information
- **CLI Tool**: Command-line interface for interactive queries

## Installation

```bash
go get github.com/richoandika/city-timezones-go
```

**Requirements**: Go 1.21 or higher

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/richoandika/city-timezones-go/pkg/citytimezones"
)

func main() {
    cities, err := citytimezones.LookupViaCity("Chicago")
    if err != nil {
        log.Fatal(err)
    }

    if len(cities) > 0 {
        city := cities[0]
        fmt.Printf("City: %s, %s\n", city.City, city.Province)
        fmt.Printf("Timezone: %s\n", city.Timezone)
        fmt.Printf("Coordinates: %.4f, %.4f\n", city.Lat, city.Lng)
    }
}
```

## Usage

### Basic Lookups

```go
// Exact city name match (case-insensitive)
cities, err := citytimezones.LookupViaCity("Chicago")

// Partial matching across city, state, and province
cities, err := citytimezones.FindFromCityStateProvince("springfield mo")

// Search by ISO2 or ISO3 country codes
cities, err := citytimezones.FindFromIsoCode("DE")

// Retrieve all cities
allCities, err := citytimezones.GetCityMapping()
```

### Advanced Searches

```go
// Configure search options
options := citytimezones.SearchOptions{
    CaseSensitive: true,
    ExactMatch:    false,
}

cities, err := citytimezones.SearchCities("Chicago", options)
if err != nil {
    log.Fatal(err)
}

for _, city := range cities {
    fmt.Printf("%s, %s (%s)\n", city.City, city.Country, city.Timezone)
}
```

### Cache Management

```go
// Retrieve cache statistics
stats := citytimezones.GetCacheStats()
fmt.Printf("Cache hit rate: %.2f%%\n", stats.HitRate)
fmt.Printf("Total hits: %d\n", stats.Hits)
fmt.Printf("Total misses: %d\n", stats.Misses)

// Check current cache size
size := citytimezones.CacheSize()
fmt.Printf("Cached entries: %d\n", size)

// Clear cache when needed
citytimezones.ClearCache()
```

## CLI Tool

Install the command-line tool:

```bash
go install github.com/richoandika/city-timezones-go/cmd/citytimezones@latest
```

Or build from source:

```bash
make build
```

### CLI Examples

```bash
# Search for a specific city
citytimezones -city Chicago

# Partial search with multiple keywords
citytimezones -search "springfield mo"

# Search by ISO country code
citytimezones -iso DE -limit 5

# Output results as JSON
citytimezones -city Tokyo -output json

# Display version information
citytimezones -version
```

## Performance

City Timezones Go delivers exceptional performance:

- **10x faster** than the original JavaScript implementation
- **7.5x lower memory footprint** (~2MB vs ~15MB)
- **Sub-100µs lookups** with LRU caching enabled
- **Thread-safe operations** with minimal lock contention

For detailed benchmarks and optimization guidelines, see [PERFORMANCE.md](docs/PERFORMANCE.md).

## Documentation

Comprehensive documentation is available in the `docs/` directory:

- [API Reference](docs/API.md) - Complete API documentation with examples
- [Performance Guide](docs/PERFORMANCE.md) - Benchmarks and optimization techniques
- [Development Guide](docs/DEVELOPMENT.md) - Contributing and development workflow
- [FAQ](docs/FAQ.md) - Frequently asked questions
- [Roadmap](docs/ROADMAP.md) - Future features and planned improvements

## Examples

The [`examples/`](examples/) directory contains practical code samples:

- [Basic Usage](examples/basic/) - Simple lookups and search operations
- [Advanced Usage](examples/advanced/) - Concurrent access, caching strategies, error handling
- [CLI Usage](examples/cli/) - Command-line interface examples

## Contributing

Contributions are welcome! Please read our [Contributing Guide](CONTRIBUTING.md) before submitting pull requests.

### Development Setup

```bash
# Clone the repository
git clone https://github.com/richoandika/city-timezones-go.git
cd city-timezones-go

# Install dependencies
go mod download

# Run tests
make test

# Build the project
make build
```

For detailed development instructions, see [DEVELOPMENT.md](docs/DEVELOPMENT.md).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

This library is a Go port of the [city-timezones](https://github.com/kevinroberts/city-timezones) JavaScript library created by Kevin Roberts. The Go implementation introduces significant performance improvements, type safety, and thread-safe operations while maintaining compatibility with the original dataset.

Special thanks to:

- **Kevin Roberts** - Original JavaScript library author
- **Go Community** - For excellent tooling and development standards
- **All Contributors** - For their valuable contributions to this project

---

**Project Status**: Active maintenance and development

For bug reports and feature requests, please use the [GitHub Issues](https://github.com/richoandika/city-timezones-go/issues) page.

For security vulnerabilities, see our [Security Policy](SECURITY.md).
