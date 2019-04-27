package main

import "fmt"

func main() {
	var x [5]int
	//will not modify x as a copy is made
	initialize(x)
	fmt.Println(x)

	initializeUsingPtr(&x)
	fmt.Println(x)

	initializeUsingSlice(x[:])
	fmt.Println(x)
}

// the array is copied as part of this method
func initialize(numbers [5]int) {
	fmt.Println("initializing to 5")
	for i := range numbers {
		numbers[i] = 5
	}
}

func initializeUsingPtr(numbers *[5]int) {
	fmt.Println("initializing using pointers to 10")
	for i := 0; i < len(numbers); i++ {
		numbers[i] = 10
	}
}

func initializeUsingSlice(numbers []int) {
	fmt.Println("initializing using slices to 20")
	for i := 0; i < len(numbers); i++ {
		numbers[i] = 20
	}
}
