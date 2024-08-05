#!/bin/bash

# Create a temporary folder
TEMP_DIR="deployment"
mkdir -p $TEMP_DIR

# Copy cloud/cloud.go to the temporary folder and rename it to function.go
cp cloud/cloud.go $TEMP_DIR/function.go

# Step 4: Copy go.mod to the temporary folder
cp go.mod $TEMP_DIR/

# SZip the contents of the temporary folder
ZIP_FILE="deploy.zip"
zip -r $TEMP_DIR/$ZIP_FILE $TEMP_DIR

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
    --region $REGION \
    --source $TEMP_DIR

# Print the deployment status
if [ $? -eq 0 ]; then
    echo "Function $FUNCTION_NAME deployed successfully."
else
    echo "Failed to deploy function $FUNCTION_NAME."
fi

# Clean up
# rm -rf $TEMP_DIR
# rm $ZIP_FILE