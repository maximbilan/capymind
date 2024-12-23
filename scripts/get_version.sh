#!/bin/bash

get_version() {
  local script_dir version_file
  script_dir=$(dirname "$0")
  version_file="$script_dir/../VERSION"

  if [[ ! -f "$version_file" ]]; then
    echo "Error: $version_file not found." >&2
    exit 1
  fi

  cat "$version_file"
}

CURRENT_VERSION=$(get_version)
echo "The current version is: $CURRENT_VERSION"