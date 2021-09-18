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
	t.Run("Table Test", func(t *testing.T) {
		areaTests := []struct {
			shape Shape
			want  float64
		}{
			{shape: Rectangle{12, 6}, want: 72.0},
			{shape: Circle{10}, want: 314.1592653589793},
			{shape: Triangle{12, 6}, want: 36.0},
		}

		for _, tableTest := range areaTests {
			got := tableTest.shape.Area()
			if got != tableTest.want {
				t.Errorf("%#v, got %v want %v", tableTest.shape, got, tableTest.want)
			}
		}
	})
}
