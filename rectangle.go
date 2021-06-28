package main

import "fmt"

type rectangle struct {
	a, b int
}

func (r rectangle) display(){
	fmt.Println("Side lengths: ", r.a*2, " - ", r.b*2)
}

func (r rectangle) area() int {
	return r.a * r.b
}

func (r rectangle) circumference() int {
	return 2 * (r.a + r.b)
}

func main(){
	r1 := rectangle{3,8}
	r1.display()
	fmt.Println("Area: ", r1.area())
	fmt.Println("Circumference: ", r1.circumference())
}