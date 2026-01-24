package main

import "fmt"

func main() {

	arr := []int{10, 20, 30, 40, 50}

	slice1 := arr[1:4] // Slicing from index 1 to 3
	slice2 := arr[:3]  // Slicing from start to index 2
	slice3 := arr[2:]  // Slicing from index 2 to end
	slice4 := arr[:]   // Slicing the entire array

	fmt.Println("Original array:", arr)
	fmt.Println("Slice from index 1 to 3:", slice1)
	fmt.Println("Slice from start to index 2:", slice2)
	fmt.Println("Slice from index 2 to end:", slice3)
	fmt.Println("Slice of the entire array:", slice4)
}
