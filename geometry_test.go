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
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tableTest := range areaTests {
		got := tableTest.shape.Area()
		if got != tableTest.want {
			t.Errorf("got %v want %v", got, tableTest.want)
		}
	}
}
