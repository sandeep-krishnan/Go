package main

import "fmt"

func main() {
	var n  = 5
	factorial(n)
}

func factorial(n int)  {
	var factorial int = 1;
	for i := 2; i <= n; i++ {
		//fmt.Println(i)
		factorial *= i
	}

	fmt.Println("Factorial is ", factorial)
}