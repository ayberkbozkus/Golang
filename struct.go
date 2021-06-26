package main

import "fmt"

type worker struct { // underlying type
	name string
	age int
	isMarried bool
}

var employee struct { 
	name string // field
	age int
	isMarried bool
}

type manager struct { 
	worker
	hasDegree bool
}

func main(){

	fmt.Println(employee)
	fmt.Printf("%#v \n", employee)
	fmt.Println(employee.age)

	employee.name = "Gurcan"
	employee.age = 40
	employee.isMarried = true

	fmt.Println(employee)
	fmt.Printf("%#v \n", employee)
	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.isMarried)

	var e1 worker
	e1.name = "Gürcan"
	e1.age = 40
	e1.isMarried = true

	var e2 worker
	e2.name = "Arin"
	e2.age = 5
	e2.isMarried = false

	fmt.Printf("%#v \n", e1)
	fmt.Printf("%#v \n", e2)

	e3 := worker{
		name : "Elis",
		age : 3,
		isMarried: false,
	}

	fmt.Println(e3)

	m1 := manager{
		worker: worker{
			name: "Ayşe",
			age: 28,
			isMarried: false,
		},
		hasDegree: true,
	}

	fmt.Println(m1)

	theBoss := struct{ // anonim struct
		name string
		money bool
	}{name: "THE BOSS", money: true}

	fmt.Println(theBoss)
}