#!/bin/bash

# Create a temporary folder
TEMP_DIR="deployment"
mkdir -p $TEMP_DIR

# Copy the function file to the temporary folder
cp function.go $TEMP_DIR/function.go

# copy needed folders to the temporary folder
cp -r firestore $TEMP_DIR/
cp -r localizer $TEMP_DIR/
cp -r telegram $TEMP_DIR/
cp -r utils $TEMP_DIR/

# Copy go.mod to the temporary folder
cp go.mod $TEMP_DIR/

# Copy credentials.json to the temporary folder
cp credentials.json $TEMP_DIR/

# Replace "Handler" with "handler" in the function.go file
sed -i '' 's/Handler/handler/g' $TEMP_DIR/function.go

# SZip the contents of the temporary folder
ZIP_FILE="deploy.zip"
zip -r $TEMP_DIR/$ZIP_FILE $TEMP_DIR

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
PARAMS=("CAPY_TELEGRAM_BOT_TOKEN=$CAPY_TELEGRAM_BOT_TOKEN" "CAPY_PROJECT_ID=$CAPY_PROJECT_ID")
ENV_VARS=""
for PARAM in "${PARAMS[@]}"; do
  ENV_VARS+="$PARAM,"
done
ENV_VARS=${ENV_VARS%,}

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
    --source $TEMP_DIR \
    --set-env-vars $ENV_VARS \
    --memory $MEMORY

# Print the deployment status
if [ $? -eq 0 ]; then
    echo "Function $FUNCTION_NAME deployed successfully."
else
    echo "Failed to deploy function $FUNCTION_NAME."
fi

# Clean up
rm -rf $TEMP_DIR
