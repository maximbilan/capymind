package bot

type TextButton struct {
	TextID   string
	Callback string
}

type JobResult struct {
	TextID  string
	Buttons []TextButton
}
