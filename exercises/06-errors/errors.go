package main

import (
	"errors"
	"fmt"
)

func divide(one, two float32) (float32, error) {
	if one == 0 || two == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	result := one / two
	return result, nil
}

func main() {
	result, err := divide(0, 1)
	fmt.Println(result)
	fmt.Println(err)

	result, err = divide(1, 0)
	fmt.Println(result)
	fmt.Println(err)

	result, err = divide(2, 1)
	fmt.Println(result)
	fmt.Println(err)
}

// TODO:
// Write a function that divides two ints
// Return an error when dividing by zero
