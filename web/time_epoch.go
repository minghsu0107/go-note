package main

import (
	"fmt"
	"time"
)

// Use time.Now with Unix or UnixNano to get elapsed time
// since the Unix epoch in seconds or nanoseconds, respectively

func main() {

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)   // 1588643297
	fmt.Println(millis) // 1588643297631
	fmt.Println(nanos)  // 1588643297631362000

	fmt.Println(time.Unix(secs, 0))                // 2020-05-05 09:48:17 +0800 CST
	fmt.Println(time.Unix(0, nanos))               // 2020-05-05 09:48:17.631362 +0800 CST
	fmt.Println(time.Unix(secs, nanos-(secs*1e9))) // 2020-05-05 09:48:17.631362 +0800 CST
}
