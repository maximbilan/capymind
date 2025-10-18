#!/bin/bash

gcloud auth print-identity-token $CAPY_SERVICE_ACCOUNT --audiences=$CAPY_THERAPY_SESSION_URL