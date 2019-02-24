package main

import (
	"fmt"
	"strings"
)
import "sand.com/concurrency/util"

func main() {
	ch := make(chan bool)
	wordMap := make(map[string]int)

	go func() {
		for _, sentence := range util.Sentences {
			fmt.Println(sentence)
			words := strings.Split(sentence, " ")
			for _, word := range words {
				wordMap[word]++
			}
		}
		ch <- true
	}()

	if <-ch {
		fmt.Println(wordMap)
	}

}
