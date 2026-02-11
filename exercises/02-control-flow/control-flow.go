// TODO:
// Write a function that returns:
// - "even" if number is even
// - "odd" if number is odd
package main

import "fmt"

func eventOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	}
	return "odd"
}

func main() {
	fmt.Println(eventOrOdd(1))
	fmt.Println(eventOrOdd(2))
	fmt.Println(eventOrOdd(3))
}
