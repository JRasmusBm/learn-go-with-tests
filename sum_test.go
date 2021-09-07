package arrays_and_slices

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("- sums up collections of integers of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("want %v got %v, given %v", got, want, numbers)
		}
	})
}
