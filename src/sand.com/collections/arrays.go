package main

import "fmt"

var days = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

func main() {
	fmt.Println(days)

	for _, value := range days {
		fmt.Println("Day using range is " + value)
	}

	for i := 0; i < len(days); i++ {
		fmt.Println("Day using i is " + days[i])
	}
}
