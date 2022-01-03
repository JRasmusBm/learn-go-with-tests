package racer

import (
	"net/http"
	"time"
)

func calculateDuration(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func Racer(a, b string) (winner string) {
	if calculateDuration(a) < calculateDuration(b) {
		return a
	}

	return b
}
