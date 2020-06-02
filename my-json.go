package main

import (
	"encoding/json"
	"fmt"
)

type PersonSt struct {
	Name  string // field name should be in Capital
	Addr  string
	Phone string
}

func main() {
	p1 := PersonSt{Name: "ddd", Addr: "12/51F", Phone: "12345"}
	barr, err := json.Marshal(p1)
	fmt.Println(">>>", err)
	if err != nil {
		fmt.Println("?>>>", barr) // not printing :-(
	}

	var p2 PersonSt
	err2 := json.Unmarshal(barr, &p2)
	fmt.Println(p2, err2)
}
