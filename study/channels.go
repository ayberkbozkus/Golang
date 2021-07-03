package main

import (
	"fmt"
	"math"
	// "sync"
)

type circle struct {
	r float64
}

func (c circle) display() {
	fmt.Println("A Circle")
}

// func (c circle) area() float64 {
// 	return math.Pi * c.r * c.r
// }

func hello(channel chan string){
	channel <- "Hello!"
}


func area( c circle, myChannel chan float64){
	result := math.Pi * c.r * c.r
	myChannel <- result
}

func main() {
	
	var myChannel chan float64
	myChannel = make(chan float64)
	
	// myChannel := make(chan string)

	c1 := circle{5}
	
	go area(c1, myChannel)
	fmt.Printf("%.2f\n", <- myChannel)

	c1.display()

	// go hello(myChannel)

	// fmt.Println(<-myChannel)
}

