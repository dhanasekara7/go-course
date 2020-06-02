package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, ok := http.Get("http://www.google.com")
	fmt.Println(res, ok)
}
