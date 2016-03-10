package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	p := strings.Fields(s)
	ages := make(map[string]int)
	for i, _ := range p {
		ages[p[i]]++
	}

	return ages
}

func main() {
	wc.Test(WordCount)
}
