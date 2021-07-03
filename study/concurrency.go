package main

import (
	"fmt"
	// "time"
	"sync"
)

func printX() {
	for i:= 0; i < 20; i++ {
		fmt.Print("X")
	}
	wg.Done()
}

func printY() {
	for i:= 0; i < 20; i++ {
		fmt.Print("Y")
	}
}

var wg sync.WaitGroup

func main() {

	wg.Add(1) // There are 2 goroutines external to main

	go printX()
	// go printY()
	wg.Wait() // wait goroutines
	printY()
	// time.Sleep(time.Second)
}