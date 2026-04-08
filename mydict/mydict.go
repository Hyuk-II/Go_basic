package mydict

import "errors"

// Dictionary
type Dictionary map[string]string

var errNotFound = errors.New("Not found")
var errCantUPdate = errors.New("Can't update non-existing word")
var errWordExists = errors.New("That word is already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]

	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word
func (d Dictionary) Add(word, def string) error {
	_, exists := d.Search(word)

	switch exists {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictionary) Update(word, def string) error {
	_, exists := d.Search(word)

	switch exists {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUPdate
	}

	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}