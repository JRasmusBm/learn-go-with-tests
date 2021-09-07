package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
  checkSums := func(t testing.TB, got, want []int) {
    t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v got %v", got, want)
		}
  }
	t.Run("- sums up the tails of each of the slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

    checkSums(t, got, want)
	})

	t.Run("- does not break on empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{})
		want := []int{2, 0}

    checkSums(t, got, want)
	})
}
