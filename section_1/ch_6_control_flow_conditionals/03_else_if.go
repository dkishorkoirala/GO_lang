/*
The `else if` Keyword


# The `else if` Keyword
Let's learn how to use the else if keyword in Go to check multiple conditions in sequence. This is useful when you need more than two possible outcomes in your program.

First, let's create a program that categorizes temperature based on different ranges:

temp := 75
    
if temp < 50 {
    fmt.Println("It's cold")
} else if temp < 80 {
    fmt.Println("It's pleasant")
} else {
    fmt.Println("It's hot")
}
After executing this code, the program will show "It's pleasant" on the screen. Here's why: the program first checks if temp (75) is less than 50. Since this is false, it moves to the next condition and checks if temp is less than 80. Since 75 is indeed less than 80, it executes that code block and prints "It's pleasant". The final else block is skipped because a condition was already met.

The conditions are checked in order from top to bottom, and only the code for the first matching condition runs. Once a condition is met, the remaining conditions are not checked at all.

Challenge

In this challenge, you'll practice using the else if keyword to handle multiple conditions.

You have a variable weather that contains a string describing the current weather. Your task is to complete the code to print different messages based on the weather condition:

If it's "sunny", print "Wear sunscreen!"
If it's "rainy", print "Bring an umbrella!"
If it's "cold", print "Wear a jacket!"
For any other weather condition, print "Check the forecast!"
*/

package main 
import "fmt"

func main(){
	weather := "rainy"

	if weather == "sunny"{
		fmt.Println("Wear sunscreen!")
	} else if  weather == "rainy"{
		fmt.Println("Bring an umbrella!")
	} else{
		fmt.Println("Check the forecast!")
	}
}