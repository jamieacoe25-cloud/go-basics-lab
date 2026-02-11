// TODO:
// Write a function that returns min and max of two ints
package main

import "fmt"

func minMax(one, two int) int {
	if one == two {
		return 0
	} else if one > two {
		return one
	}
	return two
}

func main() {
	fmt.Println(minMax(1, 2))
	fmt.Println(minMax(2, 1))
	fmt.Println(minMax(1, 1))
}
