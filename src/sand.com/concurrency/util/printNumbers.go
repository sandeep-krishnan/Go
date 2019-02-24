package util

import "fmt"

func PrintNumbers(start int, stop int, delta int) {
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}
