package hello

import "fmt"

const (
	spanish = "Spanish"

	englishPrefix = "Hello"
	spanishPrefix = "Hola"
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	if lang == spanish {
		return fmt.Sprintf("%s, %s!", spanishPrefix, name)
	}

	return fmt.Sprintf("%s, %s!", englishPrefix, name)
}
