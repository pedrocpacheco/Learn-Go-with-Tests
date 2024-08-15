package main

import "fmt"

// ! Separating our domains

// * The printed String is our domain
func Hello(name string) string {
	return "Hello, " + name + "!"
}

// * fmt.Println() is an outside domain
func main() {
	fmt.Println(Hello("world"))
}

// go run hello.go
