package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	taskspb "cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
	firestoreDB "cloud.google.com/go/firestore"
	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func createDBClient(ctx context.Context) *firestoreDB.Client {
	var client, err = firestore.NewClient(ctx)
	if err != nil {
		log.Printf("[Scheduler] Error creating firestore client, %s", err.Error())
	}
	return client
}

func createTasksClient(ctx context.Context) *cloudtasks.Client {
	var client, err = cloudtasks.NewClient(ctx)
	if err != nil {
		log.Printf("[Scheduler] Error creating cloud tasks client, %s", err.Error())
	}
	return client
}

func scheduleTask(ctx context.Context, client *cloudtasks.Client, chatId int, text string, timeOffset time.Time) {
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

	// Create a ScheduledMesaage object
	scheduledMessage := ScheduledMessage{
		ChatId: chatId,
		Text:   text,
	}

	// Pass the ScheduledMessage object to the task as a payload
	payload, err := json.Marshal(scheduledMessage)
	if err != nil {
		log.Printf("[Scheduler] Error marshalling payload, %s", err.Error())
		return
	}

	// Pass the payload to the task
	req.Task.GetHttpRequest().Body = payload

	createdTask, err := client.CreateTask(ctx, req)
	if err != nil {
		log.Printf("[Scheduler] Error scheduling a task, %s", err.Error())
		return
	}

	log.Printf("[Scheduler] Task has been successfully created: %s", createdTask.GetName())
}

func Schedule(w http.ResponseWriter, r *http.Request) {
	log.Println("Schedule capymind...")

	ctx := context.Background()

	// Firestore
	dbClient := createDBClient(ctx)
	defer dbClient.Close()

	// Cloud Tasks
	tasksClient := createTasksClient(ctx)
	defer tasksClient.Close()

	firestore.ForEachUser(ctx, dbClient, func(users []firestore.User) error {
		for _, user := range users {
			log.Printf("[Scheduler] User: %s", user.ID)
			// Handle empty fields later
			userLocale := translator.Locale(user.Locale)
			localizedMessage := translator.Translate(userLocale, "how_are_you")
			scheduledTime := time.Now().Add(9 * time.Hour)
			scheduledTime = scheduledTime.Add(-time.Duration(user.SecondsFromUTC) * time.Second)

			scheduleTask(ctx, tasksClient, user.LastChatId, localizedMessage, scheduledTime)
		}
		return nil
	})
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	var msg ScheduledMessage
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Printf("[Scheduler] Could not parse message %s", err.Error())
		return
	}

	telegram.SendMessage(msg.ChatId, msg.Text, nil, nil)
}
