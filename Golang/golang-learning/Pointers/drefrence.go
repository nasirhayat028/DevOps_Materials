package main

import "fmt"

func main() {
	i := 10

	var ptr *int = &i

	*ptr = 100

	fmt.Println("Value of i:", i)
	fmt.Println("Address of i:", ptr)
	fmt.Printf("%T %v", ptr, ptr)
	fmt.Printf("\n %T %v", *ptr, *ptr)

}
