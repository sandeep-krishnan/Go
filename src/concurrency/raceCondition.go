package main

import (
	"concurrency/util"
	"fmt"
)

func main() {
	starts := []int{1, 5, 10}

	//value of j is changing
	//can cause race conditions
	for _, j := range starts {
		go util.PrintNumbers(1, 10, j)
	}

	fmt.Println("=============")

	//alternative approach
	//anonymous function with j as input
	for _, j := range starts {
		go func(s int) {
			util.PrintNumbers(1, 10, s)
		}(j)
	}

	fmt.Scanln()
}
