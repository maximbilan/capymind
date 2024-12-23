#!/bin/bash

# File containing the version constant (relative to the script location)
VERSION_FILE="./version.go"

# Pattern to match the version line
VERSION_PATTERN='const AppVersion = "'

# Check if version.go exists
if [[ ! -f "$VERSION_FILE" ]]; then
  echo "Error: $VERSION_FILE not found."
  exit 1
fi

# Extract the current version
CURRENT_VERSION=$(grep "$VERSION_PATTERN" "$VERSION_FILE" | sed -E 's/.*"([0-9]+\.[0-9]+\.[0-9]+)".*/\1/')

if [[ -z "$CURRENT_VERSION" ]]; then
  echo "Error: Could not find the current version in $VERSION_FILE"
  exit 1
fi

# Output the current version
echo "$CURRENT_VERSION"