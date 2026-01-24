package main

import "fmt"

func addition(a int, b int) int {
	add := a + b
	return add
}

func greeting(str string) string {
	fmt.Println("Hello:", str)
	return str
}

func main() {
	result := addition(5, 10)
	fmt.Println("Sum is:", result)
	greeting("Nasir")
}
