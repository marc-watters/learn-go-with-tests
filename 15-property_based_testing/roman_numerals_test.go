package romannumerals

import "testing"

func TestRomanNumerals(t *testing.T) {
	t.Run("converts 1 to I", func(t *testing.T) {
		got := ConvertToRoman(1)
		want := "I"

		if got != want {
			t.Errorf("\ngot: \t%q\nwant:\t%q", got, want)
		}
	})

	t.Run("2 gets converted to II", func(t *testing.T) {
		got := ConvertToRoman(2)
		want := "II"

		if got != want {
			t.Errorf("\ngot: \t%q\nwant:\t%q", got, want)
		}
	})
}
