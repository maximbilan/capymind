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

# Create a new tag
git tag -a "releases/$CURRENT_VERSION" -m "$CURRENT_VERSION"

# Push the tag to the remote repository
git push origin "releases/$CURRENT_VERSION"