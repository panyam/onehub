package utils

import (
	"os"
)

/**
 * Gets the value of an environment value and if it is missing returns the
 * default value.
 */
func GetEnvOrDefault(envvar string, defaultValue string) string {
	out := os.Getenv(envvar)
	if out == "" {
		out = defaultValue
	}
	return out
}
