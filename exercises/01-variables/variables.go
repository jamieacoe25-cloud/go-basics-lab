// TODO:
// 1. Declare an int variable
// 2. Declare a string variable
// 3. Write a function that returns both
package main

import "fmt"

func DeclareVariables() (int, string) {
	number := 2
	word := "string"
	return number, word

}

func main() {
	number, world := DeclareVariables()
	fmt.Println("number:", number)
	fmt.Println("string: ", world)
}
