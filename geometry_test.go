package geometry

import (
  "testing"
)

func TestPerimeter(t *testing.T) {
  t.Run("- Calculates the perimeter of a rectangle", func(t *testing.T) {
    rectangle := Rectangle{10.0, 10.0}
    got := rectangle.Perimeter()
    want := 40.0

    if got != want {
      t.Errorf("got %v want %v, given %v", got, want, rectangle)
    }
  })
}

func TestArea(t *testing.T) {
  t.Run("- Calculates the area of a rectangle", func(t *testing.T) {
    rectangle := Rectangle{10.0, 10.0}
    got := rectangle.Area()
    want := 20.0

    if got != want {
      t.Errorf("got %v want %v, given %v", got, want, rectangle)
    }
  })

  t.Run("- Calculates the area of a circle", func(t *testing.T) {
    circle := Circle{10.0}
    got := circle.Area()
    want := 314.1592653589793

    if got != want {
      t.Errorf("got %v got %v, given %v", got, want, circle)
    }
  })
}
