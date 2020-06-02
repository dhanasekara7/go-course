package main

import "fmt"

func main() {
	x := 10
	switch {
	case x <= 10:
		fmt.Println("x<=10")
	case x > 10:
		fmt.Println("x > 10")
	}
}
