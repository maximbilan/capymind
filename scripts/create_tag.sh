#!/bin/bash

# Get current version
source ./scripts/get_version.sh
CURRENT_VERSION=$(get_version)

# Create a new tag
git tag -a "releases/$CURRENT_VERSION" -m "$CURRENT_VERSION"

# Push the tag to the remote repository
git push origin "releases/$CURRENT_VERSION"