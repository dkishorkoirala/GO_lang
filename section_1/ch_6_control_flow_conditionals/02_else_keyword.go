/*
The `else` Keyword


The else keyword lets you specify code to run when the if condition is false. It provides an alternative execution path.

Let's create a program that checks a person's age and prints different messages based on whether they're an adult or not:

age := 15
    
if age >= 18 {
    fmt.Println("You are an adult")
} else {
    fmt.Println("You are a minor")
}
After executing this code, the program checks if age is greater than or equal to 18. Since 15 is less than 18, the condition is false, so it executes the code in the else block. The message "You are a minor" will show on the screen.

The else statement always follows an if statement and its code block runs only when the if condition evaluates to false. This gives us a way to handle both possible outcomes of a condition.


Challenge

In this challenge, you'll practice using the else keyword in Go. We have a variable isRaining that indicates whether it's raining outside. Your task is to complete the code to print a message based on this condition. If it's raining, the program should print "Bring an umbrella!". Otherwise, it should print "Enjoy the sunshine!".
*/

package main 
import "fmt"

func main(){
	isRaining := true

	if isRaining{
		fmt.Println("Bring an umbrella!")
	} else {
		fmt.Println("Enjoy the sunshine!")
	}
}