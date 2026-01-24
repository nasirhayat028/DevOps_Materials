package main

import "fmt"

func main() {
	type Person struct {
		name string
		age  int
	}

	p := Person{name: "Alice", age: 30}

	fmt.Println("Person Name:", p.name)
	fmt.Println("Person Age:", p.age)

	type table struct {
		name   string
		rollno int
		grade  string
		pass   bool
	}

	data := table{name: "NAsir", rollno: 38, grade: "A+", pass: true}

	fmt.Println("Table:", data)
}
