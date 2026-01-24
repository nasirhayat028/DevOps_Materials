package main

import (
	"fmt"
)

func main() {
	// name:= "Nasir",

	var name string = "nasir"
	var welcomeMessage string = "Welcome to Go programming"
	var age int = 20

	fmt.Println("Hello, "+name, "!", "\n", welcomeMessage)
	fmt.Println("Your age is:", age)

	fmt.Printf("Type of name variable is: %v", name+"\n")
}
