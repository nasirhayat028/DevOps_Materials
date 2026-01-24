package main

import (
	"fmt"
	"reflect"
)

func main() {
	Country := "Pakistan"
	{
		City := "Karachi"
		fmt.Println("City:", City)
		fmt.Println("Country:", Country)
	}
	fmt.Println("Country outside inner block:", Country)
	// fmt.Println("City outside inner block:", City) // This will cause an error

	fmt.Println("Type of Country variable is:", reflect.TypeOf(Country)) // it will find the type of Country variable
	// fmt.Println("Type of City variable is:", reflect.TypeOf(City)) // This will cause an error
}
