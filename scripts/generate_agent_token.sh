#!/bin/bash

set -euo pipefail

if [ -z "${CAPY_THERAPY_SESSION_URL:-}" ]; then
  echo "CAPY_THERAPY_SESSION_URL is required" >&2
  exit 1
fi

args=(auth print-identity-token "--audiences=$CAPY_THERAPY_SESSION_URL")

# If a service account is provided, impersonate it to ensure SA credentials
if [ -n "${CAPY_SERVICE_ACCOUNT:-}" ]; then
  args+=("--impersonate-service-account=$CAPY_SERVICE_ACCOUNT")
fi

gcloud "${args[@]}"