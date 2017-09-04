package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3}
	fmt.Println(dup(numbers, 3))
}

func dup(numbers []int, times int) []int {
	result := make([]int, 0)
	for _, n := range numbers {
		result = append(result, repeat(n, times)...)
	}
	return result
}

func repeat(number int, times int) []int {
	arr := make([]int, 0)
	for i := 0; i < times; i++ {
		arr = append(arr, number)
	}
	return arr
}
