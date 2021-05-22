// Package clause
package main

// Import statement
import "fmt"

var packVar = "Test"

// My Code
func main() {

	var (
		name, lastname string = "tester", "second name"
		name1          string = "secondname"
		age            int    = 40
		isMarried      bool   = true
	)
	height := 17.5
	var packVar = "Test2"

	fmt.Println("Hello World", name, name1, lastname, age, isMarried, height)
	fmt.Printf("%T\n%T\n%T\n", name, age, isMarried)
	funcVar()
	fmt.Println(packVar)
	fmt.Printf("%v = %T", int(height), height)
}

func funcVar() {
	fmt.Println(packVar)
}
