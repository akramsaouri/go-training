package main

import (
	"fmt"
)

func main() {
	wanted := 6
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	found := false

	for found == false {
		middle := array[len(array)/2-1]
		fmt.Println("array ", array)
		fmt.Println("middle ", middle)

		if middle == wanted {
			found = true
			break
		}
		if wanted >= middle {
			// select right side
			array = array[pos(array, middle):]
		} else {
			// select left side
			array = array[:pos(array, middle)]
		}
	}
	fmt.Println("found it!")

}

func pos(slice []int, value int) int {
	for e, i := range slice {
		if e == value {
			return i
		}
	}
	return -1
}
