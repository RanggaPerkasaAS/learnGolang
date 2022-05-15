package main

import "fmt"

func main() {
	var nama [4]string
	nama[0] = "Rangga"
	nama[1] = "Perkasa"
	nama[2] = "Aprili"
	nama[3] = "hallo"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	var nilai = [...]int{
		80,
		90,
		30,
	}

	fmt.Println(nilai[1])
	fmt.Println(len(nilai))
}
