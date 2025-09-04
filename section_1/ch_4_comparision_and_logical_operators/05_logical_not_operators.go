/*
Logical NOT Operator


The logical NOT operator (!) reverses a boolean value. It turns true into false and false into true.

Create a program that uses the NOT operator to check if a number is not even:

num := 7
isEven := num%2 == 0
isOdd := !isEven
fmt.Println("Is", num, "odd?", isOdd)
This outputs:

Is 7 odd? true
The NOT operator inverted isEven (which was false for 7) to true, correctly identifying that 7 is odd.

Challenge

In this challenge, you'll practice using the logical NOT operator (!) in Go.

The logical NOT operator inverts a boolean value: true becomes false and false becomes true.

We've provided a boolean variable isRaining. Your task is to use the NOT operator to invert its value and store the result in the result variable.
*/
package main
import "fmt"

func main(){
	isRaining := true
	result := !isRaining

	fmt.Println("Original value:", isRaining)
	fmt.Println("After using NOT operator:", result)
}