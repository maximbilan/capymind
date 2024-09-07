#!/bin/bash

# Check if TELEGRAM_TOKEN and CLOUD_FUNCTION_URL are set
if [ -z "$CAPY_TELEGRAM_BOT_TOKEN" ]; then
  echo "Error: CAPY_TELEGRAM_BOT_TOKEN is not set."
  exit 1
fi

# Set the description

DESCRIPTION_EN="CapyMind is a personal mental health journal designed to help you track your thoughts, emotions, and progress over time. By offering a simple and secure platform to make journal entries, set reminders, and receive personalized therapy insights, CapyMind empowers you to reflect on your mental well-being. With support for multiple languages and time zones, it fits seamlessly into your daily routine, providing a personalized space for self-reflection and growth."

curl -X POST https://api.telegram.org/bot$CAPY_TELEGRAM_BOT_TOKEN/setMyDescription \
    -d "{
        \"description\": \"$DESCRIPTION_EN\",
        \"language_code\": \"en\"
    }"

if [ $? -eq 0 ]; then
  echo "EN description has been set successfully."
else
  echo "Failed to set EN description"
fi

SHORT_DESCRIPTION_EN="CapyMind is a personal mental health journal designed to help you track your thoughts, emotions, and progress over time."

curl -X POST https://api.telegram.org/bot$CAPY_TELEGRAM_BOT_TOKEN/setMyShortDescription \
    -d "{
        \"short_description\": \"$SHORT_DESCRIPTION_EN\",
        \"language_code\": \"en\"
    }"

if [ $? -eq 0 ]; then
  echo "EN short description has been set successfully."
else
  echo "Failed to set EN short description"
fi

DESCRIPTION_UK="CapyMind - це особистий щоденник здоров'я розуму, призначений для відстеження ваших думок, емоцій та прогресу з часом. Надаючи просту та безпечну платформу для ведення щоденника, встановлення нагадувань та отримання персоналізованих висновків про терапію, CapyMind допомагає вам рефлексувати над своїм психічним самопочуттям. З підтримкою кількох мов та часових поясів він легко вписується в вашу щоденну рутину, надаючи персоналізований простір для саморефлексії та зростання."

curl -X POST https://api.telegram.org/bot$CAPY_TELEGRAM_BOT_TOKEN/setMyDescription \
    -d "{
        \"description\": \"$DESCRIPTION_UK\",
        \"language_code\": \"uk\"
    }"

if [ $? -eq 0 ]; then
    echo "UK description has been set successfully."
else
    echo "Failed to set UK description"
fi

SHORT_DESCRIPTION_UK="CapyMind - це особистий щоденник здоров'я розуму, призначений для відстеження ваших думок, емоцій та прогресу з часом."

curl -X POST https://api.telegram.org/bot$CAPY_TELEGRAM_BOT_TOKEN/setMyShortDescription \
    -d "{
        \"short_description\": \"$SHORT_DESCRIPTION_UK\",
        \"language_code\": \"uk\"
    }"

if [ $? -eq 0 ]; then
    echo "UK short description has been set successfully."
else
    echo "Failed to set UK short description"
fi

# Set the commands

COMMANDS_EN='[
  { "command": "start", "description": "Start the bot" },
  { "command": "note", "description": "Make a note" },
  { "command": "last", "description": "Get last note" },
  { "command": "analysis", "description": "Get an analysis" },
  { "command": "language", "description": "Set a language" },
  { "command": "timezone", "description": "Set a timezone" },
  { "command": "help", "description": "Help" }
]'

curl -X POST https://api.telegram.org/bot$CAPY_TELEGRAM_BOT_TOKEN/setMyCommands \
    -H "Content-Type: application/json" \
    -d "{
        \"commands\": $COMMANDS_EN,
        \"language_code\": \"en\"
    }"

if [ $? -eq 0 ]; then
    echo "EN commands have been set successfully."
else
    echo "Failed to set EN commands"
fi

COMMANDS_UK='[
  { "command": "start", "description": "Запуск бота" },
  { "command": "note", "description": "Зробити запис" },
  { "command": "last", "description": "Останній запис" },
  { "command": "analysis", "description": "Зробити аналіз" },
  { "command": "language", "description": "Налаштувати мову" },
  { "command": "timezone", "description": "Налаштувати часовий пояс" },
  { "command": "help", "description": "Допомога" }
]'

if [ $? -eq 0 ]; then
    echo "UK commands have been set successfully."
else
    echo "Failed to set UK commands"
fi

curl -X POST https://api.telegram.org/bot$CAPY_TELEGRAM_BOT_TOKEN/setMyCommands \
    -H "Content-Type: application/json" \
    -d "{
        \"commands\": $COMMANDS_UK,
        \"language_code\": \"uk\",
    }"
