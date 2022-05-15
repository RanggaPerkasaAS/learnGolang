package main

import "fmt"

func main() {
	var angka = 5

	if angka%2 == 0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}

	var kabisat = 2000

	if kabisat%400 == 0 {
		fmt.Println("Tahun Kabisat")
	} else {
		fmt.Println("Bukan Tahun Kabisat")
	}

	var nama = "Perkasa"

	if nama == "Rangga" {
		fmt.Println("Hallo")
	} else if nama == "Perkasa" {
		fmt.Println("Hai")
	} else {
		fmt.Println("kamu siapa")
	}

	//SHORT STATMENT
	if length := len(nama); length > 8 {
		fmt.Println("Terlalu oanjang")
	} else {
		fmt.Println("pendek")
	}
}
