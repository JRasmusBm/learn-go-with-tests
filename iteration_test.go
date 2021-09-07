package iteration

import (
  "testing"
)

func TestRepeat(t *testing.T) {
  t.Run("- Repeats the provided symbol", func(t *testing.T) {
    result := Repeat("a")
    expected := "aaaaa"

    if result != expected {
      t.Errorf("result %s got %s", result, expected)
    }
  })
}
