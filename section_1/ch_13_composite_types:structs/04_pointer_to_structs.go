/*Pointers to Structs


When working with structs, you can create pointers to them just like with other types. This is useful for passing large structs efficiently.

Define a Person struct and create a pointer to it:

type Person struct {
    name string
    age  int
}

func main() {
    person := Person{name: "Alice", age: 30}
    
    // Create a pointer to the struct
    personPtr := &person
    
    // Access fields through pointer
    fmt.Println(personPtr.name)
    fmt.Println(personPtr.age)
}
The output shows we can access fields through the pointer:

Alice
30
You can also modify fields through the pointer:

personPtr.age = 31
fmt.Println(person.age)  // Prints: 31

Challenge

In this challenge, you'll practice working with pointers to structs in Go. You have a Person struct with name and age fields. Your task is to update the person's information using a pointer to the struct.

The code already creates a Person struct and a pointer to it. You need to modify the person's information through the pointer.*/

package main

import "fmt"

// Person struct with name and age fields
type Person struct {
	name string
	age  int
}

func main() {
	// Create a Person struct
	person := Person{
		name: "John",
		age:  25,
	}

	// Create a pointer to the Person struct
	personPtr := &person

	// Print the original person information
	fmt.Println("Original person:")
	fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)

	// TODO: Use the pointer to change the person's name to "Jane" and age to 30
	// Hint: You can access struct fields through a pointer using the dot notation
	personPtr.name = "Jane"
	personPtr.age = 30

	// Print the updated person information
	fmt.Println("Updated person:")
	fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)
}