package iteration

import (
	"fmt"
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("- Repeats the provided symbol n times", func(t *testing.T) {
		result := Repeat("a", 5)
		expected := "aaaaa"

		if result != expected {
			t.Errorf("result %s got %s", result, expected)
		}
	})
}
