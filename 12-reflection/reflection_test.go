package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct{ Name string }{"Marc"},
			[]string{"Marc"},
		},
		{
			"struct with two string field",
			struct {
				Name string
				City string
			}{"Marc", "Bletchley Park"},
			[]string{"Marc", "Bletchley Park"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Marc", 99},
			[]string{"Marc"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("\ngot: \t%v\nwant:\t%v\n", got, test.ExpectedCalls)
			}
		})
	}
}
