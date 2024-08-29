#!/bin/bash

# Set the function name
FUNCTION_NAME="handler"
# Set the entry point
ENTRY_POINT="handler"
# Set the runtime
RUNTIME="go122"
# Set the project ID
PROJECT_ID=$CAPY_PROJECT_ID
# Set the region
REGION=$CAPY_SERVER_REGION

# Set environment variables
ENV_PARAMS=("CAPY_PROJECT_ID=$CAPY_PROJECT_ID" "CLOUD=true")
ENV_VARS=""
for PARAM in "${ENV_PARAMS[@]}"; do
  ENV_VARS+="$PARAM,"
done
ENV_VARS=${ENV_VARS%,}

# Set the secret environment variables
SECRET_PARAMS=("CAPY_TELEGRAM_BOT_TOKEN=telegram_bot_token")
SECRETS=""
for PARAM in "${SECRET_PARAMS[@]}"; do
  SECRETS+="$PARAM:latest,"
done
SECRETS=${SECRETS%,}

# Set memory parameter
MEMORY="256MB"

# Deploy the function
gcloud functions deploy $FUNCTION_NAME \
    --runtime $RUNTIME \
    --trigger-http \
    --allow-unauthenticated \
    --entry-point $ENTRY_POINT \
    --project $PROJECT_ID \
    --gen2 \
    --region $REGION \
    --set-env-vars $ENV_VARS \
    --set-secrets $SECRETS \
    --memory $MEMORY

# Print the deployment status
if [ $? -eq 0 ]; then
    echo "Function $FUNCTION_NAME deployed successfully."
else
    echo "Failed to deploy function $FUNCTION_NAME."
fi
