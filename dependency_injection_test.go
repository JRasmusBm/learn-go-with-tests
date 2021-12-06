package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("Addresses the greeting to the supplied name", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Rasmus")

		got := buffer.String()
		want := "Hello, Rasmus"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
