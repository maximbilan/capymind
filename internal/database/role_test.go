package database

import (
	"testing"
)

func TestIsAdmin(t *testing.T) {
	role1 := Admin
	result := IsAdmin(&role1)
	if result != true {
		t.Fatalf("Expected true, got %v", result)
	}
}
