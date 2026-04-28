package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func ConfigPath() (string, error) {
	configHome, err := xdgPath("XDG_CONFIG_HOME", ".config")
	if err != nil {
		return "", err
	}

	return filepath.Join(configHome, "focus-gremlin", "config.json"), nil
}

func DataPath() (string, error) {
	dataHome, err := xdgPath("XDG_DATA_HOME", ".local", "share")
	if err != nil {
		return "", err
	}

	return filepath.Join(dataHome, "focus-gremlin", "history.db"), nil
}

func RuntimeDir() (string, error) {
	runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
	if runtimeDir == "" {
		return "", fmt.Errorf("XDG_RUNTIME_DIR is not set")
	}

	return filepath.Join(runtimeDir, "focus-gremlin"), nil
}

func xdgPath(envKey string, fallbackParts ...string) (string, error) {
	if value := os.Getenv(envKey); value != "" {
		return value, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolve home dir: %w", err)
	}

	parts := append([]string{home}, fallbackParts...)
	return filepath.Join(parts...), nil
}
