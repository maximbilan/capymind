package translator

const translationsJSON = `{
    "en": {
        "welcome": "Welcome to CapyMind 👋 Your personal journal for mental health notes is here to help you on your journey. Reflect on your thoughts and emotions, use reminders to stay on track, and explore therapy insights to deepen your self-awareness.",
        "welcome_onboarding": "Welcome to CapyMind 👋 Before we begin, let’s adjust your settings for a personalized experience 🙂",
        "start_note" : "Share your thoughts and feelings by entering them in the text field and sending them my way. Your personal reflections will be securely saved in your journal 👇",
        "finish_note" : "Your thoughts have been successfully saved. Thank you for trusting CapyMind. Remember, each note is a step forward on your journey to better mental well-being 🙂",
        "your_last_note": "Here’s your most recent note 👇\n\n",
        "no_notes": "You haven’t added any entries yet. Start by sharing your thoughts and feelings with CapyMind.",
        "commands_hint": "Here are the commands you can use to interact with CapyMind 👇\n\n/start Begin using the bot\n/note Make a journal entry\n/last View your most recent entry\n/analysis Receive an analysis of your journal\n/settings Settings\n/language Set your language preference\n/timezone Set your time zone\n/feedback Give feedback \n/help Get assistance with using CapyMind\n",
        "locale_set": "Your language settings have been successfully updated 🌍",
        "language_select": "Select your preferred language 👇",
        "timezone_select": "Select your time zone 👇",
        "timezone_set": "Your time zone has been updated successfully 🤘 You’ll start receiving reminders to log your entries ⏰",
        "how_are_you_morning": "Good morning! How are you feeling today? Any dreams during the night or thoughts on your mind to share?",
        "how_are_you_evening": "Good evening! How was your day? Any reflections or thoughts you'd like to share before you rest?",
        "make_record_to_journal": "Make a record to your journal 💭",
        "make_record_to_journal_short": "Make a record 💭",
        "no_analysis": "Not enough entries have been made to generate an analysis. Begin by sharing your thoughts and feelings with CapyMind.",
        "analysis_waiting": "Your analysis is being generated. Please hold on for a moment 😴",
        "how_to_use": "Getting started 🙋‍♂️",
        "weekly_analysis": "Weekly analysis 🧑‍⚕️\n\n",
        "do_you_want_sleep_analysis": "Would you like to analyze your sleep patterns? 🌙",
        "sleep_analysis": "Sleep analysis 🛌",
        "user_progress_message": "You have made a total of %d entries in your journal.\nKeep up the great work! 🚀",
        "total_user_count": "The total number of users is %d",
        "total_note_count": "The total number of notes is %d",
        "start_feedback": "Your feedback is valuable to us! 🙏 Feel free to share your thoughts down below 👇",
        "finish_feedback": "Thank you for sharing your feedback with CapyMind! 🚀",
        "feedback_last_week": "Feedback from last week 📈",
        "no_feedback": "No feedback has been provided yet",
        "settings_descr": "Please select the setting you would like to change:",
        "language": "Language 🌍",
        "timezone": "Time zone ⏰"
    },
    "uk": {
        "welcome": "Ласкаво просимо до CapyMind 👋 Ваш особистий журнал для записів про психічне здоров'я тут, щоб допомогти вам на вашому шляху. Рефлексуйте над своїми думками та емоціями, використовуйте нагадування, щоб залишатися на шляху, та досліджуйте інсайти терапії, щоб поглибити свою самосвідомість.",
        "welcome_onboarding": "Ласкаво просимо до CapyMind 👋 Перед початком давайте налаштуємо ваші параметри 🙂",
        "start_note": "Поділіться своїми думками та почуттями, введіть їх у текстове поле та надішліть. Ваші особисті роздуми будуть безпечно збережені у вашому журналі 👇",
        "finish_note": "Ваші думки успішно збережені. Дякуємо вам за довіру CapyMind. Пам'ятайте, кожен запис - це крок вперед на вашому шляху до кращого психічного самопочуття 🙂",
        "your_last_note": "Ось ваш останній запис 👇\n\n",
        "no_notes": "Ви ще не додали жодних записів. Почніть, поділившись своїми думками та почуттями з CapyMind.",
        "commands_hint": "Ось команди, які ви можете використовувати для взаємодії з CapyMind 👇\n\n/start Почати використання бота\n/note Зробити запис у журнал\n/last Переглянути ваш останній запис\n/analysis Отримати аналіз вашого журналу\n/settings Налаштування\n/language Встановити мову\n/timezone Встановити ваш часовий пояс\n/feedback Залишити відгук\n/help Отримати допомогу з використання CapyMind\n",
        "locale_set": "Ваші налаштування мови успішно оновлено 🌍",
        "language_select": "Виберіть вашу мову 👇",
        "timezone_select": "Виберіть свій часовий пояс 👇",
        "timezone_set": "Ваш часовий пояс успішно оновлено 🤘 Ви почнете отримувати нагадування про ведення записів ⏰",
        "how_are_you_morning": "Доброго ранку! Як ви себе почуваєте сьогодні? Чи були у вас сни протягом ночі або думки на ранок, якими ви хочете поділитися?",
        "how_are_you_evening": "Доброго вечора! Як минув ваш день? Можливо, у вас є думки або враження, якими ви хотіли б поділитися перед сном?",
        "make_record_to_journal": "Зробити запис у свій журнал 💭",
        "make_record_to_journal_short": "Зробити запис 💭",
        "no_analysis": "Недостатньо записів для генерації аналізу. Почніть ділитись своїми думками та почуттями з CapyMind.",
        "analysis_waiting": "Ваш аналіз генерується. Будь ласка, зачекайте 😴",
        "how_to_use": "Допомога 🙋‍♂️",
        "weekly_analysis": "Аналіз за останній тиждень 🧑‍⚕️\n\n",
        "do_you_want_sleep_analysis": "Бажаєте проаналізувати ваші сновидіння? 🌙",
        "sleep_analysis": "Аналіз сну 🛌",
        "user_progress_message": "Кількість записів у вашому журналі: %d!\nЧудова робота!\nПродовжуйте в тому ж дусі! 🚀",
        "total_user_count": "Загальна кількість користувачів: %d",
        "total_note_count": "Загальна кількість записів: %d",
        "start_feedback": "Ваш відгук для нас дуже важливий! 🙏 Не соромтеся ділитися своїми думками нижче 👇",
        "finish_feedback": "Дякуємо вам за відгук про CapyMind! 🚀",
        "feedback_last_week": "Відгуки з минулого тижня 📈",
        "no_feedback": "Жодного відгуку ще не надано",
        "settings_descr": "Будь ласка, виберіть параметр, який ви хочете змінити:",
        "language": "Мова 🌍",
        "timezone": "Часовий пояс ⏰"
    }
}`
