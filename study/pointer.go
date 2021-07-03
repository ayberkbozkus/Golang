package main

import (
	"fmt"
)

func main() {

	name := "Arin"

	fmt.Println(name)
	fmt.Println(&name)

	x := 22
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%T, %v \n", x, x)
	fmt.Printf("%T, %v \n", &x, &x)

	var y *int = &x // just int pointer not other type
	fmt.Println(y)
	fmt.Printf("%T, %v \n", y, y) 

	z := &name // just string pointer
	fmt.Printf("%T, %v \n", z, z) 

	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(*(&x)) // dereferencing

	x1 := &x

	fmt.Println(x,x1)
	fmt.Println(x,*x1)

	w1 := [4]int{1, 10, 100, 1000}
	w2 := w1

	fmt.Println(w1, "\t", w2)

	w2[0] = 3

	fmt.Println(w1, "\t", w2)

	s1 := []int{1, 10, 100, 1000}
	s2 := s1

	fmt.Println(s1, "\t", s2)

	s2[0] = 3

	fmt.Println(s1, "\t", s2)
	
	mySlc := []int{1,10,100}

	double(mySlc) 
	fmt.Println(mySlc) 

	num := 5
	doubleValue(num) 
	doubleValueP(&num)
	fmt.Println(num)
}

func double(slc []int){ // pass by refence
	for i :=0; i< len(slc); i++ {
		slc[i] *=2
	}
	fmt.Println(slc)
}

func doubleValue(num int){ // pass by value
	num *= 2
	fmt.Println(num)
}

func doubleValueP(num *int){ // pass by value
	*num *= 2
	fmt.Println(*num)
}