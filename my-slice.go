package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	slice1 := arr[0:3]
	fmt.Println(len(slice1), cap(slice1))
	slice1 = append(slice1, 41, 51, 61)
	for i, k := range slice1 {
		fmt.Println(i, k)
	}
	fmt.Println(len(slice1), cap(slice1))
}
