package reflection

import "reflect"

func walk(x any, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := range val.NumField() {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
