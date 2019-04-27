package main

import (
	"fmt"
	"sand.com/concurrency/util"
	"strings"
)

func main() {
	wordMap := make(map[string]int)
	done := make(chan struct{})

	go func() {
		// the close will be triggered after the
		// function is finished and before return
		defer close(done)
		for _, sentence := range util.Sentences {
			words := strings.Split(sentence, " ")
			for _, word := range words {
				wordMap[word]++
			}
		}
	}()

	//waiting for output from channel
	//this will block until the defer is triggered
	<-done

	fmt.Println(wordMap)
}
