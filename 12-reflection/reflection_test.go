package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

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
		{
			"nested fields",
			Person{
				"Marc",
				Profile{99, "Bletchley Park"},
			},
			[]string{"Marc", "Bletchley Park"},
		},
		{
			"pointers to things",
			&Person{
				"Marc",
				Profile{99, "Bletchley Park"},
			},
			[]string{"Marc", "Bletchley Park"},
		},
		{
			"slices",
			[]Profile{
				{99, "Bletchley Park"},
				{100, "Cambridge"},
			},
			[]string{"Bletchley Park", "Cambridge"},
		},
		{
			"arrays",
			[2]Profile{
				{99, "Bletchley Park"},
				{100, "Cambridge"},
			},
			[]string{"Bletchley Park", "Cambridge"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
