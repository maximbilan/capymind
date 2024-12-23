#!/bin/bash

# Open ./version.go file

# Read the version number from the file
VERSION = $(grep "const Version" version.go | sed -E 's/.*"(.+)"$/\1/')

echo "Current version: $VERSION"

# Increment the version number
VERSION = $(echo $VERSION | awk -F. '{$NF = $NF + 1;} 1' | sed 's/ /./g')

echo "New version: $VERSION"

# Write the new version number to the file
sed -i '' "s/const Version = \".*\"/const Version = \"$VERSION\"/" version.go