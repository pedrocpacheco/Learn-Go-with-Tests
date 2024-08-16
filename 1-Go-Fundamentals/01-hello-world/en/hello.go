package main

import "fmt"

// Constants declared in bloco
const (
	spanish    = "Spanish"
	french     = "French"
	portuguese = "Portuguese"

	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	frenchHelloPrefix     = "Bonjour, "
	portugueseHelloPrefix = "Ola, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

/*
Private Function -> 1st letter in lowercase
Named Return -> Cria variavel "prefix" dentro da função
*/
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return // retorna a variavel nomeada na assinatura
}

func main() {
	fmt.Println(Hello("world", ""))
}
