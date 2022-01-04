package sync

import (
	"sync"
	"testing"
)

func assertCounterValue(t testing.TB, counter *Counter, want int) {
	t.Helper()

	got := counter.Value()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSync(t *testing.T) {
	t.Run("Incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		want := 3

		for i := 0; i < want; i++ {
			counter.Inc()
		}

		assertCounterValue(t, &counter, want)
	})

	t.Run("Runs safely concurrently", func(t *testing.T) {
		counter := Counter{}
		want := 1000

		var wg sync.WaitGroup
		wg.Add(want)

		for i := 0; i < want; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounterValue(t, &counter, want)
	})
}
