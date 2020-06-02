package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 2
	fmt.Println(x[1])

	var strarr [5]string
	fmt.Println(strarr[0])

	// iterate through array
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Println(i, v)
	}

}
