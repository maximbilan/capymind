package bot

import "github.com/capymind/internal/firestore"

type Session struct {
	Job  Job
	User firestore.User
}
