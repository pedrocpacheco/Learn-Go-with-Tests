package main

import "fmt"

// ! Separating our domains

// * The printed String is our domain
func Hello() string {
	return "Hello, World!"
}

// * fmt.Println() is an outside domain
func main() {
	fmt.Println(Hello())
}

// go run hello.go
