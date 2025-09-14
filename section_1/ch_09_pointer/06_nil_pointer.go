/*Nil Pointers


A nil pointer is a pointer that doesn't point to anything. It's the zero value for pointer types.

Declare a nil pointer:

var ptr *int
fmt.Println(ptr)
fmt.Println(ptr == nil)
The output shows the pointer is nil:

<nil>
true
Attempting to dereference a nil pointer causes a panic:

var ptr *int
// This will crash your program
// fmt.Println(*ptr)
Always check if a pointer is nil before dereferencing:

if ptr != nil {
    fmt.Println(*ptr)
}

Challenge

In this challenge, you'll work with nil pointers in Go. A nil pointer is a pointer that doesn't point to any memory address.

Your task is to check if a pointer is nil and print an appropriate message.

The code already has a pointer variable that is set to nil. You need to add code to check if the pointer is nil and print a message accordingly.
*/
package main

import "fmt"

func main() {
	var ptr *int
	
	// TODO: Check if ptr is nil
	// If it is nil, print "The pointer is nil"
	// If it is not nil, print "The pointer is not nil"
	if ptr == nil {
		fmt.Println("The pointer is nil")
	}else{
		fmt.Println("The pointer is not nil")
	}
	
	fmt.Println("Program completed")
}