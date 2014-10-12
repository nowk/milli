# milli

[![Build Status](https://travis-ci.org/nowk/milli.svg?branch=master)](https://travis-ci.org/nowk/milli)
[![GoDoc](https://godoc.org/github.com/nowk/milli?status.svg)](http://godoc.org/github.com/nowk/milli)

Sometimes you just need milliseconds.

## Examples

    d := time.Date(2014, time.September, 11, 4, 31, 12, 500000000, time.UTC)
    m := milli.GetTime(d)

    // m => 1410409872500

If you just want the millisecond value

    m := milli.GetMilliseconds(d)

    // m => 500


## License

MIT