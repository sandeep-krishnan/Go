package main

import "fmt"

func main() {
	i := 10
	j := 20

	fmt.Println("i is ", i)

	p := &i
	*p = *p * 10

	fmt.Println("i is ", i, " j is ", j)
	swap(i, j)
	fmt.Println("i is ", i, " j is ", j)
	swapUsingPointers(&i, &j)
	fmt.Println("i is ", i, " j is ", j)
}

func swap(x int, y int) {
	t := x
	x = y
	y = t

	fmt.Println("x : ", x, " y : ", y)
}

func swapUsingPointers(x *int, y *int) {
	t := *x
	*x = *y
	*y = t

	fmt.Println("x : ", *x, " y : ", *y)
}
