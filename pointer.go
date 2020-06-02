package main

import "fmt"

func main() {
	// new returns the pointer
	ptr := new(int)
	*ptr = 3
	fmt.Println(*ptr)
}
