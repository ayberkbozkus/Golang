package main

import "fmt"

func main() {
	underArray := [...]int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(underArray)

	mySlc := underArray[2:5]
	fmt.Println(mySlc)
	mySlc2 := underArray[:6]
	fmt.Println(mySlc2)
	mySlc3 := underArray[3:]
	fmt.Println(mySlc3)
	mySlc4 := underArray[:]
	fmt.Println(mySlc4)

	mySlc[0] = 100
	fmt.Println(mySlc)
	fmt.Println(mySlc2)
	fmt.Println(mySlc4)


	mySlc5 := []int{1,2,3}
	mySlc6 := append(mySlc5, 4, 5)
	fmt.Println(mySlc6)
	mySlc5[0] = 500
	fmt.Println(mySlc5)
	fmt.Println(mySlc6)

	mySlc6 =append(mySlc5,mySlc6...)
	fmt.Println(mySlc6)

	mySlc6 = mySlc6[2:] // remove the first 2 elements

	mySlc6 = mySlc6[:len(mySlc6)-3] //  remove the last 3 elements

	var mySlc7 []int
	fmt.Printf("%#v", mySlc7) // null, just declare 
	fmt.Println()
	var mySlc8 := make([]int,0) // empty Slice
	fmt.Printf("%#v", mySlc8)
}