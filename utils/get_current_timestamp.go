package utils

import "time"

func GetCurrentTimestampString() string {
	return time.Now().UTC().Format(time.RFC3339)
}
