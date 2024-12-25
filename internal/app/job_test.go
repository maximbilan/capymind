package app

import (
	"testing"
)

func TestJobFromMessage(t *testing.T) {
	job := createJob("/language en", nil)
	if job == nil {
		t.Fatalf("Job is nil")
	}
	if job.Command != Language {
		t.Fatalf("Command is not language")
	}
	if len(job.Parameters) == 0 {
		t.Fatalf("Parameters are not empty")
	}
	if job.Parameters[0] != "en" {
		t.Fatalf("Parameter is not en")
	}
	if *job.Input != "/language en" {
		t.Fatalf("Input is not /language en")
	}
}

func TestJobFromCallbackQuery(t *testing.T) {
	job := createJob("/timezone 25200", nil)
	if job == nil {
		t.Fatalf("Job is nil")
	}
	if job.Command != Timezone {
		t.Fatalf("Command is not timezone")
	}
	if len(job.Parameters) == 0 {
		t.Fatalf("Parameters are not empty")
	}
	if job.Parameters[0] != "25200" {
		t.Fatalf("Parameter is not 25200")
	}
	if *job.Input != "/timezone 25200" {
		t.Fatalf("Input is not /timezone 25200")
	}
}

func TestNil(t *testing.T) {
	job := createJob("", nil)
	if job != nil {
		t.Fatalf("Job is not nil")
	}
}
