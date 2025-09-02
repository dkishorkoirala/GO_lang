/*
The Modulo Operator


The modulo operator (%) returns the remainder after division of one number by another.

Calculate the remainder when dividing 10 by 3:

a := 10
b := 3
remainder := a % b
fmt.Println(remainder)
Output:

1
This is useful for determining if a number is even or odd:

number := 15
isEven := number % 2 == 0
fmt.Println("Is", number, "even?", isEven)
Output:

Is 15 even? false

Challenge

In this challenge, you'll practice using the modulo operator (%) in Go. The modulo operator returns the remainder after division.

We have two numbers already defined. Your task is to calculate the remainder when number is divided by divisor using the modulo operator, and store it in the remainder variable.
*/
package main
import "fmt"

func main(){
	number := 17
	divisor := 5
	var remainder int = number % divisor
	fmt.Println("The remainder when", number, "is divided by", divisor, "is:", remainder)
}