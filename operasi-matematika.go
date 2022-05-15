package main

import "fmt"

func main() {
	var hasil = 8 + 8
	fmt.Println(hasil)

	var a = 10
	var b = 2
	var c = a % b
	fmt.Println(c)

	//AUGMENTED ASSIGMENT
	var g = 10
	g += 10
	fmt.Println(g)

	//unary assigment
	var i = 10
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
}
