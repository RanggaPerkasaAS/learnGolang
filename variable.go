package main

import "fmt"

func main() {
	//DENGAN PENDEKLARASIAN TIPE DATA
	var name string

	name = "Rangga"
	fmt.Println(name)

	//TANPA PENDEKLARASIAN TIPE DATA
	name = "RanggaPAS"
	fmt.Println(name)

	var nama = "Rani Hanifah"
	fmt.Println(nama)

	var angka = 9
	fmt.Println(angka)

	//VAR BISA TIDAK DIDEKLARASIKAN TETAPI DIGANTI :=
	negara := "indonesia"
	fmt.Println(negara)

	negara = "Inggris"
	fmt.Println(negara)

	//MULTIPLE VARIABLE
	var (
		Firstname = "Rangga"
		Lastname  = "Perkasa"
	)
	fmt.Println(Firstname + " " + Lastname)
}
