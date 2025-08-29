/*Go Comments


Comments in Go are notes that the compiler ignores. They help explain your code to other developers (and yourself).

First, let's look at single-line comments in Go. We use two forward slashes // to create a single-line comment:

// This is a single-line comment
fmt.Println("Hello, World!")
After executing this code, "Hello, World!" will show on the screen. The comment is completely ignored by Go when running the program.

Next, let's see how to write multi-line comments. We use /* to start and */ to end a multi-line comment:

/* This is a 
//    multi-line comment */
// fmt.Println("Hello, World!")
// After executing this code, we get the same result - "Hello, World!" will show on the screen. The multi-line comment is also ignored during execution.

// Challenge

// Beginner
// In this challenge, you'll practice adding comments to a Go program. Comments help document your code and make it easier for others to understand.

// Your task is to add appropriate comments to the code below:

// Add a single-line comment at the top of the main function explaining that this program displays a welcome message
// Add a multi-line comment before the print statement explaining that we're using fmt.Println to display text to the console*/


package main

import "fmt"
//calling main function
func main() {
    // TODO: Add a single-line comment here explaining that this program displays a welcome message
    
    // TODO: Add a multi-line comment here explaining that we're using fmt.Println to display text
    /*running the print
	statement in the function main*/
    fmt.Println("Welcome to Go programming!")
}