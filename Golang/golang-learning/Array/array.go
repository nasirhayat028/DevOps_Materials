package main

import "fmt"

func main() {

	var names [6]string = [6]string{"Nasir", "Hayat", "Razzaq", "Wajahat", "Sajid", "Basit"}

	fmt.Println(names)
	fmt.Println("Length of names array:", len(names))

	grades := [5]int{90, 85, 88, 92, 80}

	fmt.Println(grades)

	fmt.Println("First name:", names[0])
	fmt.Println("Second name:", names[1])

	fruits := [...]string{"Apple", "Banana", "Mango", "Orange", "Pineapple"}

	fmt.Println("Fruits:", fruits)

}
