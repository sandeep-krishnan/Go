package main

import "fmt"

func main() {
	var n  = 5
	result := factorial(n)
	fmt.Println("result is ", result)
}

func factorial(n int) int {
	var factorial int = 1;
	for i := 2; i <= n; i++ {
		//fmt.Println(i)
		factorial *= i
	}

	//fmt.Println("Factorial is ", factorial)
	return factorial
}