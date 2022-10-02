package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) (m map[string]int) {
	m = make(map[string]int)

	for _, word := range strings.Fields(s) {
		if _, ok := m[word]; ok {
			m[word] = m[word] + 1
		} else {
			m[word] = 1
		}
	}
	return
}

func main() {
	wc.Test(WordCount)
}
