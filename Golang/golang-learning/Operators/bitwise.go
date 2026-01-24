package main

import "fmt"

func main() {
	var a, b uint = 12, 5

	fmt.Println("a & b =", a&b)   // AND operator
	fmt.Println("a | b =", a|b)   // OR operator
	fmt.Println("a ^ b =", a^b)   // XOR operator
	fmt.Println("a &^ b =", a&^b) // AND NOT operator
	fmt.Println("a << 2 =", a<<2) // Left shift
	fmt.Println("b >> 1 =", b>>1) // Right shift
}
