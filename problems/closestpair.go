package main

import (
	"fmt"
	"math"
	"sort"
)

func closestPair(numbers []int) (int, int) {
	sort.Ints(numbers)
	var n1, n2 int
	diff := math.MaxInt32
	for i := 0; i < len(numbers)-1; i++ {
		curr := numbers[i]
		next := numbers[i+1]
		if next-curr < diff {
			diff = next - curr
			n1 = curr
			n2 = next
		}
	}
	return n1, n2
}

func main() {
	numbers := []int{10, 6, 2, 5}
	fmt.Println(closestPair(numbers))
}
