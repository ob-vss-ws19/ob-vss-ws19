package stringutil

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello", "olleH"},
		{"hallo", "ollah"},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
func ExampleReverse() {
	fmt.Print(Reverse("Hallo"))
	// Output:
	// ollaH
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("ljhdsagjlhfasdgljfadsgfljhasgfljahgfdsajlhgafsdljh")
	}
}
