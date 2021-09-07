package arrays_and_slices

import (
	"testing"
  "reflect"
)

func TestSumAll(t *testing.T) {
	t.Run("- Sums each of the collections", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v got %v", got, want)
    }
	})
}
