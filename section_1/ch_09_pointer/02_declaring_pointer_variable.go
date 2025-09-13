/*Declaring Pointer Variables


In Go, you can declare pointers in two main ways: using var with a pointer type or using the address operator with :=.

Declare a pointer using var with a pointer type:

func main() {
    var x int = 42
    var ptr *int    // Declare a pointer to an integer
    ptr = &x        // Assign the address of x to ptr
    
    fmt.Println("Value of x:", x)
    fmt.Println("Value of ptr:", ptr)
    fmt.Println("Value at ptr:", *ptr)
}
The output shows the pointer stores the address of x:

Value of x: 42
Value of ptr: 0xc000018030
Value at ptr: 42
You can also declare and initialize a pointer in one line using :=:

y := 100
ptr := &y    // Declare and initialize pointer in one step

Challenge
In this challenge, you'll practice declaring and using pointers in Go. A pointer is already set up for you, and your task is to create a new pointer that points to the same value.

Complete the TODO in the code to declare a new pointer variable called 'secondPointer' that points to the same memory address as the existing pointer.
*/
package main

import "fmt"

func main() {
	originalValue := 42
	pointerToValue := &originalValue
	
	secondPointer := pointerToValue
	
	fmt.Printf("Original value: %v\n", originalValue)
	fmt.Printf("Value through first pointer: %v\n", *pointerToValue)
	fmt.Printf("Value through second pointer: %v\n", *secondPointer)
	
	originalValue = 100
	fmt.Printf("\nAfter changing original value to 100:\n")
	fmt.Printf("Original value: %v\n", originalValue)
	fmt.Printf("Value through first pointer: %v\n", *pointerToValue)
	fmt.Printf("Value through second pointer: %v\n", *secondPointer)
}