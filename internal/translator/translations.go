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
        "ai_analysis_prompt": "You’re a professional therapist at CapyMind. You have received the following entries. (A brief summary, with entries sorted as follows: the most recent ones at the top of the list.) What is your professional opinion? Entries: ",
        "how_to_use": "Getting started 🙋‍♂️",
        "weekly_analysis": "Analysis for the past week 🧑‍⚕️\n\n"
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
        "how_are_you_evening": "Доброго вечора! Як пройшов ваш день? Чи є відгуки або думки, якими ви хочете поділитися перед відпочинком?",
        "make_record_to_journal": "Зробити запис у свій журнал 💭",
        "make_record_to_journal_short": "Зробити запис 💭",
        "no_analysis": "Недостатньо записів для генерації аналізу. Почніть ділитись своїми думками та почуттями з CapyMind.",
        "analysis_waiting": "Ваш аналіз генерується. Будь ласка, зачекайте 😴",
        "ai_analysis_prompt": "Ви професійний терапевт в Capymind. Ви отримуєте наступні записи. (Короткий зміст, записи відсортовані наступним чином: останні записи спочатку списку) Яка ваша професійна думка? Записи: ",
        "how_to_use": "Допомога 🙋‍♂️",
        "weekly_analysis": "Аналіз за останній тиждень 🧑‍⚕️\n\n"
    }
}`

const promptsJSON = `{
    "en": {
        "ai_analysis_system_message": "You are a skilled therapist at CapyMind, specializing in reviewing and providing feedback on user journals. Your responses should be structured into three distinct parts: 1. Praise & Encouragement: Begin by acknowledging the user’s efforts, offering positive reinforcement for progress made. If progress is minimal, provide motivational support to encourage continued effort (2-3 sentences) 2. Analysis: Analyze the user’s recent journal entries, identifying key patterns or themes in their thoughts, emotions, or behaviors (up to 10 sentences) 3. Recommendations: Finish by offering 3-4 sentences of practical suggestions or steps the user can take to continue their personal growth.",
        "ai_analysis_user_message": "Below is a list of my recent journal entries. Please provide feedback: "
    },
    "uk": {
        "ai_analysis_system_message": "Ви є кваліфікованим терапевтом в CapyMind, який спеціалізується на перегляді та наданні відгуку аналізуючи записи користувачів. Ваші відповіді повинні бути структуровані на три відмінні частини: 1. Похвала та підтримка: Почніть з визнання зусиль користувача, запропонуйте позитивне підкріплення за досягнуті успіхи. Якщо прогрес мінімальний, надайте мотиваційну підтримку для підтримки подальших зусиль (2-3 речення) 2. Аналіз: Проаналізуйте нещодавні записи користувача, визначивши ключові шаблони або теми у їхніх думках, емоціях або поведінці (до 10 речень) 3. Рекомендації: Завершіть, запропонувавши 3-4 речення практичних порад або кроків, які користувач може зробити для особистого росту.",
        "ai_analysis_user_message": "Нижче наведено список моїх нещодавніх записів у журналі. Будь ласка, надайте відгук. Записи: "
    }
}`
