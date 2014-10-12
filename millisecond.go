package milli

import (
	"time"
)

const (
	milli int64 = int64(time.Millisecond)
	micro int   = int(time.Microsecond)
)

// ParseTime returns the Time in milliseconds
func ParseTime(t time.Time) int64 {
	n := t.UnixNano()
	return n / milli
}

// GetMilliseconds returns the millisecond portion of the time
func GetMilliseconds(t time.Time) (n int) {
	m := ParseTime(t)
	return int(m) % micro
}
