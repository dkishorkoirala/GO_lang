/*
Division Operator


The division operator (/) divides the left operand by the right operand.

Create two variables and perform division:

a := 10
b := 3
result := a / b
fmt.Println(result)
Output:

3
Note that when dividing integers, Go truncates the result (removes decimal part).

For decimal division, use floating-point numbers:

x := 10.0
y := 3.0
decimalResult := x / y
fmt.Println(decimalResult)
Output:

3.3333333333333335

Challenge

In this challenge, you'll practice using the division operator (/) in Go. We have two numbers already defined. Your task is to divide the first number by the second number and store the result in the provided variable.

Complete the code to perform the division and print the result.
*/

package main
import "fmt"

func main (){
	number1 := 20.0
	number2 := 4.0
	result := number1 / number2
	fmt.Println("The result of",number1,"divided by",number2,"is:", result)
}