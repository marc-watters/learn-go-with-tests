package reflection

import "reflect"

func walk(x any, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := range val.NumField() {
		field := val.Field(i)

		if field.Kind() == reflect.String {
			fn(field.String())
		}

		if field.Kind() == reflect.Struct {
			walk(field.Interface(), fn)
		}
	}
}
