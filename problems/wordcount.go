package main

import "fmt"

func main() {
	fmt.Println(wordcount([]string{"a", "b", "a", "c", "b"}))
	fmt.Println(wordcount([]string{"c", "b", "a"}))
	fmt.Println(wordcount([]string{"c", "c", "c", "c"}))
}

func wordcount(words []string) map[string]int {
	m := make(map[string]int)
	for _, w := range words {
		if count, ok := m[w]; ok {
			m[w] = count + 1
		} else {
			m[w] = 1
		}
	}
	return m
}
