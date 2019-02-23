package main

import (
	"fmt"
	"strings"
)

func main() {

	var days = [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	//the input is a slice of days => [:], not directly the array
	printArray(days[:])

	//the input is [7] of string, not []string
	//the function makes a copy and modifies it
	changeToLowerCasePassByValue(days)

	for i := 0; i < len(days); i++ {
		fmt.Println("Day using i is " + days[i])
	}

	changeToLowerCasePassByReference(days[:])

	printArray(days[:])
}

func printArray(days []string) {
	for _, value := range days {
		fmt.Println("Day using range is " + value)
	}
}

/*
In this case the original value of the days array will not change
This is because go makes a copy of the input and then modifies that
*/
func changeToLowerCasePassByValue(days [7]string) {
	fmt.Println("Changing to lower case - Pass by Copy")
	for i, _ := range days {
		days[i] = strings.ToLower(days[i])
	}
}

func changeToLowerCasePassByReference(days []string) {
	fmt.Println("Changing to lower case - Pass by Value")
	for i, _ := range days {
		days[i] = strings.ToLower(days[i])
	}
}
