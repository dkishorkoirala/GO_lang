/*
What is a variable


A variable in Go is a named storage location that holds a value. Think of it like a labeled box where you can store information that your program needs to remember.

To create variables, use the var keyword with explicit types:

var name string = "Gopher"  // Creates a string variable
var age int = 5            // Creates an integer variable
var price float64 = 19.99  // Creates a decimal number variable
After executing this code, we have three variables: a string called name, an integer called age, and a float (decimal number) called price.

To create variables using var and to let Go figure out the type for us we can write without the types:

var name = "Gopher"  // Go knows this is a string
var age = 5          // Go knows this is an integer
var price = 19.99    // Go knows this is a float64
After executing this code, we get the same result as before, but Go automatically determined the types based on the values.

Challenge

In this challenge, you'll practice declaring variables using the var keyword in Go. Your task is to declare two variables: name as a string with the value "Gopher" and age as an integer with the value 5. Then, the program will print these values.
*/

package main

import "fmt"

func main() {
    // TODO: Declare a variable named 'name' as a string with value "Gopher" using var
    var name string = "Gopher"
    
    // TODO: Declare a variable named 'age' as an int with value 5 using var
    var age int = 5
    
    // Don't modify the code below
    fmt.Println("My name is", name, "and I am", age, "years old.")
}