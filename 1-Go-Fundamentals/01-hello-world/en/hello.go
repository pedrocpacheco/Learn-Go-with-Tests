package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	// Default Prefix
	prefix := englishHelloPrefix

	// Switch-Case to alter Prefix if needed
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case "Portuguese":
		prefix = "Ola, "
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
