/*Dereferencing Pointers


Dereferencing pointers (*) lets you access the value stored at a memory address. This is called "dereferencing" the pointer.

Create a variable and a pointer to it:

func main() {
    number := 42
    pointer := &number
    
    // Dereference the pointer to get the value
    value := *pointer
    
    fmt.Println("Pointer:", pointer)
    fmt.Println("Dereferenced value:", value)
}
The output shows both the pointer and the value it points to:

Pointer: 0xc000018030
Dereferenced value: 42
The asterisk (*) before a pointer variable accesses the actual value at that memory address.

Challenge

In this challenge, you'll practice dereferencing pointers in Go. We have a variable message with a string value, and a pointer messagePtr that points to it.

Your task is to dereference the pointer to access and print the value it points to.
*/
package main

import "fmt"

func main() {
	message := "Hello, pointers!"
	
	messagePtr := &message
	
	value := *messagePtr
	
	fmt.Println("The pointer points to the value:", value)
}