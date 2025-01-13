package hello

import "fmt"

const englishPrefix = "Hello"

func Hello(name string) string {
	return fmt.Sprintf("%s, %s!", englishPrefix, name)
}
