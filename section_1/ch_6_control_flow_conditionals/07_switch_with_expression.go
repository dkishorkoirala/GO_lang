/*
Switch with Expressions


Let's learn about switch expressions in Go, which allow us to evaluate a value directly in the switch statement.

First, let's see how to use a switch statement with an expression to check if a number is positive, negative, or zero:

num := -5
switch {
case num > 0:
    fmt.Println("Positive number")
case num < 0:
    fmt.Println("Negative number")
default:
    fmt.Println("Zero")
}
After executing this code, the program will show on the screen:

Negative number
Notice that we're using switch { with no value after it. This means Go will evaluate each case expression as a boolean condition. Since num is -5, the second case num < 0 evaluates to true, so "Negative number" is printed.

We can also initialize a variable right in the switch statement. Let's see how that works:

switch num := -5; {
case num > 0:
    fmt.Println("Positive number")
case num < 0:
    fmt.Println("Negative number")
default:
    fmt.Println("Zero")
}
After executing this code, we get the same result:

Negative number
In this example, we declared and initialized num directly in the switch statement. The semicolon separates the initialization from the switch expression (which is empty here). This makes our code more compact while keeping the same functionality.

Challenge

In this challenge, you'll practice using a simple switch expression in Go. You need to complete the code to determine the description of a day based on its name using a switch statement.

A variable dayName is already defined with the value "Wednesday". Your task is to use a switch expression to set the description variable to:
- "Start of the week" for Monday
- "Midweek" for Wednesday
- "Almost weekend" for Friday
- "Weekend" for Saturday or Sunday
- "Regular day" for any other day
*/

package main
import "fmt"

func main(){
	dayName := "Wednesday"

	var description string

	switch {
	case dayName == "Monday":
		description = "Start of the week"
	case dayName == "Wednesday":
		description = "Midweek"
	case dayName=="Friday":
		description = "Almost weekend"
	case dayName == "Saturday" || dayName == "Sunday":
		description = "Weekend"
	default:
		description = "Regular day"
	}
	fmt.Printf("%s is %s\n", dayName, description)
}