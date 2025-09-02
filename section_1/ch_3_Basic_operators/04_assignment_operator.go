/*Assignment Operator


The assignment operator (=) assigns a value to a variable. It's how we store new values in existing variables.

Create a variable and assign a value:

var score int
score = 100
fmt.Println(score)
Output:

100
You can also reassign values to change a variable:

score = 150
fmt.Println("Updated score:", score)
Output:

Updated score: 150

Challenge

In this challenge, you'll practice using the assignment operator (=) in Go.

We've created some variables for you. Your task is to use the assignment operator to change the value of the favoriteColor variable to "blue" and then print it out.*/

package main 
import "fmt"

func main(){
	favoriteColor := "red"
	favoriteColor = "blue"
	fmt.Println("My favorite color is:", favoriteColor)
}