/*
Comparison Operators - Part 2


Go provides additional comparison operators for comparing numerical values.

Use the greater than operator (>) to check if one value is larger than another:

a := 10
b := 5
result := a > b
fmt.Println(result)
This outputs:

true
Other comparison operators include less than (<), greater than or equal to (>=), and less than or equal to (<=):

x := 7
y := 7
fmt.Println(x < y)  // false
fmt.Println(x >= y) // true
fmt.Println(x <= y) // true

Challenge
In this challenge, you'll practice using comparison operators in Go. We have defined several variables with different values. Your task is to complete the missing comparison expressions to make the program print the expected output.

You need to use the following comparison operators: < (less than), > (greater than), <= (less than or equal to), and >= (greater than or equal to).
*/
package main

import "fmt"

func main(){
	age1 := 25
	age2 := 30
	age3 := 25

	isAge1LessThanAge2 := age1 < age2
	isAge2GreaterThanAge1 := age2 > age1
	isAge1LessThanOrEqualToAge3 := age1 <= age3
	isAge2GreaterThanOrEqualToAge3 := age2 >= age3

	fmt.Println("Is age1 less than age2?", isAge1LessThanAge2)
	fmt.Println("Is age2 greater than age1?", isAge2GreaterThanAge1)
	fmt.Println("Is age1 less than or equal to age3?", isAge1LessThanOrEqualToAge3)
	fmt.Println("Is age2 greater than or equal to age3?", isAge2GreaterThanOrEqualToAge3)
}