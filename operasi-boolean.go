package main

import "fmt"

func main() {
	var (
		a = 2
		b = 3
	)
	var result bool = a > b
	var hasil bool = a < b

	var jawaban bool = result && hasil
	fmt.Println(jawaban)

	var (
		nilai   = 70
		absensi = 90
	)

	var lulus = (nilai >= 70 && absensi >= 70)
	// 	var lulusujian bool = nilai >= 70
	// 	var lulusabsensi bool = absensi >= 80

	// 	var lulus bool = lulusabsensi && lulusujian
	fmt.Println(lulus)
}
