package firestore

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// Path to the credentials file
func credentialsPath() string {
	var path = "credentials.json"
	if os.Getenv("DEBUG_MODE") == "true" {
		path = "./../" + path
	} else {
		path = "./" + path
	}
	return path
}

// Client for Firestore
func NewClient(ctx context.Context) (*firestore.Client, error) {
	projectID := os.Getenv("CAPY_PROJECT_ID")
	var client *firestore.Client
	var err error

	if os.Getenv("CLOUD") == "true" {
		client, err = firestore.NewClient(ctx, projectID)
	} else {
		path := credentialsPath()
		client, err = firestore.NewClient(ctx, projectID, option.WithCredentialsFile(path))
	}

	if err != nil {
		return nil, err
	}

	return client, nil
}
