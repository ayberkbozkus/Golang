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

	height = 17.5
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

	const x = 65 // constant --> compile time

	s := string(x) // ---> run time

	fmt.Printf("%v, %T\n", x, x) // 65
	fmt.Printf("%v, %T\n", s, s) // A

	y := strconv.Itoa(x)
	fmt.Printf("%v, %T\n", y, y) // 65

	const k = 3 // typless

	a,b := 15,20

	set1 := (a == b)
	set2 := (a < b)

	fmt.Printf("%v\n", (set1 && set2)) // AND
	fmt.Printf("%v\n", (set1 || set2)) // OR
	fmt.Println(sum(5,10))
}

func funcIf() {
	

	if x := 27; x%2 == 0 {
		fmt.Println(x, "even value") // fmt.Printf("%v", x) 
	} else if x%2 != 0 {
		fmt.Println(x, "odd value")
	} else {
		fmt.Println("Error")
	}
}

func funcSwitch(){
	grade := 3
	switch grade {
	case 5:
		fmt.Println("very good")
	
	case 4:
		fmt.Println("good")
		fallthrough // continue
	
	case 3:
		fmt.Println("okay")
	
	case 2:
		fmt.Println("not bad")
	
	case 1:
		fmt.Println("bad")
	
	default:
		fmt.Println("invalid grade")
	}
}

func funcFor(){
	
	
	for i :=1; i<=10; i++ {
		fmt.Println(i)
	}

	k := 10
	for k<=100{ // while loop
		fmt.Println(k)
		k += 10
	}

	for i := 0; i<10; i++ {
		if i%3 ==0{ // go start
			continue
		}
		if i%8 ==0{  // end loop
			break
		}
		fmt.Println(i)
	}

}

func funcVar() {
	fmt.Println(packVar)
}

func sum(x,y int) int {
	return x+y
}

