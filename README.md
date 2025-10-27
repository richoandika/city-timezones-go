<div align="center">

# 🌍 City Timezones Go

[![Go Reference](https://pkg.go.dev/badge/github.com/richoandika/city-timezones-go.svg)](https://pkg.go.dev/github.com/richoandika/city-timezones-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/richoandika/city-timezones-go)](https://goreportcard.com/report/github.com/richoandika/city-timezones-go)
[![CI](https://github.com/richoandika/city-timezones-go/workflows/CI/badge.svg)](https://github.com/richoandika/city-timezones-go/actions)
[![Coverage](https://codecov.io/gh/richoandika/city-timezones-go/branch/main/graph/badge.svg)](https://codecov.io/gh/richoandika/city-timezones-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue.svg)](https://golang.org/)
[![Release](https://img.shields.io/github/release/richoandika/city-timezones-go.svg)](https://github.com/richoandika/city-timezones-go/releases)

**A high-performance Go library for city timezone lookups**

7,326 cities • Zero dependencies • Thread-safe • 96.2% test coverage

[Features](#features) • [Installation](#installation) • [Quick Start](#quick-start) • [Documentation](#documentation) • [Contributing](#contributing)

</div>

---

## Overview

City Timezones Go is a fast, efficient Go library for looking up timezone information by city name. It provides multiple search strategies including exact matching, partial matching, and ISO code lookups with a simple, intuitive API.

> **Note:** This is a Go port of the [city-timezones](https://github.com/kevinroberts/city-timezones) JavaScript library by [Kevin Roberts](https://github.com/kevinroberts), delivering **10x performance improvements** and full type safety.

## Features

| Feature | Description |
|---------|-------------|
| 🔍 **Multiple Search Strategies** | Exact match, partial matching, ISO code lookups |
| ⚡ **High Performance** | LRU caching, lazy loading, 10x faster than JavaScript |
| 🔒 **Thread-Safe** | Full concurrent access support with proper synchronization |
| 🛡️ **Security First** | Input validation, XSS prevention, injection protection |
| 📦 **Zero Dependencies** | No external packages required |
| 🧪 **Well-Tested** | 96.2% test coverage with race condition testing |
| 🌐 **Comprehensive Data** | 7,326 cities with timezone information |
| 🖥️ **CLI Tool** | Command-line interface included |

## Installation

```bash
go get github.com/richoandika/city-timezones-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/richoandika/city-timezones-go/pkg/citytimezones"
)

func main() {
    // Look up a city by name
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

## Core API

### Search Functions

```go
// Exact city name match (case-insensitive)
cities, err := citytimezones.LookupViaCity("Chicago")

// Partial matching across multiple fields
cities, err := citytimezones.FindFromCityStateProvince("springfield mo")

// Search by ISO2 or ISO3 country codes
cities, err := citytimezones.FindFromIsoCode("DE")

// Advanced search with options
options := citytimezones.SearchOptions{
    CaseSensitive: true,
    ExactMatch:    false,
}
cities, err := citytimezones.SearchCities("Chicago", options)

// Get all cities
allCities, err := citytimezones.GetCityMapping()
```

### Cache Management

```go
// Get cache statistics
stats := citytimezones.GetCacheStats()
fmt.Printf("Hit rate: %.2f%%\n", stats.HitRate)

// Check cache size
size := citytimezones.CacheSize()

// Clear cache if needed
citytimezones.ClearCache()
```

## CLI Tool

```bash
# Install
go install github.com/richoandika/city-timezones-go/cmd/citytimezones@latest

# Or build from source
make build

# Search for a city
citytimezones -city Chicago

# Search with partial matching
citytimezones -search "springfield mo"

# Search by ISO code
citytimezones -iso DE -limit 5

# Output as JSON
citytimezones -city Tokyo -output json

# Show version
citytimezones -version
```

## Performance

City Timezones Go is highly optimized:

- **10x faster** than the JavaScript version
- **7.5x less memory** usage (~2MB vs ~15MB)
- **Microsecond lookups** with caching (<100µs)
- **Thread-safe** with minimal lock contention

See [PERFORMANCE.md](docs/PERFORMANCE.md) for detailed benchmarks.

## Documentation

| Document | Description |
|----------|-------------|
| [API Reference](docs/API.md) | Complete API documentation |
| [FAQ](docs/FAQ.md) | Frequently asked questions |
| [Performance Guide](docs/PERFORMANCE.md) | Benchmarks and optimization tips |
| [Development Guide](docs/DEVELOPMENT.md) | Contributing and development workflow |
| [Roadmap](docs/ROADMAP.md) | Feature roadmap and future plans |
| [Examples](examples/) | Runnable code examples |

## Examples

Check out the [`examples/`](examples/) directory for more:

- **[Basic Usage](examples/basic/)** - Simple lookups and searches
- **[Advanced Usage](examples/advanced/)** - Concurrent access, caching, error handling
- **[CLI Usage](examples/cli/)** - Command-line examples

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

Quick start:
```bash
# Clone and setup
git clone https://github.com/richoandika/city-timezones-go.git
cd city-timezones-go
go mod download

# Run tests
make test

# Build
make build
```

See [DEVELOPMENT.md](docs/DEVELOPMENT.md) for detailed development instructions.

## Support

| Type | Link |
|------|------|
| 📖 **Documentation** | [docs/](docs/) |
| 🐛 **Bug Reports** | [Issues](https://github.com/richoandika/city-timezones-go/issues) |
| 💡 **Feature Requests** | [Issues](https://github.com/richoandika/city-timezones-go/issues) |
| 💬 **Discussions** | [Discussions](https://github.com/richoandika/city-timezones-go/discussions) |
| 🔒 **Security** | [SECURITY.md](SECURITY.md) |

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Attribution

This library is a Go port of the original [city-timezones](https://github.com/kevinroberts/city-timezones) JavaScript library created by [Kevin Roberts](https://github.com/kevinroberts).

### Key Improvements

- ⚡ 10x faster execution
- 💾 7.5x less memory usage
- 🔒 Compile-time type safety
- 🧵 Full thread-safe operations
- 🧪 96.2% test coverage
- 📦 Zero dependencies

## Acknowledgments

- **[Kevin Roberts](https://github.com/kevinroberts)** - Original JavaScript library author
- **Go Community** - For excellent tooling and standards
- **Contributors** - All who help improve this project

---

<div align="center">

**Made with ❤️ for the Go community**

If you find this library useful, please consider giving it a ⭐

![Code Size](https://img.shields.io/github/languages/code-size/richoandika/city-timezones-go)
![Last Commit](https://img.shields.io/github/last-commit/richoandika/city-timezones-go)

[⬆ Back to Top](#-city-timezones-go)

</div>
