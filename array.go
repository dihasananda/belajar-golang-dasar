package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "dihas"
	name[1] = "ananda"

	fmt.Println(name)
	fmt.Println(name[0], name[1])

	var values = [...]int{
		90, 80, 95,
	}
	values[2] = 100

	fmt.Println(values, values[0], len(values))
}
