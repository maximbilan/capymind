package bot

import "github.com/capymind/third_party/firestore"

// Use Firestore for the database
var db firestore.Firestore

// Use the following storages for the database
var adminStorage firestore.AdminStorage
var userStorage firestore.UserStorage
var feedbackStorage firestore.FeedbackStorage
var noteStorage firestore.NoteStorage
