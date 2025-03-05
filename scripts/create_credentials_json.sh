#!/bin/bash
# Download the credentials.json file

gcloud iam service-accounts keys create credentials.json \
    --iam-account $CAPY_SERVICE_ACCOUNT