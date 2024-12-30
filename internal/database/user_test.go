package database

import (
	"testing"
	"time"
)

func TestIsActive(t *testing.T) {
	yesterday := time.Now().AddDate(0, 0, -1)

	user := User{
		Timestamp: &yesterday,
	}
	if !user.IsActive() {
		t.Error("Expected user to be active, got inactive")
	}

	tenDaysAgo := time.Now().AddDate(0, 0, -10)
	user.Timestamp = &tenDaysAgo

	if user.IsActive() {
		t.Error("Expected user to be inactive, got active")
	}
}
