package main

import "fmt"

const (
	englishGreetingPrefix = "Hello, "
	spanishGreetingPrefix = "Hola, "
	frenchGreetingPrefix  = "Bonjour, "

	spanish = "Spanish"
	french  = "French"
)

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishGreetingPrefix
	case french:
		prefix = frenchGreetingPrefix
	default:
		prefix = englishGreetingPrefix
	}
	return
}

func getHello(name, language string) string {
	if len(name) == 0 {
		name = "World!"
	}
	return getGreetingPrefix(language) + name
}

func main() {
	fmt.Println(getHello("Anuj", ""))
}
