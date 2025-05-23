package helpers

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/translator"
)

type statFunc func(ctx *context.Context, locale translator.Locale, adminStorage database.AdminStorage) *string
type feedbackFunc func(ctx *context.Context, locale translator.Locale, feedbackStorage database.FeedbackStorage) []string

var wg sync.WaitGroup

func GetTotalUserCount(ctx *context.Context, locale translator.Locale, adminStorage database.AdminStorage) *string {
	count, err := adminStorage.GetTotalUserCount(ctx)
	if err != nil {
		log.Printf("[Admin] Error during fetching total user count: %v", err)
		return nil
	}
	message := fmt.Sprintf(translator.Translate(locale, "total_user_count"), count)
	return &message
}

func GetTotalActiveUserCount(ctx *context.Context, locale translator.Locale, adminStorage database.AdminStorage) *string {
	count, err := adminStorage.GetActiveUserCount(ctx)
	if err != nil {
		log.Printf("[Admin] Error during fetching active user count: %v", err)
		return nil
	}
	message := fmt.Sprintf(translator.Translate(locale, "total_active_user_count"), count)
	return &message
}

func GetTotalNoteCount(ctx *context.Context, locale translator.Locale, adminStorage database.AdminStorage) *string {
	count, err := adminStorage.GetTotalNoteCount(ctx)
	if err != nil {
		log.Printf("[Admin] Error during fetching total note count: %v", err)
		return nil
	}
	message := fmt.Sprintf(translator.Translate(locale, "total_note_count"), count)
	return &message
}

func GetStats(ctx *context.Context, locale translator.Locale, adminStorage database.AdminStorage, feedbackStorage database.FeedbackStorage) []string {
	totalUserCount := waitForStatFunction(GetTotalUserCount, ctx, locale, adminStorage)
	totalActiveUserCount := waitForStatFunction(GetTotalActiveUserCount, ctx, locale, adminStorage)
	totalNoteCount := waitForStatFunction(GetTotalNoteCount, ctx, locale, adminStorage)
	feedback := waitForFeedback(PrepareFeedback, ctx, locale, feedbackStorage)

	wg.Wait()

	var array []string

	if totalUserCount != nil {
		array = append(array, *totalUserCount)
	}
	if totalActiveUserCount != nil {
		array = append(array, *totalActiveUserCount)
	}
	if totalNoteCount != nil {
		array = append(array, *totalNoteCount)
	}

	array = append(array, feedback...)

	return array
}

func PrepareFeedback(ctx *context.Context, locale translator.Locale, feedbackStorage database.FeedbackStorage) []string {
	var array []string

	feedback, err := feedbackStorage.GetFeedbackForLastWeek(ctx)
	if err != nil {
		log.Printf("[Bot] Error getting feedbacks from firestore, %s", err.Error())
		return array
	}

	if len(feedback) == 0 {
		// No feedback
		return array
	}

	array = append(array, "")
	array = append(array, translator.Translate(locale, "feedback_last_week"))
	array = append(array, "")

	for _, f := range feedback {
		var name string

		if f.User.FirstName != nil {
			name += *f.User.FirstName + " "
		}
		if f.User.LastName != nil {
			name += *f.User.LastName
		}
		name = name + ":"
		array = append(array, name+"\n"+f.Feedback.Text+"\n")
	}

	return array
}

func waitForStatFunction(statFunc statFunc, ctx *context.Context, locale translator.Locale, adminStorage database.AdminStorage) *string {
	wg.Add(1)
	ch := make(chan *string)
	go func() {
		defer wg.Done()
		result := statFunc(ctx, locale, adminStorage)
		ch <- result
	}()
	result := <-ch
	return result
}

func waitForFeedback(feedbackFunc feedbackFunc, ctx *context.Context, locale translator.Locale, feedbackStorage database.FeedbackStorage) []string {
	wg.Add(1)
	ch := make(chan []string)
	go func() {
		defer wg.Done()
		result := feedbackFunc(ctx, locale, feedbackStorage)
		ch <- result
	}()
	result := <-ch
	return result
}
