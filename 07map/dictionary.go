package main

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definiton, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return definiton, nil
}
