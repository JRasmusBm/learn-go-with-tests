package main

import (
	"testing"
)

func TestHello(t *testing.T) {
  assertEqual := func(t testing.TB, result, expected string) {
    t.Helper()

    if result != expected {
      t.Errorf("result %s got %s", result, expected)
    }
  }

	t.Run("- greets people", func(t *testing.T) {
		result := Hello("world")
		expected := "Hello, world!"

		assertEqual(t, result, expected)
	})

	t.Run("- greets world on empty string", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, world!"

		assertEqual(t, result, expected)
	})
}
