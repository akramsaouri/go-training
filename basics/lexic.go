package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("> first_word:")
	first_word, _ := reader.ReadString('\n')
	fmt.Println("> second_word:")
	second_word, _ := reader.ReadString('\n')
	fmt.Println(lexic(first_word, second_word))
}

func lexic(first_word, second_word string) (longest string) {
	if len(first_word) > len(second_word) {
		longest = first_word
	} else {
		longest = second_word
	}
	return
}
