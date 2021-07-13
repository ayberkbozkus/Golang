// Package clause
package main

// Import statement
import (
	"fmt"
	"strconv"
	"os"
	"strings"
	"bufio"

)

func main(){

	fmt.Print("Enter your grade: ")
	grade, _ := getGrade()
	var result string

	if grade >= 50 {
		result = "Pass"
	} else {
		result = "Failed"
	}

	fmt.Println(result)

}

func getGrade() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	input2 = strings.TrimSpace(input)
	num, err := strconv.Atoi(input2)

	if err != nil {
		fmt.Println(err)
	}

	return num, nil
}