package main

import "fmt"

func main() {

	var i int = 100

	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	case 3:
		fmt.Println("i is 3")
	default:
		fmt.Println("i is:", i)
	}
}
