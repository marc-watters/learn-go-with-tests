package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("\nrepeated:\t%q\nexpected:\t%q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	count := b.N
	for i := 0; i < count; i++ {
		Repeat("a", count)
	}
}
