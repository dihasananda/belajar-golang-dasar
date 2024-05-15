package main

import "fmt"

func main() {
	type NoKTP string
	var ktpDihas NoKTP = "11111111"
	var contoh string = "2222222"
	var contohKtp = NoKTP(contoh)

	fmt.Println(ktpDihas, contohKtp)
}
