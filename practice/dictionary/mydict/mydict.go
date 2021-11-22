package mydict

import "errors"

type Dictionary map[string]string

var NotFoundData = errors.New("Not found data")

func (d Dictionary) Search(key string) (string, error) {
	value, exists := d[key]
	if exists {
		return value, nil
	} else {
		return "", NotFoundData
	}
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)
	if err != nil {
		if err == NotFoundData {
			d[key] = value
		}
		return nil
	} else {
		return errors.New("Already exists")
	}
}
