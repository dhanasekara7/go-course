package main

import (
	"fmt"
	"os"
)

func main() {
	//file, _ := os.Create("go-course/outfile.txt")
	//barr := []byte{1, 2, 3}
	//nb, _ := file.Write(barr)
	//fmt.Printf("%d", nb)

	file, _ := os.Create("go-course/outfile.txt")
	//barr := []byte{1, 2, 3}
	nb, _ := file.WriteString("this is fun")
	fmt.Printf("%d", nb)

}
