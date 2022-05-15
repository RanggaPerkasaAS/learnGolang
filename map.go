package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":   " Rangga",
		"Alamat": "Tulungagung",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["Alamat"])

	//menambah data map
	person["Tittle"] = "Proggramer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["Alamat"])
	fmt.Println(person["Tittle"])

	//function map
	fmt.Println(len(person)) //untuk mencari jumlah data
	delete(person, "Tittle")
	fmt.Println(person)
}
