package scheduler

import "testing"

func TestMessageType(t *testing.T) {
	want := "morning"
	if got := Morning; string(got) != want {
		t.Errorf("Morning = %v, want %v", got, want)
	}

	want = "evening"
	if got := Evening; string(got) != want {
		t.Errorf("Evening = %v, want %v", got, want)
	}

	want = "weekly_analysis"
	if got := WeeklyAnalysis; string(got) != want {
		t.Errorf("WeeklyAnalysis = %v, want %v", got, want)
	}

	want = "user_stats"
	if got := UserStats; string(got) != want {
		t.Errorf("UserStats = %v, want %v", got, want)
	}
}

func TestGetMessage(t *testing.T) {
	want := "how_are_you_morning_monday"
	if got := getMessage(Morning, 1); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_morning_tuesday"
	if got := getMessage(Morning, 2); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_morning_wednesday"
	if got := getMessage(Morning, 3); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_morning_thursday"
	if got := getMessage(Morning, 4); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_morning_friday"
	if got := getMessage(Morning, 5); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_morning_saturday"
	if got := getMessage(Morning, 6); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_morning_sunday"
	if got := getMessage(Morning, 0); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_evening_monday"
	if got := getMessage(Evening, 1); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_evening_tuesday"
	if got := getMessage(Evening, 2); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_evening_wednesday"
	if got := getMessage(Evening, 3); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_evening_thursday"
	if got := getMessage(Evening, 4); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_evening_friday"
	if got := getMessage(Evening, 5); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_evening_saturday"
	if got := getMessage(Evening, 6); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_evening_sunday"
	if got := getMessage(Evening, 0); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}
}

func TestGetMessageInvalidWeekday(t *testing.T) {
	want := ""
	if got := getMessage(Morning, -1); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = ""
	if got := getMessage(Morning, 7); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}
}

func TestGetMessageInvalidMessageType(t *testing.T) {
	want := ""
	if got := getMessage(Regular, 1); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}
}
