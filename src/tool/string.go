package tool

import "time"

//GetCurrentTimeStr is time conver to string
func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetCurrentTimeString is time conver to string
func GetCurrentTimeString() string {
	return time.Now().Format("2006-01-02-150405")
}
