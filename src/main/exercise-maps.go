package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordmap := make(map[string]int)
	words := strings.Fields(s)
	for _, w := range words {
		wordmap[w] = wordmap[w] + 1
	}
	return wordmap
}

func main() {
	wc.Test(WordCount)
}
