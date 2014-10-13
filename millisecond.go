package milli

import (
	"errors"
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

const (
	boundl  = int64(time.Second * time.Microsecond)
	boundh  = boundl * 10
	micro64 = int64(micro)
)

// NewTime converts milliseconds to time, returns as UTC
// Must be milliseconds
func NewTime(m int64) (t time.Time, err error) {
	if m < boundl || m >= boundh {
		return t, errors.New("error: number was not in milliseconds")
	}

	s := m / micro64
	n := (m % micro64) * milli

	t = time.Unix(s, n)
	return t.UTC(), nil
}
