package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	m := make(map[string]int)
	sFields := strings.Fields(s)

	sLength := len(sFields)
	for i := 0; i < sLength; i++ {
		token := sFields[i]
		_, ok := m[token]
		if ok == true {
			m[token] += 1
		} else {
			m[token] = 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
