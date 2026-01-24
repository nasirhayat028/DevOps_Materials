package main

import (
	"fmt"
)

func multi(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

func main() {
	sum, prod := multi(4, 5)
	fmt.Println("Sum is:", sum)
	fmt.Println("Product is:", prod)
}
