package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["this"] = 1
	m["is"] = 2
	m["fun"] = 3

	str, presence := m["this"]
	fmt.Println(str, presence)

	str2, presence2 := m["this2"]
	fmt.Println(str2, presence2)
}
