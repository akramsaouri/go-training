package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(occurence("this is to test test ok ? "))
	fmt.Println(caraOccurence("this is to test test ok ? "))
}

func occurence(phrase string) map[string]int {
	occu := make(map[string]int)
	words := strings.Fields(phrase)
	for _, word := range words {
		_, exist := occu[word]
		if exist {
			occu[word] += 1
		} else {
			occu[word] = 1
		}
	}
	return occu
}

func caraOccurence(phrase string) map[string]int {
	occu := make(map[string]int)
	letters := strings.Split(phrase, "")
	for _, letter := range letters {
		if letter != " " {
			_, exist := occu[letter]
			if exist {
				occu[letter] += 1
			} else {
				occu[letter] = 1
			}
		}
	}
	return occu
}
