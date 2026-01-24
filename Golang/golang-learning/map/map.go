package main

import "fmt"

func main() {

	my_map := map[string]string{
		"name":    "Nasir",
		"country": "Pakistan",
	}

	fmt.Println("Original map:", my_map)
	fmt.Println("Value for 'name':", my_map["name"])

	value, found := my_map["country"]

	fmt.Println("Value for 'country':", value, "Found:", found)

	// my_map2 := make(map[string]int("CGPA": 3, "age": 20))

	my_map2 := map[string]int{
		"CGPA": 3,
		"age":  20,
	}
	fmt.Println("New map:", my_map2)

	my_map3 := make(map[string]int)
	my_map3["height"] = 5
	my_map3["weight"] = 70

	fmt.Println("Third map:", my_map3)

	delete(my_map, "country")
	fmt.Println("Map after deleting 'country':", my_map)

}
