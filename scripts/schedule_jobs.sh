#!/bin/bash

FUNCTION_NAME="schedule"
PROJECT_ID=$CAPY_PROJECT_ID
REGION=$CAPY_SERVER_REGION

gcloud scheduler jobs create http schedule-morning-messages \
  --project $PROJECT_ID \
  --uri="https://$REGION-$PROJECT_ID.cloudfunctions.net/$FUNCTION_NAME?type=morning" \
  --schedule="0 0 * * *" \
  --http-method=GET \
  --location=$REGION

gcloud scheduler jobs create http schedule-evening-messages \
  --project $PROJECT_ID \
  --uri="https://$REGION-$PROJECT_ID.cloudfunctions.net/$FUNCTION_NAME?type=evening" \
  --schedule="0 12 * * *" \
  --http-method=GET \
  --location=$REGION

echo "Setup complete!"