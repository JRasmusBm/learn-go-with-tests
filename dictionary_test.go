package main

import (
	"testing"
)

func TestLookup(t *testing.T) {
	word := "book"
	definition := "a bound collection of pages meant for reading"

	t.Run("When the key is in the dictionary", func(t *testing.T) {
		t.Run("Returns its definition", func(t *testing.T) {
			dictionary := Dictionary{word: definition}

			assertDefinition(t, dictionary, word, definition)
		})
	})

	t.Run("When the key is not in the dictionary", func(t *testing.T) {
		t.Run("Returns an error", func(t *testing.T) {
			dictionary := Dictionary{word: definition}

			_, err := dictionary.Lookup("hello")

			assertError(t, err, ErrNotFound)
		})
	})
}

func TestAdd(t *testing.T) {
	word := "book"
	definition := "a bound collection of pages meant for reading"

	t.Run("When the word is not in the dictionary", func(t *testing.T) {
		t.Run("Adds the word and its definition to the dictionary", func(t *testing.T) {
			dictionary := Dictionary{}
			err := dictionary.Add(word, definition)

			assertError(t, err, nil)

			assertDefinition(t, dictionary, word, definition)
		})
	})

	t.Run("When the word is already in the dictionary", func(t *testing.T) {
		t.Run("Returns an error", func(t *testing.T) {
			word := "book"
			definition := "a bound collection of pages meant for reading"
			dictionary := Dictionary{word: definition}

			err := dictionary.Add(word, "another definition")

			assertError(t, err, ErrWordExists)
		})
	})
}

func TestUpdate(t *testing.T) {
	word := "book"
	definition := "a bound collection of pages meant for reading"

	t.Run("When the word is in the dictionary", func(t *testing.T) {
		t.Run("Sets a new definition for the word", func(t *testing.T) {
			dictionary := Dictionary{word: definition}
			newDefinition := "an item that rests in bookshelves"

			err := dictionary.Update(word, newDefinition)

			assertError(t, err, nil)
			assertDefinition(t, dictionary, word, newDefinition)
		})
	})

	t.Run("When the word is not in the dictionary", func(t *testing.T) {
		t.Run("Returns a useful error", func(t *testing.T) {
			dictionary := Dictionary{}
			newDefinition := "an item that rests in bookshelves"

			err := dictionary.Update(word, newDefinition)

			assertError(t, err, ErrWordDoesNotExist)
		})
	})
}

func TestDelete(t *testing.T) {
	word := "book"
	definition := "a bound collection of pages meant for reading"

	t.Run("When the word is in the dictionary", func(t *testing.T) {
		t.Run("Deletes the word from the dictionary", func(t *testing.T) {
			dictionary := Dictionary{word: definition}

			dictionary.Delete(word)

			_, err := dictionary.Lookup(word)

			if err != ErrNotFound {
				t.Errorf("Expected %v to be deleted", word)
			}
		})
	})

	t.Run("When the word is not in the dictionary", func(t *testing.T) {
		t.Run("Silently does nothing", func(t *testing.T) {
			dictionary := Dictionary{}

			dictionary.Delete(word) // Does not panic!
		})
	})
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
