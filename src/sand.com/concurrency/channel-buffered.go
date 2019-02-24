package main

import "fmt"

func main() {
	//bi-directional buffered channel of size 3
	//FIFO
	ch := make(chan int, 3)
	ch <- 11
	ch <- 12

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//this will block as channel is empty
	fmt.Println(<-ch)
}
