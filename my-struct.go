package main

import "fmt"

type Person struct {
	name string
	phno string
}

func main() {
	p1 := new(Person)
	p1.name = "ddd"
	p1.phno = "12345"
	fmt.Println(p1)

	p2 := Person{name: "kkk", phno: "54321"}
	fmt.Println(p2)

}
