package main

import (
	"fmt"
	"sand.com/concurrency/util"
	"strings"
)

func main() {

	wordCh := producer(util.Sentences)
	wordMap := consumer(wordCh)

	//receive only channel, can't add
	//the following will throw compile error
	//wordCh <- "abc"

	fmt.Println(wordMap)
}

func producer(sentences []string) <-chan string {
	wordCh := make(chan string)

	go func() {
		defer close(wordCh)
		for _, sentence := range sentences {
			words := strings.Split(sentence, " ")
			for _, word := range words {
				wordCh <- strings.ToLower(word)
			}
		}
	}()

	return wordCh
}

func consumer(wordCh <-chan string) map[string]int {
	wordMap := make(map[string]int)

	for word := range wordCh {
		wordMap[word]++
	}

	return wordMap
}
