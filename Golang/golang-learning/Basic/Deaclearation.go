package main

import "fmt"

func main() {
	var r int
	var n string

	n = "Nasir"
	r = 38

	fmt.Println(n, r)
	fmt.Printf("name: %s \nage: %d", n, r)

	var (
		secondName = "Hayat"
		age        = 19
	)
	fmt.Println("\n", secondName, age)

	var a, b string = "Razzaq", "khan"

	fmt.Println(a, b)
}
