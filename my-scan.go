package main

import "fmt"

func main() {
	var appleNum int

	fmt.Printf("Number of apps?")
	num, err := fmt.Scan(&appleNum)

	fmt.Printf(appleNum)
}
