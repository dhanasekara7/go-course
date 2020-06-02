package main

func main() {
	// Iota
	// as long as const values are not important

	//days of week
	//months of year

	type Grades int

	const (
		A Grades = iota
		B
		C
		D
	)

}
