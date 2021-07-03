package main

import "fmt"

func main() {
	myMap := map[string]int {
		"Ahmet": 40,
		"Ayşe": 17,
		"Selim": 14,
		"Mustafa": 70,
	}
	fmt.Println(myMap)
	fmt.Println(myMap["Ahmet"], myMap["Selim"])

	isMarried := map[string]bool {
		"Ahmet": true,
		"Ayşe": false,
		"Selim": false,
		"Mustafa": false,
	}
	fmt.Println(isMarried)

	myMap2 := make(map[string]int)

	fmt.Println(myMap2)

	studentGrades := map[string]int {
		"Arin": 80,
		"Ahmet": 29,
		"Selim": 72,
		"Ayşe": 77,
		"Çınar": 0,
	}

	_, ok := studentGrades["Elis"]
	fmt.Println(ok)
	_, ok = studentGrades["arin"]
	fmt.Println(ok)

	fmt.Println(studentGrades)
	studentGrades["Mahmut"] = 55
	fmt.Println(studentGrades)

	delete(studentGrades, "Selim")
	fmt.Println(studentGrades)

	sg := studentGrades // pass by reference
	fmt.Println(len(studentGrades))
	fmt.Println(len(sg))
	delete(sg, "Arin")
	fmt.Println(len(studentGrades))
	fmt.Println(len(sg))


	for k, v := range studentGrades{
		fmt.Println(k,v)
	}

	myMap3 := map[string][]string {
		"Yaşar Kemal": []string{"Ince Mehmet", "Yer Demir Gök Bakır"},
		"Sebahattin Ali": []string{"Kuyucaklı Yusuf", "Kürk Mantolu Madonna", "Değirmen"},
	}
	fmt.Println(myMap3)
	fmt.Println(myMap3["Yaşar Kemal"])

	for key, value := range myMap3{
		fmt.Println("Writer: ", key)
		for i,v := range value{
			fmt.Println("\t", i+1, "book : ", v)
		}
	}
}