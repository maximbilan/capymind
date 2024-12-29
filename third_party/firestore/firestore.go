//coverage:ignore file

package firestore

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type Firestore struct{}

var client *firestore.Client

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
func newClient(ctx *context.Context) (*firestore.Client, error) {
	projectID := os.Getenv("CAPY_PROJECT_ID")
	var client *firestore.Client
	var err error

	if os.Getenv("CLOUD") == "true" {
		client, err = firestore.NewClient(*ctx, projectID)
	} else {
		path := credentialsPath()
		client, err = firestore.NewClientWithDatabase(*ctx, projectID, "development", option.WithCredentialsFile(path))
	}

	if err != nil {
		return nil, err
	}

	return client, nil
}

// Create a new Firestore database connection
func (db Firestore) CreateClient(ctx *context.Context) {
	newClient, err := newClient(ctx)
	if err != nil {
		log.Printf("[Firestore] Error creating firestore client, %s", err.Error())
	}
	client = newClient
}

// Close the Firestore database connection
func (db Firestore) CloseClient() {
	client.Close()
}
