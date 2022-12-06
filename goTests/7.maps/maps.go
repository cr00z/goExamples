package maps

import "errors"

var (
	ErrWordExists        = errors.New("word exists")
	ErrWordDoesNotExists = errors.New("word does not exists")
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, bool) {
	value, inDict := d[key]
	return value, inDict
}

func (d Dictionary) Add(key string, value string) error {
	if _, inDict := d.Search(key); inDict {
		return ErrWordExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key string, value string) error {
	if _, inDict := d.Search(key); !inDict {
		return ErrWordDoesNotExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) error {
	// if _, inDict := d.Search(key); !inDict {
	// 	return ErrWordDoesNotExists
	// }
	delete(d, key)
	return nil
}
