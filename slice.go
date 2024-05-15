package main

import "fmt"

func main() {
	names := [...]string{"dihas", "Dihas", "ananda", "Ananda"}

	slice1 := names[1:3]
	fmt.Println(slice1)

	slice2 := names[:2]
	fmt.Println(slice2)

	slice3 := names[1:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)
}
