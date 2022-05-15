package main

import "fmt"

func main() {
	var nilai32 int32 = 100
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	//byte ke string
	nama := "Rangga"
	tes := nama[1]           //byte
	tesString := string(tes) //string

	fmt.Println(tes)
	fmt.Println(tesString)

}
