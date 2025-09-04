/*
Operator Precedence Basics


Operator precedence determines the order in which operations are performed in an expression. Operations with higher precedence are performed before those with lower precedence.

Create a program that demonstrates operator precedence:

result1 := 5 + 3 * 2
result2 := (5 + 3) * 2
fmt.Println("5 + 3 * 2 =", result1)
fmt.Println("(5 + 3) * 2 =", result2)
This outputs:

5 + 3 * 2 = 11
(5 + 3) * 2 = 16
In the first expression, multiplication (*) has higher precedence than addition (+), so 3 * 2 is calculated first. In the second expression, parentheses override the default precedence, forcing 5 + 3 to be calculated first.

Challenge
In this challenge, you'll practice understanding operator precedence in Go. You have several variables with different values, and you need to fix the expression in the result variable to get the expected output.

The expression currently has incorrect parentheses placement. Your task is to add parentheses in the right places to make the expression evaluate to true.
*/

package main 
import "fmt"

func main(){
	a := true
	b := false
	c := true

	result := (a && b) || c

	fmt.Println(result)
}