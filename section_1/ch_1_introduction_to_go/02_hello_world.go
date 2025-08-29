/*
Hello World


The "Hello, World!" is a simple program that outputs Hello, World! to the screen.

In Go, we use Println() to show words on the screen. The words go inside quotation marks.

Let's take a look at the "Hello World!" program in Go:

fmt.Println("Hello, World!")

Challenge

Use the code view to write a program that outputs Hello, World!

Note that anything inside quotation marks is case sensitive. For example:

fmt.Println("Hello, World!")
fmt.Println("hello, world!")
are different things (notice the capital letters in the first line).
*/

package main
import ( "fmt" )

func main(){
	fmt.Println("Hello, World!")
}