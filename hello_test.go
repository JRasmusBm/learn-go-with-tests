package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertEqual := func(t testing.TB, result, expected string) {
		t.Helper()

		if result != expected {
			t.Errorf("result %s expected %s", result, expected)
		}
	}

	t.Run("- greets person with specified name", func(t *testing.T) {
		result := Hello("Peter", "")
		expected := "Hello, Peter!"

		assertEqual(t, result, expected)
	})

	t.Run("- greets world on empty string", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, world!"

		assertEqual(t, result, expected)
	})

	t.Run("- speaks Spanish when so specified", func(t *testing.T) {
		result := Hello("Elodie", "Spanish")
		expected := "¡Hola, Elodie!"

		assertEqual(t, result, expected)
	})

	t.Run("- speaks French when so specified", func(t *testing.T) {
		result := Hello("Marc", "French")
		expected := "Bonjour, Marc!"

		assertEqual(t, result, expected)
	})

	t.Run("- speaks Swedish when so specified", func(t *testing.T) {
		result := Hello("Rasmus", "Swedish")
		expected := "Hej, Rasmus!"

		assertEqual(t, result, expected)
	})
}
