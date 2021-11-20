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
