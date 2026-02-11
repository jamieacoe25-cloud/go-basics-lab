package main

import (
	"fmt"
	"strings"
)

// TODO:
// Count word frequency in a string
// Return map[string]int
func countCharInString(sentence string) map[string]int {
	arr := strings.Split(sentence, " ")
	var stringCountMap = map[string]int{}

	for _, word := range arr {
		if stringCountMap[word] == 0 {
			stringCountMap[word] = 1
		} else {
			num := stringCountMap[word]
			stringCountMap[word] = num + 1
		}
	}

	return stringCountMap
}

func main() {
	fmt.Println(countCharInString("hello hello my name is jamie"))
	fmt.Println(countCharInString("this is this is me"))
}
