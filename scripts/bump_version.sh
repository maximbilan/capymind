#!/bin/bash

# Get current version
source ./scripts/get_version.sh
CURRENT_VERSION=$(get_version)

if [[ $? -ne 0 ]]; then
  echo "Error: Failed to read the current version."
  exit 1
fi

# Split the version into components
IFS='.' read -r MAJOR MINOR PATCH <<< "$CURRENT_VERSION"

# Parse arguments
UPDATE_PART="PATCH"
if [[ "$1" == "--minor" ]]; then
  UPDATE_PART="MINOR"
elif [[ "$1" != "" ]]; then
  echo "Usage: $0 [--minor]"
  echo "  --minor  Increment the MINOR version (default is PATCH)"
  exit 1
fi

# Update the version based on the argument
if [[ "$UPDATE_PART" == "MINOR" ]]; then
  MINOR=$((MINOR + 1))
  PATCH=0  # Reset PATCH when MINOR is incremented
else
  PATCH=$((PATCH + 1))
fi

# Construct the new version
NEW_VERSION="${MAJOR}.${MINOR}.${PATCH}"

# Write the new version to the VERSION file
SCRIPT_DIR=$(dirname "$0")
VERSION_FILE="$SCRIPT_DIR/../VERSION"
echo "$NEW_VERSION" > "$VERSION_FILE"

# Print success message
echo "Version bumped from $CURRENT_VERSION to $NEW_VERSION in $VERSION_FILE"
