/*Accessing Struct Fields


To access fields in a struct, use the dot notation (structVariable.fieldName).

Create a Person struct and access its fields:

type Person struct {
    name string
    age  int
}

func main() {
    person := Person{name: "Alice", age: 30}
    
    // Access individual fields
    fmt.Println(person.name)
    fmt.Println(person.age)
}
The output shows the individual field values:

Alice
30
You can also modify field values using the same dot notation:

person.age = 31
fmt.Println(person.age)  // Prints: 31

Challenge

In this challenge, you'll practice accessing fields from a struct. We've defined a Person struct with name, age, and isStudent fields.

Your task is to access these fields and print them using the correct format verbs.*/

package main

import "fmt"

func main() {
	// Person struct with three fields
	type Person struct {
		name      string
		age       int
		isStudent bool
	}

	// A sample person
	john := Person{
		name:      "John Doe",
		age:       25,
		isStudent: true,
	}

	// TODO: Print john's name using fmt.Printf and the %s format verb
	fmt.Printf("Name: %s\n", john.name)
	
	// TODO: Print john's age using fmt.Printf and the %d format verb
	fmt.Printf("Age: %d\n", john.age)
	
	// TODO: Print whether john is a student using fmt.Printf and the %t format verb
	fmt.Printf("Is student: %t\n", john.isStudent)
	
}