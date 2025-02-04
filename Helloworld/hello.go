package main

import (
	"fmt"
)

const (
	englishHelloPrefix = "Hello "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	hindiHelloPrefix   = "Namaskar, "

	spanish = "Spanish"
	french  = "French"
	hindi   = "Hindi"
)

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {

	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case hindi:
		prefix = hindiHelloPrefix
	default:
		prefix = englishHelloPrefix

	}

	return
}

func main() {

	fmt.Println(Hello("Mohit", ""))
}
