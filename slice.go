package main

import "fmt"

func main() {
	var bulan = [...]string{
		"Januari",
		"Februari",
		"maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = bulan[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// bulan[5] = "Kamu"
	// fmt.Println(slice1)

	// slice1[0] = "kamu kamu lagi"
	// fmt.Println(bulan)

	//SLICE APPEND
	//jika di array masih muat maka index tersebut akan di replace dengan yang baru 
	//jika array tidak muat maka akan dibuat array baru atau yang lam  tidak akan berubah
	var slice2 = bulan[4:6]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Rangga")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(bulan)

	//MAKE SLICE
	var newslice = make([]string, 2,5)
	newslice[0] = "Rangga"
	newslice[1] = "Perkasa"

	fmt.Println(newslice)

	//COPY SLICE
	var copyslice = make([]string, len(newslice), cap(newslice))
	copy(copyslice,newslice)
	fmt.Println(copyslice)
}