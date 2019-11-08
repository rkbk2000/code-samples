package utils

import "fmt"

// Person type, to store name and age
type Person struct {
	Name string
	Age  int
}

// Print Method to print a person
func Print(person Person) {
	fmt.Printf("%.v", person)
}
