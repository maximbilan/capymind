package app

import (
	"github.com/capymind/third_party/firestore"
	"github.com/capymind/third_party/telegram"
)

var bot telegram.Telegram

// Use Firestore for the database
var db firestore.Firestore

// Use the following storages for the database
var userStorage firestore.UserStorage
var feedbackStorage firestore.FeedbackStorage
var noteStorage firestore.NoteStorage
