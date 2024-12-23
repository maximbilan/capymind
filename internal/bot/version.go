package bot

import "os"

func handleVersion(session *Session) {
	versionStr := "Version: " + os.Getenv("APP_VERSION")
	setOutputText(versionStr, session)
}
