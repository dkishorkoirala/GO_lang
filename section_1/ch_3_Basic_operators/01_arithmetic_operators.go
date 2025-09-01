/*
Arithmetic Operators


Go provides standard arithmetic operators for mathematical operations.

Create variables and perform addition with +:

a := 5
b := 3
sum := a + b
fmt.Println(sum)
Output:

8
Perform subtraction with -:

difference := a - b
fmt.Println(difference)
Output:

2
Perform multiplication with *:

product := a * b
fmt.Println(product)
Output:

15

Challenge

In this challenge, you'll practice using arithmetic operators in Go. We have three predefined variables: num1, num2, and result.

Your task is to calculate the sum of num1 and num2, then multiply that result by 2, and store the final answer in the result variable.*/

package main
import "fmt"

func main(){
	num1 := 5
	num2 := 7
	result := 0
	result = (num1 + num2) * 2
	fmt.Println("The result is:", result)
}