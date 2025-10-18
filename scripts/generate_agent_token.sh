#!/bin/bash

SERVICE_ACCOUNT=$CAPY_SERVICE_ACCOUNT
PROJECT_ID=$CAPY_PROJECT_ID

gcloud iam service-accounts add-iam-policy-binding $SERVICE_ACCOUNT \
  --role=roles/iam.serviceAccountTokenCreator \
  --member=serviceAccount:$SERVICE_ACCOUNT \
  --project=$PROJECT_ID

gcloud iam service-accounts keys create key.json \
  --iam-account=$SERVICE_ACCOUNT

# Activate the service account
gcloud auth activate-service-account --key-file=key.json

# Generate the identity token
gcloud auth print-identity-token