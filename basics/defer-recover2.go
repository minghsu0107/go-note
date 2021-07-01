package main

import "log"
import "time"

func shouldNotExit() {
	for {
		// Simulate a workload.
		time.Sleep(time.Second)

		// Simulate an unexpected panic.
		if time.Now().UnixNano()&0x3 == 0 {
			panic("unexpected situation")
		}
	}
}

func NeverExit(name string, f func()) {
	defer func() {
		if v := recover(); v != nil {
			// A panic is detected.
			log.Println(name, "is crashed. Restart it now.")
			go NeverExit(name, f) // restart
		}
	}()
	f()
}

func main() {
	log.SetFlags(0)
	go NeverExit("job#A", shouldNotExit)
	go NeverExit("job#B", shouldNotExit)
	select {} // block here for ever
}
