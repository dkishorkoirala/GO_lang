/*
Println


Println is a function from the fmt package that prints text followed by a line break. You can print multiple items with a single call.

Print a simple message:

fmt.Println("Learning Go is fun!")
When you run this code, you'll see:

Learning Go is fun!
Print multiple values separated by commas:

fmt.Println("Go", "is", "awesome")
The output shows all items with spaces between them:

Go is awesome

Challenge

In this challenge, you'll practice using fmt.Println() to display multiple values on the same line.

Your task is to use fmt.Println() to print the text Alex 25 coding on a single line.

Remember that fmt.Println() can accept multiple arguments, and it will automatically add spaces between them.
*/

package main
import ("fmt")

func main(){
	fmt.Println("Alex", "25", "coding")
}