package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("go-course/my-open-file.go")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err2 := file.Read(data)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("read %d bytes : %q\n", count, data[:count])
}
