#!/bin/bash

set -euo pipefail

if [ -z "${CAPY_THERAPY_SESSION_URL:-}" ]; then
  echo "CAPY_THERAPY_SESSION_URL is required" >&2
  exit 1
fi

args=(auth print-identity-token "--audiences=$CAPY_THERAPY_SESSION_URL")

# Ensure we use service account credentials
if [ -n "${CAPY_SERVICE_ACCOUNT:-}" ]; then
  # Impersonate explicitly if provided
  args+=("--impersonate-service-account=$CAPY_SERVICE_ACCOUNT")
else
  # Verify the active account is a service account
  ACTIVE_ACCOUNT="$(gcloud config get-value account 2>/dev/null || true)"
  if ! echo "$ACTIVE_ACCOUNT" | grep -qi "gserviceaccount.com"; then
    echo "Active gcloud account '$ACTIVE_ACCOUNT' is not a service account. Set CAPY_SERVICE_ACCOUNT to impersonate one." >&2
    exit 1
  fi
fi

gcloud "${args[@]}"