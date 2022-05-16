package main

import "fmt"

func main() {
	var counter = 0

	var increment = func()  {
		fmt.Println("increment")
		counter++
	}

	increment()

	fmt.Println(counter)
}