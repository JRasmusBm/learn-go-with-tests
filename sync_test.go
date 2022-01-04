package sync

import (
	"testing"
)

func assertCounterValue(t testing.TB, counter Counter, want int) {
	t.Helper()

	got := counter.Value()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func incrementCounter(counter *Counter, times int) {
	for i := 0; i < times; i++ {
		counter.Inc()
	}
}

func TestSync(t *testing.T) {
	t.Run("Incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}

		incrementCounter(&counter, 3)

		assertCounterValue(t, counter, 3)
	})
}
