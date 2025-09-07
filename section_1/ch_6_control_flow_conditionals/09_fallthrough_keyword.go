/*The `fallthrough` Keyword


In Go, the fallthrough keyword allows us to continue execution to the next case in a switch statement, even after a match is found. Let's explore how this works.

First, let's create a simple switch statement that uses the fallthrough keyword to describe a number:

num := 15
    
switch {
case num < 20:
    fmt.Print("Less than 20 ")
    fallthrough
case num > 10:
    fmt.Print("Greater than 10 ")
    fallthrough
case num > 0:
    fmt.Print("Positive number")
}
After executing this code, we'll see all three messages printed on the screen:

Less than 20 Greater than 10 Positive number
This happens because the first case matches (15 is less than 20), and the fallthrough keyword forces Go to also execute the second case. Then the second case also uses fallthrough, which makes Go execute the third case as well. Without fallthrough, only the first matching case would execute.

Remember that fallthrough will execute the next case regardless of whether its condition is true or not. In our example, all conditions happen to be true for the number 15, but the next case would execute even if its condition was false.

Challenge

In this challenge, you'll practice using the fallthrough keyword in a switch statement. The fallthrough keyword allows execution to continue to the next case even after a match is found.

You have a variable representing a day of the week. Your task is to add the fallthrough keyword in the appropriate place to make the program print both 'It's a weekday!' and 'Time to work!' when the day is 'Monday'.*/


package main

import "fmt"

func main() {
	// A variable representing a day of the week
	day := "Monday"
	
	// Switch statement to check the day
	switch day {
	case "Monday":
		fmt.Println("It's a weekday!")
		fallthrough
		// TODO: Add the fallthrough keyword here to continue to the next case
	case "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Time to work!")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Invalid day!")
	}
}
