package translator

const translationsJSON = `{
    "en": {
        "welcome": "Welcome to CapyMind 👋 Your personal journal for mental health notes is here to help you on your journey. Reflect on your thoughts and emotions, use reminders to stay on track, and explore therapy insights to deepen your self-awareness.",
        "welcome_onboarding": "Welcome to CapyMind 👋 Before we begin, let’s adjust your settings for a personalized experience 🙂",
        "start_note" : "Share your thoughts and feelings by entering them in the text field and sending them my way. Your personal reflections will be securely saved in your journal 👇",
        "finish_note" : "Your thoughts have been successfully saved. Thank you for trusting CapyMind. Remember, each note is a step forward on your journey to better mental well-being 🙂",
        "your_last_note": "Here’s your most recent note 👇\n\n",
        "no_notes": "You haven’t added any entries yet. Start by sharing your thoughts and feelings with CapyMind.",
        "commands_hint": "Here are the commands you can use to interact with CapyMind 👇\n\n/start Begin using the bot\n/note Make a journal entry\n/last View your most recent entry\n/analysis Receive an analysis of your journal\n/language Set your language preference\n/timezone Set your time zone\n/help Get assistance with using CapyMind\n",
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
        "total_note_count": "The total number of notes is %d"
    },
    "uk": {
        "welcome": "Ласкаво просимо до CapyMind 👋 Ваш особистий журнал для записів про психічне здоров'я тут, щоб допомогти вам на вашому шляху. Рефлексуйте над своїми думками та емоціями, використовуйте нагадування, щоб залишатися на шляху, та досліджуйте інсайти терапії, щоб поглибити свою самосвідомість.",
        "welcome_onboarding": "Ласкаво просимо до CapyMind 👋 Перед початком давайте налаштуємо ваші параметри 🙂",
        "start_note": "Поділіться своїми думками та почуттями, введіть їх у текстове поле та надішліть. Ваші особисті роздуми будуть безпечно збережені у вашому журналі 👇",
        "finish_note": "Ваші думки успішно збережені. Дякуємо вам за довіру CapyMind. Пам'ятайте, кожен запис - це крок вперед на вашому шляху до кращого психічного самопочуття 🙂",
        "your_last_note": "Ось ваш останній запис 👇\n\n",
        "no_notes": "Ви ще не додали жодних записів. Почніть, поділившись своїми думками та почуттями з CapyMind.",
        "commands_hint": "Ось команди, які ви можете використовувати для взаємодії з CapyMind 👇\n\n/start Почати використання бота\n/note Зробити запис у журнал\n/last Переглянути ваш останній запис\n/analysis Отримати аналіз вашого журналу\n/language Встановити мову\n/timezone Встановити ваш часовий пояс\n/help Отримати допомогу з використання CapyMind\n",
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
        "total_note_count": "Загальна кількість записів: %d"
    }
}`

const promptsJSON = `{
    "en": {
        "ai_weekly_analysis_system_message": "You are a skilled therapist at CapyMind, specializing in reviewing and providing feedback on user journals. Your responses should be structured into three distinct parts: 1. Praise & Encouragement: Begin by acknowledging the user’s efforts, offering positive reinforcement for progress made. If progress is minimal, provide motivational support to encourage continued effort (2-3 sentences) 2. Analysis: Analyze the user’s recent journal entries, identifying key patterns or themes in their thoughts, emotions, or behaviors (up to 10 sentences) 3. Recommendations: Finish by offering 3-4 sentences of practical suggestions or steps the user can take to continue their personal growth. (Don't use 1, 2, 3 in the actual response and the names of the parts)",
        "ai_weekly_analysis_user_message": "Below is a list of my recent journal entries. Please provide feedback: ",
        "ai_sleep_analysis_system_message": "You are a sleep therapist at CapyMind, specializing in analyzing sleep patterns and providing recommendations for better sleep quality.",
        "ai_sleep_analysis_user_message": "Below are notes from my last sleep. Please provide feedback: ",
        "ai_quick_analysis_system_message": "You are a therapist at CapyMind, specializing in analyzing user journal entries and providing insights to support their mental well-being.",
        "ai_quick_analysis_user_message": "Below are my last 5 journal entries. Please provide feedback: "
    },
    "uk": {
        "ai_weekly_analysis_system_message": "Ви є кваліфікованим терапевтом в CapyMind, який спеціалізується на перегляді та наданні відгуку аналізуючи записи користувачів. Ваші відповіді повинні бути структуровані на три відмінні частини: 1. Похвала та підтримка: Почніть з визнання зусиль користувача, запропонуйте позитивне підкріплення за досягнуті успіхи. Якщо прогрес мінімальний, надайте мотиваційну підтримку для підтримки подальших зусиль (2-3 речення) 2. Аналіз: Проаналізуйте нещодавні записи користувача, визначивши ключові шаблони або теми у їхніх думках, емоціях або поведінці (до 10 речень) 3. Рекомендації: Завершіть, запропонувавши 3-4 речення практичних порад або кроків, які користувач може зробити для особистого росту. (Не використовуйте 1, 2, 3 у фактичній відповіді та назви частин)",
        "ai_weekly_analysis_user_message": "Нижче наведено список моїх нещодавніх записів у журналі. Будь ласка, надайте відгук. Записи: ",
        "ai_sleep_analysis_system_message": "Ви є терапевтом зі сну в CapyMind, який спеціалізується на аналізі сновидінь та наданні рекомендацій для покращення якості сну.",
        "ai_sleep_analysis_user_message": "Нижче наведено записи з мого останнього сну. Будь ласка, надайте відгук. Записи: ",
        "ai_quick_analysis_system_message": "Ви є терапевтом в CapyMind, який спеціалізується на аналізі записів користувачів у журналі та наданні інсайтів для підтримки їхнього психічного самопочуття.",
        "ai_quick_analysis_user_message": "Нижче наведено мої останні 5 записів у журналі. Будь ласка, надайте відгук. Записи: "
    }
}`

const searchKeywordsJSON = `{
    "en": {
        "dreams": ["dream", "dreams", "night", "sleep", "dreaming", "nightmare"]
    },
    "uk": {
        "dreams": ["сни", "сон", "сновидіння", "кошмари", "кошмар", "cні", "приснилось", "приснилося", "наснилось", "наснилося", "сниться", "снів", "снах", "снилось", "снилося", "сновидінь", "сновидіннях"]
    }
}`
