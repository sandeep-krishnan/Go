package main

import "fmt"

func main() {
	fmt.Println(reverse2("hello"))
}

func reverse(input string) string {
	chars := []rune(input)

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func reverse2(input string) string {
	var output string

	for i := len(input) - 1; i >= 0; i-- {
		output += string(input[i])
	}
	return output
}
