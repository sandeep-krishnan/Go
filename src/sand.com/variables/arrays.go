package main

import "fmt"

func main() {
	var x [5]int
	//will not modify x as a copy is made
	initialize(x)
	fmt.Println(x)
}

// the array is copied as part of this method
func initialize(numbers [5]int) {
	for i, _ := range numbers {
		numbers[i] = 10
	}
}

func initializeUsingPtr(numbers *[]int) {
	for i, _ := range numbers {
		numbers[i] = 10
	}
}
