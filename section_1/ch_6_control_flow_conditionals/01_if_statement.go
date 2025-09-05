/*
The `if` Statement

The if statement lets your program make decisions based on conditions. It executes code only when a condition is true.

Let's create a simple program that uses an if statement to check someone's age:

age := 18
    
if age >= 18 {
    fmt.Println("You are an adult")
}
After executing this code, the program checks if the age variable is greater than or equal to 18. Since 18 is equal to 18, the condition is true, and the message "You are an adult" will show on the screen.

The code inside the curly braces only runs when the condition evaluates to true. If the age was less than 18, nothing would be printed.

Challenge

In this challenge, you'll practice using an if statement in Go. We have a variable temperature with a value of 25. Your task is to write an if statement that checks if the temperature is greater than 20. If it is, print "It's warm today!"
*/

package main
import "fmt"

func  main()  {
	temperature := 25

	if temperature >20{
		fmt.Println("It's warm today!")
	}
	fmt.Println("Weather check complete.")
}