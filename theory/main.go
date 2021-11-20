package main

import (
	"fmt"
	"github.com/Mingsicogi/go_study/theory/utils"
	"strings"
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

	fmt.Println(nakedReturn("minssogi"))
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
	length = len(name)
	upperName = strings.ToUpper(name)
	return
}
