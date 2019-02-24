package main

import "fmt"

func main() {
	ch := make(chan int, 4)

	fmt.Println("Length of channel before adding ", len(ch))
	fmt.Println("Capacity of channel before adding ", cap(ch))
	addNumbersToChannel(ch)

	fmt.Println("Length of channel after adding ", len(ch))

	//close the channel
	close(ch)

	//this will error out as the channel is closed
	ch <- 13

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func addNumbersToChannel(c chan<- int) {
	c <- 1
	c <- 2

	//compile error - receive from send only channel
	//fmt.Println(<- c)
}
