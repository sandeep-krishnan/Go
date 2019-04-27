package main

import "fmt"

func main() {
	//bi-directional unbuffered
	ch := make(chan int)

	sendNonBlockingMessage(ch)
	fmt.Println(<-ch)

	//this will throw panic
	sendBlockingMessage(ch)

	fmt.Println(<-ch)
}

func sendBlockingMessage(ch chan int) {
	//fatal error: all goroutines are asleep - deadlock!
	ch <- 12
}

func sendNonBlockingMessage(ch chan int) {
	//fatal error: all goroutines are asleep - deadlock!
	go func() { ch <- 12 }()
}
