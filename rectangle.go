package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	a, b float64
}

func (r rectangle) display(){
	fmt.Println("Side lengths: ", r.a*2, " - ", r.b*2)
}

func (r rectangle) area() float64 {
	return r.a * r.b
}

func (r rectangle) circumference() float64 {
	return 2 * (r.a + r.b)
}

type circle struct {
	r float64
}

func (c circle) display(){
	fmt.Println("R: ", c.r)
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.r
}

func (c circle) diameter() float64 {
	return 2 * c.r
}

type shape interface {
	area() float64
	circumference() float64
}

func interfacefunc (i shape) {
	fmt.Println(i)
	fmt.Println(i.area())
	fmt.Println(i.circumference)
	fmt.Printf("%T", i)
	fmt.Println()
}

func printArea(shapes ...shape){
	for _, shape := range shapes{
		fmt.Println("Area: ", shape.area())
	}
}

func main(){
	r1 := rectangle{3,8}
	c1 := circle{3}
	r2 := rectangle{10,10}
	r1.display()
	fmt.Println("Area: ", r1.area())
	fmt.Println("Circumference: ", r1.circumference())
	fmt.Println(" -------------- ")
	interfacefunc(r1)
	fmt.Println(" -------------- ")
	interfacefunc(c1)
	fmt.Println(" -------------- ")
	printArea(r1,c1,r2)
}