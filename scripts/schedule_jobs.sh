#!/bin/bash

FUNCTION_NAME="schedule"
PROJECT_ID=$CAPY_PROJECT_ID
REGION=$CAPY_SERVER_REGION

gcloud scheduler jobs create http schedule-message-1 \
  --project $PROJECT_ID \
  --uri="https://$REGION-$PROJECT_ID.cloudfunctions.net/$FUNCTION_NAME" \
  --schedule="0 0 * * *" \
  --time-zone="UTC" \
  --http-method=GET \
  --location=$REGION

gcloud scheduler jobs create http schedule-message-2 \
  --project $PROJECT_ID \
  --uri="https://$REGION-$PROJECT_ID.cloudfunctions.net/$FUNCTION_NAME" \
  --schedule="0 12 * * *" \
  --time-zone="UTC" \
  --http-method=GET \
  --location=$REGION

echo "Setup complete!"