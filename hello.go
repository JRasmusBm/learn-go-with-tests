package main

import (
	"fmt"
)

const englishPrefix = "Hello, "
const spanishPrefix = "¡Hola, "
const frenchPrefix = "Bonjour, "
const swedishPrefix = "Hej, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishPrefix
	case "French":
		prefix = frenchPrefix
  case "Swedish":
    prefix = swedishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
