package config

type Config struct {
	IntervalMinutes      int      `json:"interval_minutes"`
	ActiveStart          string   `json:"active_start"`
	ActiveEnd            string   `json:"active_end"`
	SnoozeMinutes        int      `json:"snooze_minutes"`
	PromptTimeoutSeconds int      `json:"prompt_timeout_seconds"`
	Statuses             []string `json:"statuses"`
	TerminalCommand      []string `json:"terminal_command,omitempty"`
}

func DefaultConfig() Config {
	return Config{
		IntervalMinutes:      60,
		ActiveStart:          "09:00",
		ActiveEnd:            "17:00",
		SnoozeMinutes:        10,
		PromptTimeoutSeconds: 0,
		Statuses: []string{
			"on_track",
			"off_track",
			"deep_in_the_void",
			"snooze",
		},
		TerminalCommand: []string{"x-terminal-emulator", "-e"},
	}
}
