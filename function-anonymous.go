package main

import "fmt"

type Blacklist func(string) bool

func regusterUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked", name)
	}else{
		fmt.Println("Welcome" , name)
	}
}

func main() {
	var blacklist1 =  func(name string)bool{
		return name == "admin"
	}
	regusterUser("Rangga",blacklist1)
	regusterUser("admin",blacklist1)
}