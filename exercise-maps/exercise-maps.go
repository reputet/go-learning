package main

import (
    "golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var words []string
	var wc map[string]int

	words = strings.Fields(s)
	wc = make(map[string]int)

	for _, v := range words {
		wc[v]++
	}

	return wc
}

func main() {
	wc.Test(WordCount)
}

