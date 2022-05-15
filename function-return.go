package main

import "fmt"

func getHello(nama string) string {
	var result = "Hallo" + nama
	return result
}

func getname(name string) string {
	if name == "" {
		return "Hallo bro"
	} else {
		return "Hallo" + name
	}
}
//MULTILE RETURN VALUE
func kabisat(tahun int) (int, string) {
	if tahun%400 == 0 {
		return tahun , "tahun Kabisat"
	} else {
		return tahun , "Bukan Kabisat"
	}
}

func getnama() (string,string)  {
	return "Rangga", "Perkasa"
}

//NAMED RETURN VALUE
func getFullname() (awal,tengah,akhir string)  {
	awal = "Rangga"
	tengah = "Perkasa"
	akhir = "Aprili"
	return
}

func main() {
	var hasil = getname("Eko")
	fmt.Println(hasil)
	//MULTILE RETURN VALUE
	var year, kata = kabisat(2000)
	fmt.Println(year, kata)

	var namaAwal, namaAkhir = getnama()
	fmt.Println(namaAwal , namaAkhir)

	var awal,tengah,akhir = getFullname()
	fmt.Println(awal, tengah, akhir)
}
