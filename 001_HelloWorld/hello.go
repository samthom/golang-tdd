package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const spanishPrefix = "Hola, "
const englishPrefix = "Hello, "
const frenchPrefix = "Bonjour, "

// Hello return string with greeting
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "English"))
}
