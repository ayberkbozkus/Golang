package main

import "fmt"

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

func main(){
	r1 := rectangle{3,8}
	r1.display()
	fmt.Println("Area: ", r1.area())
	fmt.Println("Circumference: ", r1.circumference())
	fmt.Println(" -------------- ")
	interfacefunc(r1)
}