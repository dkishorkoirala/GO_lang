/*
What is a Pointer?


A pointer is a variable that stores the memory address of another variable. Pointers help us access or modify variables indirectly.

Declaring a pointer variable: Use the symbol before the type to declare a variable that can hold a pointer:

var pointerToInt 

int    // Declares a pointer variable that can point to an int
var pointerToString string // Declares a pointer variable that can point to a string
Creating a pointer to an existing variable: Use the & (address-of) operator to get the memory address of a variable:

x := 10
pointerToX := &x  // pointerToX now holds the memory address of x
fmt.Println(pointerToX)
Accessing the value a pointer points to: Use the (dereference) operator:

fmt.Println(*pointerToX)  // Prints the value stored at the address
The output will be:

0xc000018030  // Memory address (your address will differ)
10           // The value at that address
*/