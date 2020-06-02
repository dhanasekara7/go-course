package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := strings.Compare("12", "22")
	fmt.Println(res)

	res = strings.Compare("AB", "ab")
	fmt.Println(res)
	// A --> ascii 65
	// a --> ascii 97

	res_contains := strings.Contains("abcd", "cd")
	fmt.Println(res_contains)

	res_prefix := strings.Contains("abcd", "ab")
	fmt.Println(res_prefix)

	str_index := strings.Index("this is", "is")
	fmt.Println(str_index)

	// strings are immutable
	// replace will return new string
	res_replace := strings.Replace("abcd", "ab", "AB", 1)
	fmt.Println(res_replace)

	//ToLower
	//ToUpper
	//TrimSpace
	// all the above will return new string

	//Strconv package
	// Atoi(s) --> converts string to int
	inte, ok := strconv.Atoi("32")
	fmt.Println(inte, ok)

	str := strconv.Itoa(32)
	fmt.Println(str)

	// FormatFloat --> convert floating point number to string
	// ParseFloat --> converts string to floating point number

}
