/*
Zero Values


In Go, zero values are default values assigned to variables when declared without an explicit initial value. Let's explore these default values for different data types.

First, let's declare variables of different types without initializing them:

var number int
var decimal float64
var isValid bool
var name string

fmt.Println(number)
fmt.Println(decimal)
fmt.Println(isValid)
fmt.Println(name)
After executing this code, we'll see the zero values for each type will show on the screen:

0
0
false

Notice that int and float64 have a zero value of 0, bool has a zero value of false, and string has a zero value of an empty string (which appears as a blank line in the output).

Challenge

In this challenge, you'll explore Go's zero values - the default values that variables get when they're declared but not explicitly initialized.

Your task is to declare four variables (one for each basic type: int, float64, bool, and string) without initializing them, then print their zero values.

The code template already has the variable declarations set up for you. You just need to print each variable to see its zero value.
*/

package main
import "fmt"

func main(){
	var number int
	var decimal float64
	var isTrue bool
	var text string

	fmt.Println(number)
	fmt.Println(decimal)
	fmt.Println(isTrue)
	fmt.Println(text)
}