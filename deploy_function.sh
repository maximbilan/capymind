#!/bin/bash

# Set the function name
FUNCTION_NAME="Handler"
# Set the entry point
ENTRY_POINT="Handler"
# Set the runtime
RUNTIME="go122"
# Set the project ID
PROJECT_ID=$CAPY_PROJECT_ID
# Set the region
REGION=$CAPY_SERVER_REGION

# Deploy the function
gcloud functions deploy $FUNCTION_NAME \
    --runtime $RUNTIME \
    --trigger-http \
    --allow-unauthenticated \
    --entry-point $ENTRY_POINT \
    --project $PROJECT_ID \
    --gen2 \
    --region $REGION

# Print the deployment status
if [ $? -eq 0 ]; then
    echo "Function $FUNCTION_NAME deployed successfully."
else
    echo "Failed to deploy function $FUNCTION_NAME."
fi