package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{10, 9, 2, 2, 5, 4, 12, 3, 6, 7, 6, 4}
	fmt.Println(dedup(numbers))
}

func dedup(numbers []int) []int {
	sort.Ints(numbers)
	dedups := make([]int, 0)
	dedups = append(dedups, numbers[0])
	for i, n := range numbers {
		if i != 0 && n != numbers[i-1] {
			dedups = append(dedups, n)
		}
	}
	return dedups
}
