package main

import (
	"fmt"
	"sand.com/concurrency/util"
	"strings"
)

func main() {
	wordMap := make(map[string]int)

	wordCh := getWords(util.Sentences)

	//receive only channel, can't add
	//the following will throw compile error
	//wordCh <- "abc"

	for word := range wordCh {
		wordMap[word]++
	}

	fmt.Println(wordMap)
}

func getWords(sentences []string) <-chan string {
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
