//coverage:ignore file

package scheduler

import (
	"github.com/capymind/third_party/firestore"
	"github.com/capymind/third_party/googletasks"
	"github.com/capymind/third_party/openai"
	"github.com/capymind/third_party/telegram"
)

var bot telegram.Telegram
var aiService openai.OpenAI

// Use Firestore for the database
var db firestore.Firestore
var userStorage firestore.UserStorage
var settingsStorage firestore.SettingsStorage
var noteStorage firestore.NoteStorage
var adminStorage firestore.AdminStorage
var feedbackStorage firestore.FeedbackStorage

var tasks googletasks.GoogleTasks
