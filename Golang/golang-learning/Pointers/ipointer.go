package main

import "fmt"

func main() {
	i := 10

	var ptr *int = &i

	fmt.Println("Value of i:", i)
	fmt.Println("Address of i:", ptr)

	name := "Nasir"

	var n *string = &name

	fmt.Println("Value of name:", name)
	fmt.Println("Address of name:", n)

}
