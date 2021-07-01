package main

// logger writes to standard error and prints the date and time of each logged message
import (
	"errors"
	"log"
	"net"
)

func ClientHandler(c net.Conn) {
	defer func() {
		if v := recover(); v != nil {
			// v is the error msg
			log.Println("capture a panic:", v)
			log.Println("avoid crashing the program")
		}
		c.Close()
	}()
	panic(errors.New("just a demo.")) // a demo-purpose panic
}

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		// Handle each client connection
		// in a new goroutine.
		go ClientHandler(conn)
	}
}

/*
// we can also use panic-recover to reduce error checks:
func doSomething() (err error) {
	defer func() {
		err = recover() // can change the return value
	}()

	doStep1()
	doStep2()
	doStep3()
	doStep4()
	doStep5()

	return
}

// In reality, the prototypes of the doStepN functions
// might be different. For each of them,
// * panic with nil for success and no needs to continue.
// * panic with error for failure and no needs to contine.
// * not panic for continuing.
func doStepN() {
	...
	if err != nil {
		panic(err)
	}
	...
	if done {
		panic(nil)
	}
}
*/
