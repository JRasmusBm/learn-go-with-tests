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
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		if got != want {
			t.Errorf("got %v want %v, given %x", got, want, shape)
		}
	}
	t.Run("- Calculates the area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 20.0)

	})

	t.Run("- Calculates the area of a circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
