/*Using Pointers in Functions


Functions can accept and modify variables through pointers. This allows functions to change the original value, not just a copy.

Create a function that modifies a value through a pointer:

func main() {
    number := 10
    fmt.Println("Before:", number)
    
    // Pass the address of number
    doubleValue(&number)
    
    fmt.Println("After:", number)
}

func doubleValue(num *int) {
    // Modify the value at the pointer's address
    *num = *num * 2
}
The output shows the value changed by the function:

Before: 10
After: 20
When we pass &number, the function receives a pointer and can modify the original variable.

Challenge

In this challenge, you'll practice using pointers in functions. You need to complete a function that takes a pointer to a string and modifies the string value to be uppercase.

The function makeUppercase is already defined for you. Your task is to modify the string that the pointer points to by using the strings.ToUpper() function.*/

package main

import (
	"fmt"
	"strings"
)

// makeUppercase takes a pointer to a string and changes the string to uppercase
func makeUppercase(strPtr *string) {
	// TODO: Use strings.ToUpper() to change the string that strPtr points to
	// Hint: You need to dereference the pointer to access the string value
	*strPtr = strings.ToUpper(*strPtr)
	
}

func main() {
	// We already have a string variable
	message := "hello, world"
	
	// Print the original message
	fmt.Printf("Original message: %s\n", message)
	
	// Call makeUppercase with the address of message
	makeUppercase(&message)
	
	// Print the modified message
	fmt.Printf("Modified message: %s\n", message)
}