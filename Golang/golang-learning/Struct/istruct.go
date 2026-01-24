package main

import "fmt"

func main() {
	type me struct {
		name  string
		age   int
		class string
	}

	data := me{
		name:  "Nasir",
		age:   38,
		class: "BSCS",
	}
	fmt.Println("My Data:", data)
}
