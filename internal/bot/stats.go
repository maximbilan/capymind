package bot

import "sync"

type statFunc func(session *Session) *string
type feedbackFunc func(session *Session) []string

var wg sync.WaitGroup

func handleStats(session *Session) {
	totalUserCount := waitForStatFunction(getTotalUserCount, session)
	totalActiveUserCount := waitForStatFunction(getTotalActiveUserCount, session)
	totalNoteCount := waitForStatFunction(getTotalNoteCount, session)
	feedback := waitForFeedback(prepareFeedback, session)

	wg.Wait()

	if totalUserCount != nil {
		setOutputText(*totalUserCount, session)
	}
	if totalActiveUserCount != nil {
		setOutputText(*totalActiveUserCount, session)
	}
	if totalNoteCount != nil {
		setOutputText(*totalNoteCount, session)
	}

	for _, item := range feedback {
		setOutputText(item, session)
	}
}

func waitForStatFunction(statFunc statFunc, session *Session) *string {
	wg.Add(1)
	ch := make(chan *string)
	go func() {
		defer wg.Done()
		result := statFunc(session)
		ch <- result
	}()
	result := <-ch
	return result
}

func waitForFeedback(feedbackFunc feedbackFunc, session *Session) []string {
	wg.Add(1)
	ch := make(chan []string)
	go func() {
		defer wg.Done()
		result := feedbackFunc(session)
		ch <- result
	}()
	result := <-ch
	return result
}
