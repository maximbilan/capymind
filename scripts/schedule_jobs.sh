#!/bin/bash

FUNCTION_NAME="schedule"
PROJECT_ID=$CAPY_PROJECT_ID
REGION=$CAPY_SERVER_REGION

JOB1="schedule-morning-messages"
JOB2="schedule-evening-messages"
JOB3="schedule-weekly-analysis"
JOB4="schedule-motivational-messages"

JOBS=($JOB1 $JOB2 $JOB3 $JOB4)

for JOB in "${JOBS[@]}"; do
  gcloud scheduler jobs delete "$JOB" \
    --project "$PROJECT_ID" \
    --location "$REGION" \
    --quiet
done

gcloud scheduler jobs create http $JOB1 \
  --project $PROJECT_ID \
  --uri="https://$REGION-$PROJECT_ID.cloudfunctions.net/$FUNCTION_NAME?type=morning&offset=9" \
  --schedule="0 0 * * *" \
  --http-method=GET \
  --location=$REGION

gcloud scheduler jobs create http $JOB2 \
  --project $PROJECT_ID \
  --uri="https://$REGION-$PROJECT_ID.cloudfunctions.net/$FUNCTION_NAME?type=evening&offset=9" \
  --schedule="0 12 * * *" \
  --http-method=GET \
  --location=$REGION

gcloud scheduler jobs create http $JOB3 \
  --project $PROJECT_ID \
  --uri="https://$REGION-$PROJECT_ID.cloudfunctions.net/$FUNCTION_NAME?type=weekly_analysis&offset=5" \
  --schedule="0 12 * * 0" \ # Every Sunday
  --http-method=GET \
  --location=$REGION

gcloud scheduler jobs create http $JOB4 \
  --project $PROJECT_ID \
  --uri="https://$REGION-$PROJECT_ID.cloudfunctions.net/$FUNCTION_NAME?type=user_stats&offset=10" \
  --schedule="0 0 * * 4" \ # Every Thursday
  --http-method=GET \
  --location=$REGION

echo "Setup complete!"