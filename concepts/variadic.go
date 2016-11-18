package main

import (
	"fmt"
)

func main() {
	sum1 := sum(1, 2, 3, 4, 5)
	fmt.Println(sum1)

	sum2 := sum([]int{1, 2, 3, 4}...)
	fmt.Println(sum2)
}

func sum(nums ...int) (result int) {
	for _, num := range nums {
		result += num
	}
	return
}
