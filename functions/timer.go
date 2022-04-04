package funcs

import "time"

var Timer time.Time

func TimerStart(inn time.Time) {
	inn = time.Now()
}

/* func getUptime(inn time.Time) time.Duration {
	return time.Since(inn)
}
*/
