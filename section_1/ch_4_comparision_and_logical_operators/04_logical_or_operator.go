/*Logical OR Operator


The logical OR operator (||) returns true when at least one condition is true.

Create a program that checks if a number is either negative or even:

num := 6
isNegative := num < 0
isEven := num%2 == 0
result := isNegative || isEven
fmt.Println(result)
This outputs:

true
The result is true because num is even, even though it's not negative. Only one condition needs to be true for the OR operator to return true.

Challenge

In this challenge, you'll practice using the logical OR operator (||) in Go.

You have variables representing whether a student has completed their homework and whether they've completed their extra credit assignment. A student passes the class if they've completed either their homework or their extra credit assignment (or both).

// Your task is to use the logical OR operator to determine if the student passes the class.*/

package main
import "fmt"

func main(){
	homeworkComplete := false
	extraCreditComplete := true

	passesClass := homeworkComplete || extraCreditComplete
	
	fmt.Println("Homework complete:", homeworkComplete)
	fmt.Println("Extra credit complete:", extraCreditComplete)
	fmt.Println("Student passes class:", passesClass)
}