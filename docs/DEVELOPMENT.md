# 🛠️ Development Guide

This guide provides information for developers who want to contribute to or extend City Timezones Go.

## Table of Contents

- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Development Workflow](#development-workflow)
- [Testing](#testing)
- [Building](#building)
- [Contributing](#contributing)

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git
- Make (optional, but recommended)

### Clone and Setup

```bash
# Clone the repository
git clone https://github.com/richoandika/city-timezones-go.git
cd city-timezones-go

# Download dependencies
go mod download

# Verify setup
go test ./...
```

## Project Structure

This project follows the [Go Standard Project Layout](https://github.com/golang-standards/project-layout):

```
city-timezones-go/
├── .github/                    # GitHub configuration
│   ├── workflows/             # GitHub Actions CI/CD
│   ├── ISSUE_TEMPLATE/        # Issue templates
│   └── CODEOWNERS            # Code ownership
├── cmd/                      # Main applications
│   └── citytimezones/        # CLI application
│       └── main.go           # CLI entry point
├── data/                     # Application data
│   └── cityMap.json         # City timezone data (7,326 cities)
├── docs/                     # Documentation
│   ├── API.md               # API reference
│   ├── FAQ.md               # Frequently asked questions
│   ├── PERFORMANCE.md       # Performance guide
│   ├── ROADMAP.md           # Project roadmap
│   └── DEVELOPMENT.md       # This file
├── examples/                 # Example applications
│   ├── basic/               # Basic usage examples
│   ├── advanced/            # Advanced patterns
│   └── cli/                 # CLI usage examples
├── internal/                 # Private application code
│   └── city/                # Internal city package
│       ├── cache.go         # LRU cache implementation
│       ├── cache_test.go    # Cache tests
│       ├── custom_unmarshaler.go  # JSON unmarshaling
│       ├── errors.go        # Error types
│       ├── loader.go        # Data loading
│       ├── loader_test.go   # Loader tests
│       ├── search.go        # Search functions
│       ├── search_test.go   # Search tests
│       ├── types.go         # Data structures
│       └── validation.go    # Input validation
├── pkg/                      # Public library code
│   └── citytimezones/       # Public API package
│       ├── citytimezones.go      # Public API
│       └── citytimezones_test.go # Public API tests
├── scripts/                  # Build and utility scripts
│   └── test_runner.sh       # Comprehensive test script
├── .goreleaser.yml          # Release automation
├── Makefile                 # Build automation
├── go.mod                   # Go module file
├── go.sum                   # Go dependencies checksum
├── CHANGELOG.md             # Project changelog
├── CONTRIBUTING.md          # Contribution guidelines
├── LICENSE                  # MIT License
├── README.md                # Project README
└── SECURITY.md              # Security policy
```

### Key Directories

**`/internal/city`** - Core implementation
- Contains the main business logic
- Not importable by external packages
- Full test coverage required

**`/pkg/citytimezones`** - Public API
- Thin wrapper around internal package
- Exposes only what's needed externally
- Type aliases for public types

**`/cmd/citytimezones`** - CLI tool
- Command-line interface implementation
- Uses the public API from pkg/

**`/examples`** - Usage examples
- Runnable example applications
- Demonstrates various use cases
- Should always be kept up-to-date

## Development Workflow

### Making Changes

1. **Create a branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes**
   - Write code
   - Add tests
   - Update documentation

3. **Run tests**
   ```bash
   make test
   make test-coverage
   ```

4. **Check formatting and linting**
   ```bash
   make fmt
   make vet
   ```

5. **Commit your changes**
   ```bash
   git add .
   git commit -m "feat: add your feature description"
   ```

### Commit Message Convention

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>: <description>

[optional body]

[optional footer]
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `test`: Adding or updating tests
- `refactor`: Code refactoring
- `perf`: Performance improvements
- `chore`: Maintenance tasks

Examples:
```bash
feat: add fuzzy search support
fix: correct timezone offset calculation
docs: update API examples
test: add concurrent access tests
```

## Testing

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run tests with race detection
go test ./... -race

# Run specific package tests
go test ./internal/city -v

# Run specific test
go test ./internal/city -run TestLookupViaCity
```

### Writing Tests

#### Unit Tests

```go
func TestYourFunction(t *testing.T) {
    t.Run("should handle normal case", func(t *testing.T) {
        result, err := YourFunction("input")
        if err != nil {
            t.Errorf("unexpected error: %v", err)
        }
        if result != expected {
            t.Errorf("expected %v, got %v", expected, result)
        }
    })

    t.Run("should handle error case", func(t *testing.T) {
        _, err := YourFunction("")
        if err == nil {
            t.Error("expected error, got nil")
        }
    })
}
```

#### Benchmark Tests

```go
func BenchmarkYourFunction(b *testing.B) {
    for i := 0; i < b.N; i++ {
        YourFunction("test")
    }
}
```

### Test Coverage Goals

- Minimum: 90% coverage
- Target: 95%+ coverage
- New code: 100% coverage required

Check coverage:
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Building

### Build CLI Tool

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Run the CLI
./bin/citytimezones -city Chicago
```

### Build with Go

```bash
# Build CLI
go build -o bin/citytimezones ./cmd/citytimezones

# Build with version info
go build -ldflags="-X main.version=v1.0.0" -o bin/citytimezones ./cmd/citytimezones
```

### Cross-Compilation

```bash
# Linux AMD64
GOOS=linux GOARCH=amd64 go build -o bin/citytimezones-linux-amd64 ./cmd/citytimezones

# macOS ARM64
GOOS=darwin GOARCH=arm64 go build -o bin/citytimezones-darwin-arm64 ./cmd/citytimezones

# Windows AMD64
GOOS=windows GOARCH=amd64 go build -o bin/citytimezones-windows-amd64.exe ./cmd/citytimezones
```

## Code Quality

### Formatting

```bash
# Format all Go files
make fmt

# Or manually
gofmt -w .
```

### Linting

```bash
# Run go vet
make vet

# Run with all checks
go vet ./...
```

### Static Analysis

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run all linters
golangci-lint run
```

## Debugging

### Enable Verbose Output

```go
// Add debug prints (remove before committing)
import "log"

log.Printf("Debug: value=%v", value)
```

### Use Delve Debugger

```bash
# Install delve
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug tests
dlv test ./internal/city

# Debug CLI
dlv debug ./cmd/citytimezones
```

### Profiling

```bash
# CPU profiling
go test -cpuprofile=cpu.prof -bench=.
go tool pprof cpu.prof

# Memory profiling
go test -memprofile=mem.prof -bench=.
go tool pprof mem.prof
```

## Documentation

### Godoc Comments

All public functions should have godoc comments:

```go
// LookupViaCity searches for cities by exact city name match (case-insensitive).
// It returns all cities that match the given name.
//
// Example:
//   cities, err := LookupViaCity("Chicago")
//   if err != nil {
//       log.Fatal(err)
//   }
//   for _, city := range cities {
//       fmt.Printf("%s, %s\n", city.City, city.Country)
//   }
func LookupViaCity(cityName string) ([]CityData, error) {
    // ...
}
```

### Updating Documentation

When making changes:
1. Update relevant `.md` files in `/docs`
2. Update code examples if APIs change
3. Update README.md if needed
4. Update CHANGELOG.md

## Release Process

Releases are automated using GoReleaser:

1. **Update version**
   ```bash
   # Update CHANGELOG.md
   # Commit changes
   git add CHANGELOG.md
   git commit -m "chore: prepare v1.1.0 release"
   ```

2. **Create and push tag**
   ```bash
   git tag -a v1.1.0 -m "Release v1.1.0"
   git push origin v1.1.0
   ```

3. **GitHub Actions handles the rest**
   - Runs all tests
   - Builds binaries for multiple platforms
   - Creates GitHub release
   - Generates release notes

## Common Tasks

### Adding a New Search Function

1. Add function to `internal/city/search.go`
2. Add tests to `internal/city/search_test.go`
3. Export through `pkg/citytimezones/citytimezones.go`
4. Add public tests to `pkg/citytimezones/citytimezones_test.go`
5. Update documentation
6. Add examples

### Updating City Data

1. Update `data/cityMap.json`
2. Run tests to ensure format is correct
3. Update count in README if significant change
4. Document in CHANGELOG.md

### Adding Dependencies

**Important:** This project aims for zero dependencies. Before adding:

1. Evaluate if it's absolutely necessary
2. Consider implementing the functionality ourselves
3. If needed, propose in an issue first
4. Update go.mod: `go get package@version`

## Troubleshooting

### Tests Failing

```bash
# Clean and rebuild
make clean
go mod tidy
go test ./...
```

### Module Issues

```bash
# Reset modules
rm go.sum
go mod tidy
go mod download
```

### Build Issues

```bash
# Clean build cache
go clean -cache
go clean -modcache
make build
```

## Getting Help

- **Questions:** Open a [Discussion](https://github.com/richoandika/city-timezones-go/discussions)
- **Bugs:** File an [Issue](https://github.com/richoandika/city-timezones-go/issues)
- **Contributing:** See [CONTRIBUTING.md](../CONTRIBUTING.md)

---

[⬆ Back to README](../README.md)
