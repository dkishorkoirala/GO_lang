/*
Strings


Strings in Go are used to store text data. Let's learn how to work with them step by step.

First, let's create a string variable using the var keyword with an explicit type declaration:

var message string = "Hello, Go!"
fmt.Println(message)
After executing this code, the text Hello, Go! will show on the screen. We've created a string variable named message and displayed its value.

Now, let's create a string using Go's short declaration syntax with :=, which automatically infers the type:

greeting := "Welcome to Go programming"
fmt.Println(greeting)
After executing this code, the text Welcome to Go programming will show on the screen. The := operator creates a new variable and determines its type based on the value provided.

Both methods accomplish the same thing - they create string variables that store text. The short declaration syntax (:=) is more commonly used in Go because it's concise and convenient.

Challenge

In this challenge, you'll practice working with strings in Go.

We've already set up two string variables: firstName and lastName. Your task is to:

Create a third string variable called greeting with the value "Hello, my name is "
Print each of the three strings on separate lines using fmt.Println()
*/

package main
import "fmt"

func main(){
	firstName := "Alex"
	lastName := "Smith"

	greeting := "Hello, my name is "

	fmt.Println(greeting)
	fmt.Println(firstName)
	fmt.Println(lastName)
}