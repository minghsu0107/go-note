package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// If a value is available on messages
	// then select will take the <-messages case with that value.
	// If not it will immediately take the default case
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Here msg cannot be sent to the messages channel,
	// because the channel has no buffer and there is no receiver.
	// Therefore the default case is selected.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
	/*
	   no message received
	   no message sent
	   no activity
	*/
}
