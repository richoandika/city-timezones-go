# Examples

This directory contains practical examples demonstrating how to use the city-timezones-go library.

## Directory Structure

```
examples/
├── basic/           # Basic usage examples
├── advanced/        # Advanced features and patterns
└── cli/            # Command-line tool usage examples
```

## Running the Examples

### Basic Examples

Simple examples showing core functionality:

```bash
# Run from project root
cd examples/basic
go run main.go
```

This demonstrates:
- Looking up cities by exact name
- Partial matching searches
- ISO code lookups
- Handling empty results
- Getting all cities

### Advanced Examples

Advanced patterns including concurrency and caching:

```bash
# Run from project root
cd examples/advanced
go run main.go
```

This demonstrates:
- Custom search options (case-sensitive, exact match)
- Thread-safe concurrent lookups
- Cache performance optimization
- Proper error handling patterns

### CLI Examples

Command-line tool usage examples:

```bash
# Run from project root
cd examples/cli
./usage.sh
```

This demonstrates:
- Simple city lookups
- Partial searches
- ISO code filtering
- Timezone filtering
- JSON output format
- Batch lookups

## Quick Start

If you're new to the library, start with the basic examples:

1. **First**: Run `examples/basic/main.go` to understand core features
2. **Then**: Run `examples/advanced/main.go` to see advanced patterns
3. **Finally**: Try `examples/cli/usage.sh` to explore CLI options

## Building the Examples

All examples can be run directly with `go run`, but you can also build them:

```bash
# Build basic example
cd examples/basic
go build -o basic-demo main.go
./basic-demo

# Build advanced example
cd examples/advanced
go build -o advanced-demo main.go
./advanced-demo
```

## Example Output

### Basic Example Output:
```
=== City Timezones - Basic Examples ===

1. Looking up Chicago:
   City: Chicago, Illinois
   Country: United States of America (US)
   Timezone: America/Chicago
   Coordinates: 41.8500, -87.6500
   Population: 2695598
...
```

### Advanced Example Output:
```
=== City Timezones - Advanced Examples ===

1. Advanced Search with Options:
   Case-sensitive search for 'Chicago': found 1 results
   Exact match search for 'chicago': found 1 results

2. Concurrent Lookups (Thread Safety Demo):
   ✓ Tokyo -> Asia/Tokyo
   ✓ London -> Europe/London
   ✓ New York -> America/New_York
...
```

## Integration with Your Project

To use these patterns in your own project:

```go
import "github.com/richoandika/city-timezones-go/pkg/citytimezones"

// Copy patterns from examples/basic or examples/advanced
cities, err := citytimezones.LookupViaCity("YourCity")
if err != nil {
    // Handle error
}
// Use results
```

## Need Help?

- See the [main README](../README.md) for full documentation
- Check the [API documentation](../docs/API.md) for detailed function reference
- Review [Contributing Guide](../CONTRIBUTING.md) if you want to add more examples

## Contributing Examples

Have a useful example? We'd love to include it! Please:

1. Add your example to the appropriate directory
2. Update this README with a description
3. Ensure it follows Go best practices
4. Test that it runs without errors
5. Submit a pull request

Examples we'd love to see:
- Web API integration
- Database storage patterns
- Custom caching strategies
- Performance optimization techniques
- Real-world application use cases
