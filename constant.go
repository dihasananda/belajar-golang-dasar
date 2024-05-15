package main

import "fmt"

func main() {
	const firstName string = "dihas"
	const lastName = "ananda"

	fmt.Println(firstName, lastName)

	//firstName = "dihas"
	//lastName = "ananda"

	const (
		first = "dihas"
		last  = "ananda"
	)

	fmt.Println(first, last)
}
