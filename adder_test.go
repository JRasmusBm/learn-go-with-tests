package integers

import (
  "fmt"
  "testing"
)

func ExampleAdd()  {
  sum := Add(1, 5)

  fmt.Println(sum)
  // Output: 6
}

func TestAdder(t *testing.T) {
  t.Run("- 2 + 2 = 4", func(t *testing.T) {
    result := Add(2, 2)
    expected := 4

    if result != expected {
      t.Errorf("result %d got %d", result, expected)
    }
  })
}
