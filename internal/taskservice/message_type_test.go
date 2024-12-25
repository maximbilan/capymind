package taskservice

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
	if got := GetMessage(Morning, 1); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_morning_tuesday"
	if got := GetMessage(Morning, 2); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_morning_wednesday"
	if got := GetMessage(Morning, 3); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_morning_thursday"
	if got := GetMessage(Morning, 4); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_morning_friday"
	if got := GetMessage(Morning, 5); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_morning_saturday"
	if got := GetMessage(Morning, 6); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_morning_sunday"
	if got := GetMessage(Morning, 0); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_evening_monday"
	if got := GetMessage(Evening, 1); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_evening_tuesday"
	if got := GetMessage(Evening, 2); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_evening_wednesday"
	if got := GetMessage(Evening, 3); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_evening_thursday"
	if got := GetMessage(Evening, 4); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_evening_friday"
	if got := GetMessage(Evening, 5); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_evening_saturday"
	if got := GetMessage(Evening, 6); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = "how_are_you_evening_sunday"
	if got := GetMessage(Evening, 0); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}
}

func TestGetMessageInvalidWeekday(t *testing.T) {
	want := ""
	if got := GetMessage(Morning, -1); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}

	want = ""
	if got := GetMessage(Morning, 7); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}
}

func TestGetMessageInvalidMessageType(t *testing.T) {
	want := ""
	if got := GetMessage(Regular, 1); got != want {
		t.Errorf("getMessage() = %v, want %v", got, want)
	}
}
