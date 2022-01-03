package walk

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	t.Run("- Walks a string", func(t *testing.T) {
		var got []string
		want := []string{"Rasmus"}

		x := struct {
			Name string
		}{want[0]}

		walk(x, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Wrong calls, got %v want %v", got, want)
		}
	})
}
