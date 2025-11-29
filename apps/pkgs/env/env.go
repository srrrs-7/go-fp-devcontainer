package env

import (
	"fmt"
	"os"
	"strings"
)

// GetString returns the value of the environment variable named by the key.
// If the variable is not present, it returns the defaultValue.
func GetString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetInt returns the value of the environment variable named by the key as an integer.
// If the variable is not present or cannot be parsed, it returns the defaultValue.
func GetInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	var result int
	if _, err := fmt.Sscanf(value, "%d", &result); err != nil {
		return defaultValue
	}
	return result
}

// GetBool returns the value of the environment variable named by the key as a boolean.
// It accepts "true", "1", "yes" as true values (case-insensitive).
// If the variable is not present, it returns the defaultValue.
func GetBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	switch strings.ToLower(value) {
	case "true", "1", "yes":
		return true
	case "false", "0", "no":
		return false
	default:
		return defaultValue
	}
}
