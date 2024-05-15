package main

import "fmt"

func main() {
	//var namaVariable tipeData
	var name1 string

	name1 = "ananda"

	//var namaVariable deklarasi
	var name2 = "dihasananda"

	fmt.Println(name1, name2)

	//namaVariable :=
	name := "dihas ananda"
	fmt.Println(name)

	name = "dihas"
	fmt.Println(name)

	//jadi satu
	var (
		firstName = "dihas"
		lastName  = "ananda"
	)
	fmt.Println(firstName, lastName)
}
