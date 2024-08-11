# Description: This script sets up the gcloud access for the service account

PROJECT_ID=$CAPY_PROJECT_ID
SERVICE_ACCOUNT=$CAPY_SERVICE_ACCOUNT
PROJECT_NUMBER=$CAPY_PROJECT_NUMBER
POOL_ID="githubci"
POOL_NAME="GitHubCI"
LOCATION="global"
PROVIDER_ID="github-provider"
PROVIDER_NAME="GitHub Provider"

gcloud services enable iam.googleapis.com

gcloud iam workload-identity-pools create $POOL_ID \
    --project=$PROJECT_ID \
	--location=$LOCATION \
	--display-name=$POOL_NAME

gcloud iam workload-identity-pools providers create-oidc $PROVIDER_ID \
    --project=$PROJECT_ID \
    --location=$LOCATION \
    --workload-identity-pool=$POOL_ID \
    --display-name=$PROVIDER_NAME \
    --attribute-mapping="google.subject=assertion.sub,attribute.actor=assertion.actor,attribute.repository=assertion.repository,attribute.aud=assertion.aud" \
    --issuer-uri="https://token.actions.githubusercontent.com"

gcloud iam workload-identity-pools providers describe $PROVIDER_ID \
    --project=$PROJECT_ID \
    --location=$LOCATION \
    --workload-identity-pool=$POOL_ID \
    --format="value(name)"

PRINCIPAL_SET="principalSet://iam.googleapis.com/projects/$PROJECT_NUMBER/locations/$LOCATION/workloadIdentityPools/$POOL_ID/attribute.repository/maximbilan/$PROJECT_ID"

gcloud iam service-accounts add-iam-policy-binding $SERVICE_ACCOUNT \
  --project=$PROJECT_ID \
  --role="roles/iam.workloadIdentityUser" \
  --member=$PRINCIPAL_SET