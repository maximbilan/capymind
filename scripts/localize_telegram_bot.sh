#!/bin/bash

# Check if TELEGRAM_TOKEN and CLOUD_FUNCTION_URL are set
if [ -z "$CAPY_TELEGRAM_BOT_TOKEN" ]; then
  echo "Error: CAPY_TELEGRAM_BOT_TOKEN is not set."
  exit 1
fi

# Set the description

DESCRIPTION_EN="CapyMind is a personal mental health journal powered by AI, designed to help you track your thoughts, emotions, and progress over time. It provides a simple and secure platform for making journal entries, setting reminders, and gaining personalized therapy insights. Leveraging AI to analyze your notes and dreams, CapyMind empowers you to reflect deeply on your mental well-being. With support for multiple languages and time zones, it seamlessly integrates into your daily routine, offering a personalized space for self-reflection and growth."

curl -X POST https://api.telegram.org/bot$CAPY_TELEGRAM_BOT_TOKEN/setMyDescription \
    -H "Content-Type: application/json" \
    -d "{
        \"description\": \"$DESCRIPTION_EN\",
        \"language_code\": \"en\"
    }"

if [ $? -eq 0 ]; then
  echo "EN description has been set successfully."
else
  echo "Failed to set EN description"
fi

SHORT_DESCRIPTION_EN="CapyMind is an AI-powered mental health journal that helps you track emotions, analyze notes or dreams, and reflect on your well-being."

curl -X POST https://api.telegram.org/bot$CAPY_TELEGRAM_BOT_TOKEN/setMyShortDescription \
    -H "Content-Type: application/json" \
    -d "{
        \"short_description\": \"$SHORT_DESCRIPTION_EN\",
        \"language_code\": \"en\"
    }"

if [ $? -eq 0 ]; then
  echo "EN short description has been set successfully."
else
  echo "Failed to set EN short description"
fi

DESCRIPTION_UK="CapyMind — це персональний щоденник для підтримки психічного здоров’я, що працює на основі штучного інтелекту. Він допомагає відстежувати ваші думки, емоції та прогрес з часом. Платформа забезпечує просте та безпечне місце для запису нотаток, встановлення нагадувань і отримання персоналізованих терапевтичних порад. Використовуючи штучний інтелект для аналізу ваших нотаток і снів, CapyMind допомагає вам глибше рефлексувати над своїм психічним станом. Завдяки підтримці кількох мов і часових зон, він легко інтегрується у ваш щоденний графік, надаючи особистий простір для саморозвитку."

curl -X POST https://api.telegram.org/bot$CAPY_TELEGRAM_BOT_TOKEN/setMyDescription \
    -H "Content-Type: application/json" \
    -d "{
        \"description\": \"$DESCRIPTION_UK\",
        \"language_code\": \"uk\"
    }"

if [ $? -eq 0 ]; then
    echo "UK description has been set successfully."
else
    echo "Failed to set UK description"
fi

SHORT_DESCRIPTION_UK="CapyMind — це щоденник для психічного здоров’я з підтримкою ШІ, який допомагає відстежувати емоції, аналізувати нотатки чи сни та рефлексувати над своїм станом."

curl -X POST https://api.telegram.org/bot$CAPY_TELEGRAM_BOT_TOKEN/setMyShortDescription \
    -H "Content-Type: application/json" \
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
  { "command": "settings", "description": "Settings" },
  { "command": "support", "description": "Send feedback" },
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
  { "command": "settings", "description": "Налаштування" },
  { "command": "support", "description": "Надіслати відгук" },
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
