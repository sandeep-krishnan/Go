package main

import (
	"fmt"
	"sand.com/golang-examples/basics"
)

func main() {
	fmt.Println(basics.NAME)
	//not able to find "name" as first char is lower
	//fmt.Println(basics.name)

	//enum
	fmt.Println(basics.THREE)

	i := 1
	i++
	fmt.Println("i is ", i)

	//compilation error, ++ is a statement, not expression
	//fmt.Println("i is " , i++)

	//this is not supported, only i++ will work
	//++i
}
