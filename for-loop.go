package main

import "fmt"

func main() {

	for hitung := 0; hitung <= 10; hitung++ {
		if hitung%2 == 0 {
			fmt.Println(hitung, "genap")
		} else {
			fmt.Println(hitung, "Ganjil")
		}
		// fmt.Println("Hitungan Ke:",  hitung)
	}

	var slice = []string{"Rangga", "Perkasa", "Aprili", "Sugiyanto"}

	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i])
	}

	//for range
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}
}
