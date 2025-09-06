/*
The `switch` Statement


Let's learn about the switch statement in Go, which provides a cleaner way to check multiple conditions compared to using many if/else if statements.

First, let's create a simple program that evaluates a day number and prints the corresponding day name:

day := 3
    
switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    fmt.Println("Wednesday")
default:
    fmt.Println("Another day")
}

After executing this code, the program will check the value of day (which is 3) and match it with the appropriate case. Since day equals 3, the program will execute the code under case 3 and show on the screen:

Wednesday
The switch statement checks each case in order until it finds a match. If no match is found, it executes the code under the default case. This makes your code much cleaner than writing multiple if/else if statements.

Challenge

In this challenge, you'll practice using the switch statement in Go. You'll be given a variable representing a day of the week, and you need to complete the switch statement to display whether it's a weekday or weekend.
*/

package main
import "fmt"

func main(){
	day := "Saturday"

	switch day{
	case "Sunday", "Saturday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday")
	}
}