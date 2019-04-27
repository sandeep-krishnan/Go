package main

import (
	"fmt"
)

const (
	HELLO, WORLD string = "HELLO", "WORLD"
	SPACE               = " "
)

var (
	helloWorld = HELLO + SPACE + WORLD
)

func main() {
	helloWorldStr := helloWorld
	fmt.Println(helloWorldStr)

	fmt.Println(reverse("hello"))
}
