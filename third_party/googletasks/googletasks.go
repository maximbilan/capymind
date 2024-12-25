package googletasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	taskspb "cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
	"github.com/capymind/internal/taskservice"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GoogleTasks struct{}

var client *cloudtasks.Client

// Create cloud tasks client
func (tasks GoogleTasks) Connect(ctx *context.Context) {
	var newClient, err = cloudtasks.NewClient(*ctx)
	if err != nil {
		log.Printf("[Scheduler] Error creating cloud tasks client, %s", err.Error())
	}
	client = newClient
}

// Close cloud tasks client
func (tasks GoogleTasks) Close() {
	client.Close()
}

// Schedule a cloud task
func (tasks GoogleTasks) Schedule(ctx *context.Context, scheduledMessage taskservice.ScheduledTask, timeOffset time.Time) {
	projectID := os.Getenv("CAPY_PROJECT_ID")
	locationID := os.Getenv("CAPY_SERVER_REGION")
	queueID := "messages"
	queuePath := fmt.Sprintf("projects/%s/locations/%s/queues/%s", projectID, locationID, queueID)
	url := fmt.Sprintf("https://%s-%s.cloudfunctions.net/%s", locationID, projectID, "sendMessage")

	timestamp := timestamppb.Timestamp{
		Seconds: timeOffset.Unix(),
		Nanos:   int32(timeOffset.Nanosecond()),
	}

	req := &taskspb.CreateTaskRequest{
		Parent: queuePath,
		Task: &taskspb.Task{
			MessageType: &taskspb.Task_HttpRequest{
				HttpRequest: &taskspb.HttpRequest{
					HttpMethod: taskspb.HttpMethod_POST,
					Url:        url,
				},
			},
			ScheduleTime: &timestamp,
		},
	}

	payload, err := json.Marshal(scheduledMessage)
	if err != nil {
		log.Printf("[Scheduler] Error marshalling payload, %s", err.Error())
		return
	}
	req.Task.GetHttpRequest().Body = payload

	createdTask, err := client.CreateTask(*ctx, req)
	if err != nil {
		log.Printf("[Scheduler] Error scheduling a task, %s", err.Error())
		return
	}

	log.Printf("[Scheduler] Task has been successfully created: %s", createdTask.GetName())
}
