package go_helper

import (
	"os"
)

func GetENV(key string, default_value string) (ENV_VALUE string) {
	ENV_VALUE = os.Getenv(key)
	if ENV_VALUE == "" {
		ENV_VALUE = default_value
	}
	return
}
