package test

import "fmt"

func Hello() {
	fmt.Println("Hello World")
}

func Add(a int, b int) int {
	return a + b
}

type User struct {
	Name string
	Age  int
}

const Version = "1.0.0"
