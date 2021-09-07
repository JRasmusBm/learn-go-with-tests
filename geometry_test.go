package geometry

import (
  "testing"
)

func TestPerimeter(t *testing.T) {
  t.Run("- Calculates the perimeter of a rectangle", func(t *testing.T) {
    rectangle := Rectangle{10.0, 10.0}
    got := Perimeter(rectangle)
    want := 40.0

    if got != want {
      t.Errorf("got %v got %v, given %v", got, want, rectangle)
    }
  })
}

func TestArea(t *testing.T) {
  t.Run("- Calculates the perimeter of a rectangle", func(t *testing.T) {
    rectangle := Rectangle{10.0, 10.0}
    got := Area(rectangle)
    want := 20.0

    if got != want {
      t.Errorf("got %v got %v, given %v", got, want, rectangle)
    }
  })
}
