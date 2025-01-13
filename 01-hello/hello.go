package hello

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishPrefix = "Hello"
	spanishPrefix = "Hola"
	frenchPrefix  = "Bonjour"
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	if lang == spanish {
		return fmt.Sprintf("%s, %s!", spanishPrefix, name)
	}

	if lang == french {
		return fmt.Sprintf("%s, %s!", frenchPrefix, name)
	}

	return fmt.Sprintf("%s, %s!", englishPrefix, name)
}
