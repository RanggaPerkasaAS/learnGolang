package main

import "fmt"

func sayHello() {
	fmt.Println("Hallo")
}

func hitung() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	sayHello()
	hitung()
}
