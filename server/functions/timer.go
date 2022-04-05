package server

import "time"

func GetUptime(inn time.Time) time.Duration {
	return time.Since(inn)
}
