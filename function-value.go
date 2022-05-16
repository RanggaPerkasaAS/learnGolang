package main

import "fmt"

func getGoodBye(name string) string {
	return "good bye" + name
}

func getTahun(tahun int) (string, int){
	if tahun % 400 == 0 {
		return "tahun kabisat" , tahun
	}else{
		return "tahun bukan kabisat" , tahun
	}
}

func main() {
	//Merubah function getgoodbye menjadi value yang disimpaan pada variable

	var goodbye = getGoodBye
	var result = goodbye("Rangga")
	fmt.Println(result)

	var cekTahun = getTahun
	var hasil, kata = cekTahun(2000)
	fmt.Println(hasil,kata)
}