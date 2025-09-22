/*Creating Struct Instances


After defining a struct type, you need to create instances (values) of that type to use in your program.

Define a Person struct:

type Person struct {
    name string
    age  int
}

func main() {
    // Create a Person instance
    person1 := Person{name: "Alice", age: 30}
    
    fmt.Println(person1)
}
The output shows the struct instance with its field values:

{Alice 30}
You can also create a struct with zero values by omitting fields:

person2 := Person{}  // Creates {name:"", age:0}

Challenge

In this challenge, you'll practice creating instances of a struct. A Person struct has already been defined with fields for name, age, and isStudent.

Your task is to create a new instance of the Person struct with the following information:

Name: "Alice"
Age: 25
IsStudent: true
Then the program will print out the person's information.*/
package main

import "fmt"

// Person struct definition
type Person struct {
	name      string
	age       int
	isStudent bool
}

func main() {
	// TODO: Create a new Person struct instance with
	// name: "Alice", age: 25, isStudent: true
	var person Person
	person.name = "Alice"
	person.age = 25
	person.isStudent = true
	
	// Don't modify the code below
	fmt.Printf("Name: %s\n", person.name)
	fmt.Printf("Age: %d\n", person.age)
	fmt.Printf("Is Student: %t\n", person.isStudent)
}