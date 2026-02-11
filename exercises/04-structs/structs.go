package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func newUser(name string, age int) User {
	return User{name, age}
}

func (u *User) isAdult() bool {
	return u.Age >= 18
}

func main() {
	child := newUser("Charile", 1)
	fmt.Println(child.isAdult())

	adult := newUser("Jamie", 30)
	fmt.Println(adult.isAdult())
}

// TODO:
// Add a method IsAdult() bool
