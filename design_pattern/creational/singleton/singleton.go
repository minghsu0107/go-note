package singleton

import "sync"

type Singleton struct{}

var singleton *Singleton
var once sync.Once

// lazy mode (initialize when used)
// for thread safety, use double checking + lock
func GetInstance() *Singleton {
	// only execute once
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}

/*
implementation of sync.once:

package sync

import (
	"sync/atomic"
)

type Once struct {
	m    Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
*/
