package main

import (
	"testing"
)

func TestLookup(t *testing.T) {
	t.Run("When the key is in the dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}

		assertDefinition(t, dictionary, "test", "this is a test")
	})

	t.Run("When the key is not in the dictionary", func(t *testing.T) {
		t.Run("Returns an error", func(t *testing.T) {
			dictionary := Dictionary{"test": "this is a test"}

			_, err := dictionary.Lookup("hello")

			assertError(t, err, ErrNotFound)
		})
	})
}

func TestAdd(t *testing.T) {
	t.Run("Adds the provided word and definition to the dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("book", "a bound collection of pages meant for reading")

		assertError(t, err, nil)

		got, _ := dictionary.Lookup("book")
		want := "a bound collection of pages meant for reading"

		assertEqualStrings(t, got, want)
	})

	t.Run("Returns a useful error when word already exists", func(t *testing.T) {
		word := "book"
		definition := "a bound collection of pages meant for reading"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "another definition")

		assertError(t, err, ErrWordExists)
	})
}

func assertEqualStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Lookup(word)

	if err != nil {
		t.Errorf("Got unexpected error: %#v", err)
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}

}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %#v want %#v", got, want)
	}

}
