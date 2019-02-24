package main

import (
	"fmt"
	"sand.com/concurrency/util"
	"strings"
)

func main() {
	wordMap := make(map[string]int)
	wordCh := make(chan string)

	go func() {
		defer close(wordCh)
		for _, sentence := range util.Sentences {
			words := strings.Split(sentence, " ")
			for _, word := range words {
				wordCh <- word
			}
		}
	}()

	//the wordCh is used directly in range
	//note the return of range is having only 1 param
	for word := range wordCh {
		wordMap[word]++
	}

	fmt.Println(wordMap)
}
