package main

import "fmt"

func main() {
	i := 10

	fmt.Println("Value of i:", i)
	fmt.Println("Address of i:", &i)
	fmt.Printf("%T %v", &i, &i)
	fmt.Printf("\n %T %v", *(&i), *(&i))
}
