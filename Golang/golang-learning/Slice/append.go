package main

import "fmt"

func main() {

	arr := []int{10, 20, 30, 40, 50}

	slice1 := arr[1:4]

	fmt.Println("Original array:", arr)
	fmt.Println("Slice from index 1 to 3:", slice1)
	fmt.Println("Slice lrngth", len(slice1))
	fmt.Println("Slice capacity", cap(slice1))

	slice1 = append(slice1, 60, 70)
	fmt.Println("After appending 60 and 70 to slice1:", slice1)
	fmt.Println("Original array after appending to slice1:", arr)
	fmt.Println("Length after append", len(slice1))
	fmt.Println("Capacity after append", cap(slice1))
}
