/*
Type Inference with `:=`


Let's explore a convenient way to create and initialize values in Go using the := shorthand. This is called type inference, where Go automatically figures out what type your data should be.

To create a text message, use :=:

message := "Hello, Go!"
fmt.Println(message)
After running this code, you'll see Hello, Go! printed to the console. Go automatically knew that message should be a string type.

Let's create a number using :=:

age := 25
fmt.Println(age)
This prints 25 to the console. Go automatically made age an integer type.

Let's create a decimal number:

price := 19.99
fmt.Println(price)
This prints 19.99 to the console. Go automatically made price a floating-point type.

Challenge

In this challenge, you'll practice using the := operator for type inference in Go.

The program already has some variables declared using the var keyword with explicit types. Your task is to declare two new variables using the := shorthand syntax that lets Go infer the types automatically.

Complete the TODOs in the code to make the program print the correct values.
*/
package main

import ("fmt")

func main(){
	// Variables declared with explicit types
    var name string = "Gopher"
    var age int = 5
    
    // TODO: Declare a variable called 'language' with the value "Go" using := syntax
    language := "Go"
    
    // TODO: Declare a variable called 'version' with the value 1.18 using := syntax
    version := 1.18
    
    // Print all variables
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Language:", language)
    fmt.Println("Version:", version)
}