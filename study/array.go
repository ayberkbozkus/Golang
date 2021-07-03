package main

import "fmt"

func main() {
	cities := []string{"istanbul", "roma", "tahran", "belgrad","",""}
	fmt.Println(cities)
	fmt.Println(cities[0])
	fmt.Println(len(cities[0]))

	var mArr [10]int
	fmt.Println(mArr)

	for i:=0; i< len(cities); i++ {
		fmt.Println(i, cities[i])
	}

	myArr := [10]int{0,1,2,3,4,5,6,7,8,9}

	myArr = mySquare(myArr)

	fmt.Println(myArr)

	for index, city := range cities {
		fmt.Println(index,city)
	}
	
}

func mySquare(arr[10]int) [10]int {
	for i := 0; i< len(arr); i++{
		arr[i] = arr[i] * arr[i]
	}
	return arr
}