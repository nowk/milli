package milli

import (
	"github.com/nowk/assert"
	"testing"
	"time"
)

func TestUnixTimeMillisecond(t *testing.T) {
	d := time.Date(2014, time.August, 15, 4, 31, 12, 500000000, time.UTC)
	m := ParseTime(d)
	assert.Equal(t, int64(1408077072500), m)
}

func TestGetMillisecondValue(t *testing.T) {
	d := time.Date(2014, time.August, 15, 4, 31, 12, 500000000, time.UTC)
	m := GetMilliseconds(d)
	assert.Equal(t, 500, m)
}

func TestMillisecondsToTime(t *testing.T) {
	d := time.Date(2014, time.August, 15, 4, 31, 12, 500000000, time.UTC)
	u, err := NewTime(int64(1408077072500))
	assert.Nil(t, err)
	assert.Equal(t, u, d)
}

func TestMillisecondsToTimeWhenNotMillisecondLength(t *testing.T) {
	_, err := NewTime(int64(14080770725000))
	assert.Equal(t, "error: number was not in milliseconds", err.Error())

	_, err = NewTime(int64(14080770725))
	assert.Equal(t, "error: number was not in milliseconds", err.Error())
}
