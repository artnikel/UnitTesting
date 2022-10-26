package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(SearchString("Hello", 4))
	//fmt.Println(ContainString("Hello Guys", "Guy"))
	fmt.Println(LenString("Hello World"))
}

func LenString(input string) int {
	count := len(input)
	return count
}

func ContainString(input string, word string) bool {
	contain := strings.Contains(input, word)
	return contain
}

func SearchString(input string, number int) byte {
	search := input[number]
	return search
}
