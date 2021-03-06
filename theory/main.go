package main

import (
	"fmt"
	"github.com/Mingsicogi/go_study/theory/utils"
	"strings"
	"time"
)

func main() {
	fmt.Print("Hello Go!")
	utils.SayBye()

	println()

	//const name string = "minssogi"
	name := "minssogi"
	fmt.Println(name)

	fmt.Println(multiply(200, 4))

	lengthOfName, upperCaseOfName := lengthAndUpperStr("minssogi")
	fmt.Println(lengthOfName, upperCaseOfName)

	dynamicArguments("mins", "jings", "jeon", "eun", "mike")

	a, b := nakedReturn("minssogi")
	fmt.Println(a, b)

	fmt.Println(sum1to10(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	fmt.Println(canIDrink(19))

	fmt.Println(isItFruit("orange"))

	fmt.Println("================================")

	pointer()

	arrayAndSlice()

	mapPractice()

	fmt.Println(createPerson("minssogi", 20, "computer", "exercise", "shopping"))
}

// simple function
func multiply(a int, b int) int {
	return a * b
}

// multiple return
func lengthAndUpperStr(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// dynamic arguments
func dynamicArguments(names ...string) {
	fmt.Println(names)
}

// naked return
func nakedReturn(name string) (length int, upperName string) {
	defer forDeferTest("Additional nakedReturn function finish!")
	length = len(name)
	upperName = strings.ToUpper(name)
	fmt.Println("nakedReturn business logic finish")
	return
}

// defer => execute command after function finished
func forDeferTest(message string) {
	time.Sleep(1 * time.Second)
	fmt.Println(message)
}

// for loop
func sum1to10(numbers ...int) (total int) {
	for _, number := range numbers {
		//fmt.Println(i, number)
		total += number
	}

	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return
}

// if - variable expression
func canIDrink(age int) bool {
	if checkAge := age - 1; checkAge > 18 {
		return true
	}

	return false
}

// switch case
func isItFruit(something string) bool {
	switch {
	case "apple" == strings.ToLower(something):
		return true
	case "banana" == strings.ToLower(something):
		return true
	case "orange" == strings.ToLower(something):
		return true
	case "grape" == strings.ToLower(something):
		return true
	}

	return false
}

// & => address, * => see though
func pointer() {
	a := 10
	b := &a

	fmt.Println("a =", a, &a)
	fmt.Println("b =", b, &b, *b)

	a = 100000
	fmt.Println("change a value as", a)

	fmt.Println("a =", a, &a)
	fmt.Println("b =", b, &b, *b)
}

// array, slice
func arrayAndSlice() {
	fruits := [5]string{"apple", "banana", "orange"}
	fmt.Println(fruits)

	names := []string{"mins", "jeon"}
	newNames := append(names, "mark")
	fmt.Println(newNames)
}

// map
func mapPractice() {
	empInfo := map[string]string{"name": "minssogi", "age": "20"}
	fmt.Println(empInfo)

	for key, value := range empInfo {
		fmt.Println(key, value)
	}
}

// struct
type person struct {
	name      string
	age       int
	favorites []string
}

func createPerson(name string, age int, favorites ...string) (newPerson person) {
	newPerson = person{name: name, age: age, favorites: favorites}
	return
}
