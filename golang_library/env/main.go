package env

import "os"

func Env(key string, default_value string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return default_value
	}
	return val
}
