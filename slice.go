package main

import "fmt"

func main() {
	mySlc := []int {1,2,3}
	fmt.Println(mySlc)
	fmt.Println(len(mySlc))

	var mySlcEmpty []int
	fmt.Println(mySlcEmpty)
	fmt.Println(len(mySlcEmpty))

	mySlcEmpty = make([]int, 4)

	fmt.Println(mySlcEmpty)
	
	myArr := [3]int{1,2,3}
	fmt.Println(myArr)
	myArr2 := myArr
	fmt.Println(myArr2)
	myArr2[0] = 100
	fmt.Println(myArr2, " -- myArr2") // it change 
	fmt.Println(myArr, " -- myArr2") // it not change (pass By Value)

	mySlcChange := []int{1,2,3}
	fmt.Println(mySlcChange)
	mySlcChange2 := myArr
	fmt.Println(mySlcChange2)
	mySlcChange2[0] = 100
	fmt.Println(mySlcChange2) // it change
	fmt.Println(mySlcChange) // it change (pass By Reference)
}