package main

import (
	"fmt"
	"math/rand"
	"time"
)

var helloWorldArr = [][]string{
	{"hello world", "english"},
	{"hola mundo", "spanish"},
}

func main() {
	helloWorldRecord := getRandomMessage()
	fmt.Printf(helloWorldRecord[0] + " is in " + helloWorldRecord[1])
	fmt.Println()
}

func getRandomMessage() []string {
	//start - declaring an unused variable and using _ to ignore it
	var unUsed int
	_ = unUsed
	//end - declaring an unused variable and using _ to ignore it

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return helloWorldArr[rnd.Intn(len(helloWorldArr))]
}
