/*
Initializing Variables


Let's learn how to initialize variables in Go. Initializing means giving a variable a starting value when you create it.

First, let's see how to initialize variables using the var keyword:

var age int = 25
var name string = "Gopher"
var isActive bool = true
    
fmt.Println(age, name, isActive)
After executing this code, the values of our variables will show on the screen:

25 Gopher true
We can also use Go's short declaration operator (:=) to make our code more concise:

age := 25
name := "Gopher"
isActive := true
    
fmt.Println(age, name, isActive)
After executing this code, we get the same result as before:

25 Gopher true
The short declaration operator (:=) automatically figures out the type based on the value you provide, which makes your code cleaner and easier to read.

Challenge

In this challenge, you'll practice initializing variables in Go. Your task is to initialize three variables with different data types and then print them.

The program already has the variable declarations, but you need to initialize them with values.
*/

package main
import "fmt"

func main(){
	var name string = "John"
	var age int = 25
	var isStudent bool = true

	fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Is Student:", isStudent)
}