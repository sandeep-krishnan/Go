package main

import "fmt"
import "concurrency/util"

func main() {
	go util.PrintNumbers(1, 10, 2)
	go util.PrintNumbers(2, 10, 2)

	//this is just to block the main thread till the execution completes
	fmt.Scanln()
}
