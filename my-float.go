package main

import (
	"fmt"
)

func main() {
	// float32 --> upto 6 digits of precision
	// float64 --> upto 15 digits of precision

	var x float64 = 123.45
	var y float64 = 1.2345e2

	var z complex128 = complex(2, 3)

	fmt.Println(x, y, z)

}
