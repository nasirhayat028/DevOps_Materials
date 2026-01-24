package main

import "fmt"

func main() {

	var a, b int = 10, 20

	fmt.Println("a && b:", a != 0 && b != 0) // true
	fmt.Println("a && b:", a < b && b > 15)  //true
	fmt.Println("a && b:", b < a && b == 20) //false

	fmt.Println("a || b:", a != 0 || b == 0) // true
	fmt.Println("a || b:", a == 0 || b > 15) // true
	fmt.Println("a || b:", a > 30 || a > b)  // false

	fmt.Println("! (a < b):", !(a < b)) // false
	fmt.Println("! (a > b):", !(a > b)) // true
	//reverse the res..
}
