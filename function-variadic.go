package main

import "fmt"

func sum(angka...int) int {
	var total =  0

	for _, value := range angka {
		total += value
	}

	return total
}

func main()  {
	var total = sum(10,10,10,10)
	fmt.Println(total)
}