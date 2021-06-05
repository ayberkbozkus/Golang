// Package clause
package main

// Import statement
import (
	"fmt"
	"strconv"
)

var packVar = "Test"

// My Code
func main() {

	var (
		name, lastname string = "tester", "second name"
		name1          string = "secondname"
		age            int    = 40
		isMarried      bool   = true

		height float32 = 72.5
		weight int 	   = 172
	)
	height := 17.5
	var packVar = "Test2"

	fmt.Println("Hello World", name, name1, lastname, age, isMarried, height)
	fmt.Printf("%T\n%T\n%T\n", name, age, isMarried)
	funcVar()
	fmt.Println(packVar)
	fmt.Printf("%v = %T", int(height), height)

	// tye corversion
	height = float32(weight)

	// change value
	name, name1 = name1, name

	if true {
		fmt.Println("name1", name1, "name", name)
	}

	const x := 65 // constant --> compile time

	s := string(x) // ---> run time

	fmt.Printf("%v, %T\n", x, x) // 65
	fmt.Printf("%v, %T\n", s, s) // A

	y := strconv.Itoa(x)
	fmt.Printf("%v, %T\n", y, y) // 65

	const k = 3 // typless

	

}

func funcVar() {
	fmt.Println(packVar)
}
