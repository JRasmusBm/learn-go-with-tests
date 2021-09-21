package main

import (
	"testing"
)

func TestLookup(t *testing.T) {
	t.Run("When the key is in the dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}

		got, err := dictionary.Lookup("test")

		t.Run("Returns the corrresponding value", func(t *testing.T) {
			assertEqualStrings(t, got, "this is a test")
		})

		t.Run("Does not return an error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Got unexpected error: %#v", err)
			}
		})
	})

	t.Run("When the key is not in the dictionary", func(t *testing.T) {
		t.Run("Returns an error", func(t *testing.T) {
			dictionary := Dictionary{"test": "this is a test"}

			_, err := dictionary.Lookup("hello")

			assertError(t, err, ErrNotFound)
		})
	})
}

func assertEqualStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %#v want %#v", got, want)
	}

}
