package main

import (
	"errors"
)

var ErrNotFound = errors.New("Key not found in dictionary")

type Dictionary map[string]string

func (d Dictionary) Lookup(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
