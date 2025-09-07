/*
Switch without Expression


In Go, we can use a switch statement without an expression after the keyword. This is useful when we want to evaluate multiple conditions in a clean way.

Let's create a program that categorizes a person based on their age:

age := 25
    
switch {
case age < 13:
    fmt.Println("Child")
case age < 20:
    fmt.Println("Teenager")
case age < 30:
    fmt.Println("Young Adult")
default:
    fmt.Println("Adult")
}
After executing this code, the program will check each case condition in order. Since our age is 25, it's not less than 13 or 20, but it is less than 30, so the third case matches. The text "Young Adult" will show on the screen:

Young Adult
Notice that without an expression after switch, each case becomes a boolean condition that evaluates to either true or false. The first case that evaluates to true will execute its code block.

If none of the cases match, the default case would execute. In our example, if age was 35, the output would be "Adult".

Challenge

In this challenge, you'll practice using a switch statement without an expression. This type of switch is useful when you want to evaluate multiple conditions without repeating if-else statements.

We've provided a variable timeOfDay that represents the current hour (in 24-hour format). Your task is to complete the switch statement without an expression to print an appropriate greeting based on the time of day.
*/

package main
import "fmt"

func main(){
	timeOfDay := 14

	switch {
	case timeOfDay < 12:
		fmt.Println("Good morning!")
	case timeOfDay < 23:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Hello!")
	}
}