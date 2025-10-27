#!/bin/bash

# City Timezones CLI - Usage Examples
# This script demonstrates various CLI usage patterns

echo "=== City Timezones CLI - Usage Examples ==="
echo ""

# Make sure the CLI is built
if [ ! -f "../../bin/citytimezones" ]; then
    echo "Building CLI tool..."
    cd ../.. && make build && cd examples/cli
    echo ""
fi

CLI="../../bin/citytimezones"

# Example 1: Simple city lookup
echo "1. Simple city lookup:"
echo "   Command: $CLI -city Chicago"
$CLI -city Chicago
echo ""

# Example 2: Partial search
echo "2. Partial search:"
echo "   Command: $CLI -search \"springfield mo\""
$CLI -search "springfield mo"
echo ""

# Example 3: ISO code lookup with limit
echo "3. ISO code lookup (limited results):"
echo "   Command: $CLI -iso DE -limit 5"
$CLI -iso DE -limit 5
echo ""

# Example 4: Timezone filter
echo "4. Filter by timezone:"
echo "   Command: $CLI -timezone \"America/New_York\" -limit 3"
$CLI -timezone "America/New_York" -limit 3
echo ""

# Example 5: JSON output
echo "5. JSON output format:"
echo "   Command: $CLI -city Tokyo -output json"
$CLI -city Tokyo -output json
echo ""

# Example 6: Multiple major cities
echo "6. Looking up multiple cities:"
cities=("London" "Paris" "Tokyo" "Sydney")
for city in "${cities[@]}"; do
    echo "   Lookup: $city"
    $CLI -city "$city" -limit 1 | head -n 3
    echo ""
done

echo "=== Examples Complete ==="
