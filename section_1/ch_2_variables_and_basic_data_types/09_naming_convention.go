/*
Naming Conventions


Let's learn about naming conventions in Go, which are important patterns that make our code easier to read.

First, let's see how to name variables using camelCase (starting with lowercase and capitalizing subsequent words):

var userName = "Alex"
const maxLoginAttempts = 3
    
fmt.Println(userName)
fmt.Println(maxLoginAttempts)
After executing this code, we'll see our variable values displayed on the screen:

Alex
3
In Go, we follow these simple naming rules:

Use camelCase for variables and functions (like userName)
Avoid using snake_case (like user_name)
Don't use PascalCase (like UserName) for local variables
Important: Names that start with an uppercase letter (like UserName) are exported and visible outside packages
Following these naming conventions makes your Go code more readable and consistent with standard Go practices.

Challenge
In this challenge, you'll practice using proper Go naming conventions. You need to fix the variable names in the code to follow Go's standard naming conventions:

Use camelCase for variable names (start with lowercase, capitalize subsequent words)
Use PascalCase for exported variables (start with uppercase)
Avoid using underscores in variable names
Fix the variable names in the code to follow these conventions, then print them.
*/
package main

import "fmt"

func main(){
	var userName string = "John"
	var userAge int = 25
	var isActive bool = true
	var firstLogin string = "2023-01-15"

	fmt.Println("User:", userName)
	fmt.Println("Age:", userAge)
	fmt.Println("Active:", isActive)
	fmt.Println("First Login:", firstLogin)
}

