/*
Recap - Input and Output
Challenge

In this challenge, you'll practice using fmt.Printf to display formatted output with different format specifiers.

Complete the code to display information about a Gopher using the appropriate format specifiers:

Use %s for strings
Use %d for integers
Use %.2f for floating-point numbers with 2 decimal places
Use %t for boolean values
Use %T to display the type of a variable
Make sure each output line ends with a newline character (\n).
*/

package main 
import "fmt"

func main(){
	name := "Gopher"
  	age := 5
  	height := 0.75
  	isActive := true

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d years\n", age)
	fmt.Printf("Height: %.2f meters\n", height)
	fmt.Printf("Active: %t\n", isActive)
	fmt.Printf("Name type: %T", name)
}