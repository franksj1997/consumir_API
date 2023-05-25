package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	Token = getenv("TOKENAPI", "")
)

func getenv(key, defaultValue string) string {
	value, defined := os.LookupEnv(key)
	if !defined {
		return defaultValue
	}
	return value
}
