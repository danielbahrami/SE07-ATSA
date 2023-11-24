package env

import (
	"os"
	"strings"
)

var env = map[string]string{}

// Load
// This should be called in the top of the main function
func Load(filename string) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes), "\n")
	for i := range lines {
		if !strings.HasPrefix(lines[i], "#") {
			key, value, found := strings.Cut(lines[i], "=")
			if found {
				env[key] = clean(value)
			}
		}
	}
}

func clean(value string) string {
	value = strings.TrimSpace(value)
	if strings.HasPrefix(value, "\"") {
		return strings.TrimSuffix(strings.TrimPrefix(value, "\""), "\"")
	} else if strings.HasPrefix(value, "'") {
		return strings.TrimSuffix(strings.TrimPrefix(value, "'"), "'")
	}
	return value
}

func Get(key string) string {
	return env[key]
}
