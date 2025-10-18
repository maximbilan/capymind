#!/bin/bash

# Use the active credentials (set by CI auth) to mint an identity token
gcloud auth print-identity-token --audiences="$CAPY_THERAPY_SESSION_URL"