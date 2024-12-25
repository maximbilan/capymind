package database

import (
	"testing"
)

func TestIsAdmin(t *testing.T) {
	role := Admin
	result := IsAdmin(&role)
	if result != true {
		t.Fatalf("Expected true, got %v", result)
	}
}
