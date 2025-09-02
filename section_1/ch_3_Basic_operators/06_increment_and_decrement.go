/*Increment and Decrement


Increment and decrement operators provide a shorthand way to add or subtract 1 from a variable.

Go uses ++ to increment and -- to decrement a variable by 1:

count := 5
count++  // Same as: count = count + 1
fmt.Println(count)

count--  // Same as: count = count - 1
fmt.Println(count)
Output:

6
5
Important: In Go, ++ and -- are statements, not expressions. They can only be used as standalone lines, not within other expressions.

Challenge

In this challenge, you'll practice using increment (++) and decrement (--) operators in Go.

We have a variable counter with an initial value. Your task is to:

Increment the counter twice using the ++ operator
Decrement the counter once using the -- operator
The program will then print the final value of the counter.
*/
package main
import "fmt"

func main(){
	counter := 5
	counter++
	counter++
	counter--
	fmt.Println("Final counter value:",counter)
}