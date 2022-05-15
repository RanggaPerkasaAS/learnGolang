package main

import "fmt"

func main() {
	var nama = "Rangga"

	switch nama {
	case "Rangga":
		fmt.Println("Hallo Rangga")
	case "Perkasa":
		fmt.Println("Hallo Rangga")
	default:
		fmt.Println("Kenalan dong")
	}

	//SHORT STATMENT GOLANG
	// switch length := len(nama); length > 8 {
	// case true:
	// 	fmt.Println("Selamat")
	// case false:
	// 	fmt.Println("huuuuu")
	// }

	//SWITCH TNPA KONDISI

	var length = len(nama)
	switch {
	case length > 8:
		fmt.Println("Panjang")
	case length == 5:
		fmt.Println("Pas")
	default:
		fmt.Println("ndak pas")
	}
}
