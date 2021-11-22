package main

import (
	"fmt"
	"github.com/Mingsicogi/go_study/practice/dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "first world"}

	// search
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

	// add
	err = dictionary.Add("second", "bye world")
	if err != nil {
		fmt.Println(err)
	}
	value, err = dictionary.Search("second")
	if err == nil {
		fmt.Println(value)
	}

	err = dictionary.Add("second", "bye world")
	if err != nil {
		fmt.Println(err)
	}

	// update
	err = dictionary.Update("second", "good bye world")
	if err != nil {
		fmt.Println(err)
	}
	value, err = dictionary.Search("second")
	if err == nil {
		fmt.Println(value)
	}

	err = dictionary.Update("asldkfj", "good bye world")
	if err != nil {
		fmt.Println(err)
	}

	// delete
	dictionary.Delete("second")

	value, err = dictionary.Search("second")
	if err == nil {
		fmt.Println(value)
	} else {
		fmt.Println(err)
	}

}
