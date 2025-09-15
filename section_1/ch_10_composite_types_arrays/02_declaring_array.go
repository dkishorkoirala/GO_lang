/*Declaring Arrays


In Go, you can declare arrays in several ways. An array's size is part of its type and must be specified when declared.

Declare an array using var with a specific size:

var scores [4]int
fmt.Println(scores)
This creates an array of 4 integers, all initialized to zero:

[0 0 0 0]
Declare and initialize an array in one line using shorthand notation:

names := [3]string{"Alice", "Bob", "Charlie"}
fmt.Println(names)
This displays:

[Alice Bob Charlie]

Challenge

In this challenge, you'll practice declaring and initializing an array in Go. Your task is to declare an array of 5 integers with the values 10, 20, 30, 40, and 50, then print the array to the console.
*/
package main
import "fmt"

func main(){
	numbers := [5]int {10,20,30,40,50}
	fmt.Println("My array of numbers:", numbers)
}