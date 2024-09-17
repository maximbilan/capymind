# capymind

CapyMind is a personal mental health journal designed to help you track your thoughts, emotions, and progress over time. By offering a simple and secure platform to make journal entries, set reminders, and receive personalized therapy insights, CapyMind empowers you to reflect on your mental well-being. With support for multiple languages and time zones, it fits seamlessly into your daily routine, providing a personalized space for self-reflection and growth.

# Prerequisites

1. Install Go
2. VSCode or other editors
3. ngrok or other tunnels

# How to run locally

1. `go build`
2. `go run cmd/main.go` to run the server
3. Start `ngrok` or other tunnel
4. Set up `CAPY_CLOUD_FUNCTION_URL` as a env variable with the tunnel url
5. `chmod +x ./scripts/setup_telegram.bot.sh`
6. `./scripts/setup_telegram.bot.sh` (Set up a Telegram token)

# How to set up GCloud service account

1. `chmod +x ./scripts/setup_gcloud_access.sh`
2. `./scripts/setup_gcloud_access.sh`

# How to deploy cloud functions

1. `chmod +x ./scripts/deploy_functions.sh`
2. `./scripts/deploy_functions.sh`

# How to localize Telegram Bot description (short and full)

1. `chmod +x ./scripts/localize_telegram_bot.sh`
2. `./scripts/localize_telegram_bot.sh`

# How to schedule (cron) jobs to send Morning and Evening messages + Weekly Analysis

1. `chmod +x ./scripts/schedule_jobs.sh`
2. `./scripts/schedule_jobs.sh`
