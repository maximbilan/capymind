#!/bin/bash

# File containing the version constant (relative to the script location)
VERSION_FILE="./version.go"

# Check if version.go exists
if [[ ! -f "$VERSION_FILE" ]]; then
  echo "Error: $VERSION_FILE not found."
  exit 1
fi

# Pattern to match the version line
VERSION_PATTERN='const AppVersion = "'

# Extract the current version
CURRENT_VERSION=$(grep "$VERSION_PATTERN" "$VERSION_FILE" | sed -E 's/.*"([0-9]+\.[0-9]+\.[0-9]+)".*/\1/')

if [[ -z "$CURRENT_VERSION" ]]; then
  echo "Error: Could not find the current version in $VERSION_FILE"
  exit 1
fi

# Split the version into components
IFS='.' read -r MAJOR MINOR PATCH <<< "$CURRENT_VERSION"

# Increment the PATCH version
PATCH=$((PATCH + 1))

# Construct the new version
NEW_VERSION="${MAJOR}.${MINOR}.${PATCH}"

# Update the version in the file
sed -i.bak -E "s/$VERSION_PATTERN[0-9]+\.[0-9]+\.[0-9]+\"/$VERSION_PATTERN${NEW_VERSION}\"/" "$VERSION_FILE"

# Print success message
echo "Version bumped from $CURRENT_VERSION to $NEW_VERSION in $VERSION_FILE"

# Optional: Remove the backup file created by sed
rm -f "${VERSION_FILE}.bak"