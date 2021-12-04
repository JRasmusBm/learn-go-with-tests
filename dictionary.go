package main

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

var (
	ErrNotFound   = DictionaryErr("Key not found in dictionary")
	ErrWordExists = DictionaryErr("Word already exists")
)

type Dictionary map[string]string

func (d Dictionary) Lookup(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Lookup(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) {
	d[word] = definition
}
