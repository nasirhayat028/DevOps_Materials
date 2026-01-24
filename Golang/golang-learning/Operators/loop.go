package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println("Value of i:", i)
		fmt.Println("Square of i:", i*i)
	}

	for n := 100; n >= 90; n-- {
		fmt.Println("Value of n:", n)
		if n == 95 {
			break
		}
	}
}
