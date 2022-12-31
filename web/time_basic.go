package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday()) // Tuesday

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff) // 91709h7m6.443124763s

	p(diff.Hours())       // 91709.11845642354
	p(diff.Minutes())     // 5.502547107385413e+06
	p(diff.Seconds())     // 3.301528264431248e+08
	p(diff.Nanoseconds()) // 330152826443124763

	p(then.Add(diff))
	p(then.Add(-diff)) // 1999-06-02 15:27:52.208262474 +0000 UTC
}
