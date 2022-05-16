package main

import "fmt"

//TYPE DECLARATION
type Filter func(string)string

func sayHello(name string, filter Filter) {
	var namedFIlter = filter(name)
	fmt.Println(namedFIlter)
}

func spamFilter(name string) string {
	if name == "Anjing" || name == "anjing" {
		return "....."
	}else if name == "cok" || name == "Cok"{
		return "..."
	}else{
		return name
	}
}

func main() {
	sayHello("Rangga",spamFilter)
	sayHello("anjing",spamFilter)
	sayHello("cok",spamFilter)
}