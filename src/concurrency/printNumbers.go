package main

import "fmt"

func main() {
	go printNumbers(1, 10, 2)
	go printNumbers(2, 10, 2)

	//this is just to block the main thread till the execution completes
	fmt.Scanln()
}

func printNumbers(start int, stop int, delta int) {
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}
