package helpers

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/capymind/internal/translator"
	"github.com/capymind/third_party/firestore"
)

type statFunc func(ctx *context.Context, locale translator.Locale) *string
type feedbackFunc func(ctx *context.Context, locale translator.Locale) []string

var wg sync.WaitGroup

var adminStorage firestore.AdminStorage
var feedbackStorage firestore.FeedbackStorage

func GetTotalUserCount(ctx *context.Context, locale translator.Locale) *string {
	count, err := adminStorage.GetTotalUserCount(ctx)
	if err != nil {
		log.Printf("[Admin] Error during fetching total user count: %v", err)
		return nil
	}
	message := fmt.Sprintf(translator.Translate(locale, "total_user_count"), count)
	return &message
}

func GetTotalActiveUserCount(ctx *context.Context, locale translator.Locale) *string {
	count, err := adminStorage.GetActiveUserCount(ctx)
	if err != nil {
		log.Printf("[Admin] Error during fetching active user count: %v", err)
		return nil
	}
	message := fmt.Sprintf(translator.Translate(locale, "total_active_user_count"), count)
	return &message
}

func GetTotalNoteCount(ctx *context.Context, locale translator.Locale) *string {
	count, err := adminStorage.GetTotalNoteCount(ctx)
	if err != nil {
		log.Printf("[Admin] Error during fetching total note count: %v", err)
		return nil
	}
	message := fmt.Sprintf(translator.Translate(locale, "total_note_count"), count)
	return &message
}

func GetStats(ctx *context.Context, locale translator.Locale) []string {
	totalUserCount := waitForStatFunction(GetTotalUserCount, ctx, locale)
	totalActiveUserCount := waitForStatFunction(GetTotalActiveUserCount, ctx, locale)
	totalNoteCount := waitForStatFunction(GetTotalNoteCount, ctx, locale)
	feedback := waitForFeedback(PrepareFeedback, ctx, locale)

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

func PrepareFeedback(ctx *context.Context, locale translator.Locale) []string {
	var array []string
	array = append(array, "feedback_last_week")
	array = append(array, "\n\n")

	feedback, err := feedbackStorage.GetFeedbackForLastWeek(ctx)
	if err != nil {
		log.Printf("[Bot] Error getting feedbacks from firestore, %s", err.Error())
	}

	if len(feedback) == 0 {
		array = append(array, translator.Translate(locale, "no_feedback"))
		return array
	}

	for _, f := range feedback {
		array = append(array, *f.User.FirstName+" "+*f.User.LastName+":"+"\n"+f.Feedback.Text+"\n\n")
	}

	return array
}

func waitForStatFunction(statFunc statFunc, ctx *context.Context, locale translator.Locale) *string {
	wg.Add(1)
	ch := make(chan *string)
	go func() {
		defer wg.Done()
		result := statFunc(ctx, locale)
		ch <- result
	}()
	result := <-ch
	return result
}

func waitForFeedback(feedbackFunc feedbackFunc, ctx *context.Context, locale translator.Locale) []string {
	wg.Add(1)
	ch := make(chan []string)
	go func() {
		defer wg.Done()
		result := feedbackFunc(ctx, locale)
		ch <- result
	}()
	result := <-ch
	return result
}
