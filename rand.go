package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math/rand"
	"time"
)

func main(){
	target:= numRand(1,100)
	fmt.Println(" Find number between 1 and 100 ")
	reader := bufio.NewReader(os.Stdin)

	for attemps := 0; attemps < 10; attemps++ {
		fmt.Println("You have ", 10-attemps, " left")
		fmt.Println(" Please enter a guess: ")
		
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println(err)
		}

		if num > target {
			fmt.Println(" Enter smaller number ")
		} else if num < target {
			fmt.Println(" Enter bigger number ")
		} else {
			fmt.Println(" True guess", attemps, "you find ")
			break
		}
	}
	
}

func numRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}