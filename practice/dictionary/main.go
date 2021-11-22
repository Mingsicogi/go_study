package main

import (
	"fmt"
	"github.com/Mingsicogi/go_study/practice/dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "first world"}

	value, err := dictionary.Search("second")
	if err == nil {
		fmt.Println(value)
	} else {
		fmt.Println(err)
	}

	value, err = dictionary.Search("first")
	if err == nil {
		fmt.Println(value)
	} else {
		fmt.Println(err)
	}
}
