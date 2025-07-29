package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "

	spanish = "Spanish"
	french  = "French"
)

func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(language)
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}
