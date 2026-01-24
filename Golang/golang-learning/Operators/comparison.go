package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Println("a == b:", a == b) // false

	fmt.Println("a != b:", a != b) // true

	fmt.Println("a < b:", a < b) // true

	fmt.Println("a <= b:", a <= b) // true

	fmt.Println("a > b:", a > b) // false

	fmt.Println("a >= b:", a >= b) // false
}
