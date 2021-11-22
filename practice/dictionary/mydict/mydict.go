package mydict

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, exists := d[key]
	if exists {
		return value, nil
	} else {
		return "", errors.New("Not found data")
	}
}
