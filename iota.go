package main

import "fmt"

const (
	_ = iota // ignore
	KB = 1 << (10 * iota) // (10 - iota) times 2
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/GB)
}