package database

import "testing"

func TestCollections(t *testing.T) {
	collections := []Collection{Users, Notes, Feedbacks, Updates}
	for _, collection := range collections {
		if collection.String() != string(collection) {
			t.Fatalf("Expected %v, got %v", collection, collection.String())
		}
	}
}
