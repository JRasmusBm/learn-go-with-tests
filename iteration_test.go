package iteration

import (
  "testing"
)

func BenchmarkRepeat(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Repeat("a")
  }
}

func TestRepeat(t *testing.T) {
  t.Run("- Repeats the provided symbol", func(t *testing.T) {
    result := Repeat("a")
    expected := "aaaaa"

    if result != expected {
      t.Errorf("result %s got %s", result, expected)
    }
  })
}
