package main

import (
	"sync"
)

// we can typedef an interface type while preserving its methods
// but typedef a non-interface type will not inherit the original methods
// type myMutex sync.Mutex

type myLocker struct {
	sync.Mutex
}

func main() {
	var locker myLocker
	locker.Lock()
	locker.Unlock()
}
