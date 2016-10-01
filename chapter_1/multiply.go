package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a int
	fmt.Println("> a:")
	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		handleError(err)
	}
	var b int
	fmt.Println("> b:")
	_, err = fmt.Scanf("%d", &b)
	if err != nil {
		handleError(err)
	}
	mupltiply(a, b)
}

func handleError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func mupltiply(a, b int) {
	fmt.Println(a, b)
	for i := 1; i < 20; i++ {
		result := i * a
		// if result is b multiplier
		if result%b == 0 {
			// convert result to string and add # to it
			fmt.Printf("%s ", strconv.Itoa(result)+"#")
		} else {
			fmt.Printf("%d ", result)
		}
	}
}
