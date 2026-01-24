package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var c int
	c = b  // assigning value of b to c
	a += 5 // a = a + 5
	b -= 5 // b = b - 5
	c *= 2 // c = c * 2 = 40
	c /= 4 // c = c / 4 = 10
	c %= 3 // c = c % 3 = 1

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("Value of c is b=c:", c)
}
