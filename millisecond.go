package milli

import (
	"time"
)

const (
	milli int64 = int64(time.Millisecond)
	micro int   = int(time.Microsecond)
)

// GetTime returns the Time in milliseconds, akin to javascript Date's getTime()
func GetTime(t time.Time) int64 {
	n := t.UnixNano()
	return n / milli
}

// GetMilliseconds returns the millisecond portion of the time
func GetMilliseconds(t time.Time) (n int) {
	m := GetTime(t)
	return int(m) % micro
}
