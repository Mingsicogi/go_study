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

func sum1to10(numbers ...int) (total int) {
	for i, number := range numbers {
		fmt.Println(i, number)
		total += number
	}

	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return
}
