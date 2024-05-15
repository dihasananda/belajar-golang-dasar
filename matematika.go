package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b
	var g = a % b

	fmt.Println(c, d, e, f, g)

	var i = 10
	i += 10
	fmt.Println(i)
	i -= 10
	fmt.Println(i)
	i *= 10
	fmt.Println(i)
	i /= 10
	fmt.Println(i)
	i %= 10
	fmt.Println(i)

	var j = 1
	j++
	fmt.Println(j)
	j--
	fmt.Println(j)
	fmt.Println(-j)
	fmt.Println(+j)
}
