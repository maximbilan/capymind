#!/bin/bash

FUNCTION_NAME="schedule"
PROJECT_ID=$CAPY_PROJECT_ID
REGION=$CAPY_SERVER_REGION

JOB1="schedule-morning-messages"
JOB2="schedule-evening-messages"

gcloud scheduler jobs delete $JOB1 \
  --project $PROJECT_ID \
  --location=$REGION \
  --quiet

gcloud scheduler jobs delete $JOB2 \
  --project $PROJECT_ID \
  --location=$REGION \
  --quiet

gcloud scheduler jobs create http $JOB1 \
  --project $PROJECT_ID \
  --uri="https://$REGION-$PROJECT_ID.cloudfunctions.net/$FUNCTION_NAME?type=morning" \
  --schedule="0 0 * * *" \
  --http-method=GET \
  --location=$REGION

gcloud scheduler jobs create http $JOB2 \
  --project $PROJECT_ID \
  --uri="https://$REGION-$PROJECT_ID.cloudfunctions.net/$FUNCTION_NAME?type=evening" \
  --schedule="0 12 * * *" \
  --http-method=GET \
  --location=$REGION

echo "Setup complete!"