package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai16, nilai32, nilai64)

	var name = "dihas ananda"
	var e = name[0]
	var eString = string(e)

	fmt.Println(name, e, eString)
}
