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

	prefix := greeting(lang)

	return fmt.Sprintf("%s, %s!", prefix, name)
}

func greeting(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}

	return
}
