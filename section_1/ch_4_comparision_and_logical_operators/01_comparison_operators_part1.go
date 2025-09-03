/*
Comparison Operators - Part 1

Comparison operators compare two values and return a boolean result (true or false).

The equality operator (==) checks if two values are equal:

a := 5
b := 5
result := a == b
fmt.Println(result)
This outputs:

true
The inequality operator (!=) checks if two values are not equal:

x := 10
y := 20
notEqual := x != y
fmt.Println(notEqual)
This outputs:

true
Challenge

In this challenge, you'll practice using the equality (==) and inequality (!=) comparison operators in Go.

We have two variables with predefined values. Your task is to compare these variables using both operators and store the results in the provided boolean variables.

Then we'll print the results to see if your comparisons are correct.
*/
package main
import "fmt"

func main(){
	num1 := 10
	num2 := 10
	name1 := "Go"
	name2 := "Python"

	numbersEqual := num1 == num2
	namesNotEqual := name1 != name2
	fmt.Println("Are the numbers equal?", numbersEqual)
	fmt.Println("Are the names different?", namesNotEqual)
}