package go_specs_greet

import "fmt"

func Introduce(name string) string {
	return fmt.Sprintf("An honor to meet you, %s.", name)
}
