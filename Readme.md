# milli

[![Build Status](https://travis-ci.org/nowk/milli.svg?branch=master)](https://travis-ci.org/nowk/milli)
[![GoDoc](https://godoc.org/github.com/nowk/milli?status.svg)](http://godoc.org/github.com/nowk/milli)

Sometimes you just need milliseconds.

## Examples

    d := time.Date(2014, time.September, 11, 4, 31, 12, 500000000, time.UTC)
    m := milli.ParseTime(d)

    // m => 1410409872500

If you just want the millisecond value

    m := milli.GetMilliseconds(d)

    // m => 500

To convert millisecond to time

    t, err := milli.NewTime(1410409872500)
    if err != nil {
      // returns if value is not in millisecond
    }
    
    t.String() // => 2014-08-15 04:31:12.5 +0000 UTC 


## License

MIT