#!/bin/bash

# Get the directory of this script
SCRIPT_DIR=$(dirname "$0")

# Path to the version.go file (adjusted for the relative path)
VERSION_FILE="$SCRIPT_DIR/../version.go"

# Check if version.go exists
if [[ ! -f "$VERSION_FILE" ]]; then
  echo "Error: $VERSION_FILE not found."
  exit 1
fi

# Path to get_version.sh
GET_VERSION_SCRIPT="$SCRIPT_DIR/get_version.sh"

# Import the current version using the get_version.sh script
CURRENT_VERSION=$("$GET_VERSION_SCRIPT")

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

# Update the version in the file
sed -i.bak -E "s/(const AppVersion = \")[0-9]+\.[0-9]+\.[0-9]+(\".*)/\1${NEW_VERSION}\2/" "$VERSION_FILE"

# Print success message
echo "Version bumped from $CURRENT_VERSION to $NEW_VERSION in $VERSION_FILE"

# Optional: Remove the backup file created by sed
rm -f "${VERSION_FILE}.bak"
