package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dat, e := ioutil.ReadFile("go-course/my-ioutil.go")
	if e != nil {
		fmt.Println(dat)
	}

	err2 := ioutil.WriteFile("go-course/zfile.txt", dat, 0777)
	fmt.Println(err2)
}
