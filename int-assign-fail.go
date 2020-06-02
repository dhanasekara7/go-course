package main

import "fmt"

func main() {
	var x int32 = 1
	var y int16 = 2
	//error when assign y to x :
	//cannot use y (type int16) as type int32 in assignment
	// x = y

	x = int32(y)
	fmt.Println(x, y)
}
