package main

import (
	"fmt"
	"github.com/bradfitz/slice"
)

const LEN = 5

func main() {
	liste := [LEN]int{20, 18, 12, 2, 102}
	fmt.Println(liste)

	sort(liste[:])
	fmt.Println(liste)

	var portion []int
	portion = add(liste[:], 5)
	fmt.Println(portion)

	reverse(portion)
	fmt.Println(portion)

	sortDesc(portion)
	fmt.Println(portion)

	portion = pull(portion, 1)
	fmt.Println(portion)

}

func sort(portion []int) {
	slice.Sort(portion, func(i, j int) bool {
		return portion[i] < portion[j]
	})
}

func sortDesc(portion []int) {
	slice.Sort(portion, func(i, j int) bool {
		return portion[i] > portion[j]
	})
}

func add(portion []int, elm int) []int {
	return append(portion, elm)
}

func reverse(portion []int) {
	var temp int
	j := 0
	for i := len(portion) - 1; i > -1; i-- {
		temp = portion[j]
		portion[j] = portion[i]
		portion[i] = temp
		j += 1
	}
}

func pull(portion []int, index int) []int {
	return append(portion[:index], portion[index+1:]...)
}
