/*Defining Custom Types


In Go, you can create your own custom types based on existing types. This gives your code more clarity and type safety.

Define a custom type for temperature in Celsius:

type Celsius float64

func main() {
    var temperature Celsius = 25.5
    fmt.Printf("Temperature: %.1f°C\n", temperature)
}
The result shows the temperature value:

Temperature: 25.5°C
Custom types behave like their underlying types but are treated as distinct types by the compiler, improving code readability and preventing logical errors.

Challenge

In this challenge, you'll practice defining a custom type in Go. Custom types allow you to create your own data types based on existing ones.

Your task is to define a custom type called Temperature that is based on float64, and then create a variable of this type.

After defining the type, the program should print the value and type of the temperature variable.*/

package main

import "fmt"

type Temperature float64

func main() {
	// Create a variable of type Temperature with value 23.5
	var roomTemp Temperature = 23.5
	
	// Print the value and type of roomTemp
	fmt.Printf("Room temperature: %v\n", roomTemp)
	fmt.Printf("Type of roomTemp: %T\n", roomTemp)
}