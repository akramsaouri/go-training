package main

import "fmt"

func main() {
	array := [5]int{3, 2, 5, 4, 1}

	fmt.Println(array)

	start := 0
	for _ = range array {
		for i := len(array) - 1; i > start; i-- {
			if array[i] < array[i-1] {
				temp := array[i]
				array[i] = array[i-1]
				array[i-1] = temp
			}
			if i == start+1 {
				start++
			}
		}
	}
	fmt.Println(array)
}
