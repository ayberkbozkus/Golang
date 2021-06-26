package main

import (
	"fmt"
	"strings"
)

type mile float64
type kilometer float64

func main() {

	var m1 mile
	m1 = 3.2
	fmt.Println(m1)
	fmt.Printf("%T %v", m1,m1)

	f1 := float64(4.4)

	// fmt.Println(f1 == m1) // error, different type

	fmt.Printf(" %T %v \n", (m1 + mile(f1)), (m1 + mile(f1)))
	fmt.Printf(" %T %v \n", (float64(m1) + f1), (float64(m1) + f1))

	k1 := kilometer(7.8)

	fmt.Printf("%T %v", k1,k1)

	// fmt.Printf(m1 + k1) // error

	m2 := mile(4.6)

	fmt.Println(m1 + m2)
	fmt.Println(m1 * m2)
	fmt.Printf("%.2f \n", m1 * m2)
	fmt.Println(m1 + m2 + 2.1)

	myString := "arin"

	fmt.Println(strings.ToUpper(myString))

	k2 := toKilometer(m1)
	fmt.Println(k2)

	m3 := toMile(k1)
	fmt.Println(m3)
}

func toKilometer(m mile) kilometer{
	return kilometer(m * 1.6)
}

func toMile(k kilometer) mile {
	return mile(k * 0.62)
}