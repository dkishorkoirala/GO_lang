/*
Printing Types


In Go, we can display the type of a variable using the %T format verb with the fmt.Printf() function. Let's explore how to do this with different data types.

First, let's create variables of different types and print their types:

name := "Gopher"
age := 25
height := 1.85
isActive := true
    
fmt.Printf("name is type: %T\n", name)
fmt.Printf("age is type: %T\n", age)
fmt.Printf("height is type: %T\n", height)
fmt.Printf("isActive is type: %T\n", isActive)
After executing this code, we'll see the type of each variable shown on the screen:

name is type: string
age is type: int
height is type: float64
isActive is type: bool
The %T format verb tells Go to display the data type instead of the value itself. This is very helpful when you're not sure what type a variable is or when you're learning about Go's type system.

Challenge

In this challenge, you'll practice using the %T format verb with fmt.Printf to display the types of different variables.

The code already has several variables of different types defined. Your task is to use fmt.Printf with the %T format verb to print out the type of each variable.
*/
package main 
import "fmt"

func main(){
	age := 25
	price := 19.99
	isAvailable := true
	productName := "Wireless Headphones"

	fmt.Printf("The type of age is: %T\n", age)
	fmt.Printf("The type of price is: %T\n", price)
	fmt.Printf("The type of isAvailable is: %T\n", isAvailable)
	fmt.Printf("The type of productName is: %T", productName)
}