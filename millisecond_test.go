package milli

import (
	"github.com/nowk/assert"
	"testing"
	"time"
)

func TestUnixTimeMillisecond(t *testing.T) {
	d := time.Date(2014, time.August, 15, 4, 31, 12, 500000000, time.UTC)
	m := GetTime(d)
	assert.Equal(t, int64(1408077072500), m)
}

func TestGetMillisecondValue(t *testing.T) {
	d := time.Date(2014, time.August, 15, 4, 31, 12, 500000000, time.UTC)
	m := GetMilliseconds(d)
	assert.Equal(t, 500, m)
}
