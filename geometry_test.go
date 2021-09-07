package geometry

import (
  "testing"
)

func TestPerimeter(t *testing.T) {
  t.Run("- Calculates the perimeter of a triangle", func(t *testing.T) {
    got := Perimeter(10.0, 10.0)
    want := 40.0

    if got != want {
      t.Errorf("got %v got %v", got, want)
    }
  })
}

func TestArea(t *testing.T) {
  t.Run("- Calculates the perimeter of a triangle", func(t *testing.T) {
    got := Area(10.0, 10.0)
    want := 20.0

    if got != want {
      t.Errorf("got %v got %v", got, want)
    }
  })
}
