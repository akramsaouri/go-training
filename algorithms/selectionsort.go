package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := rand.Perm(25)
	fmt.Println(array)

	for selection := range array {
		indexOfMin := selection
		for i := selection; i < len(array); i++ {
			if array[i] < array[indexOfMin] {
				indexOfMin = i
			}
		}
		temp := array[selection]
		array[selection] = array[indexOfMin]
		array[indexOfMin] = temp
	}

	fmt.Println(array)
}
