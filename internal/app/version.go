//coverage:ignore file

package app

import "os"

func handleVersion(session *Session) {
	setOutputText(os.Getenv("APP_VERSION"), session)
}
