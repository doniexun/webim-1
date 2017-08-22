package util

import "time"

// GetCurrentTimestamp return current utc timestamp
func GetCurrentTimestamp() int64 {
	return time.Now().UTC().Unix()
}
