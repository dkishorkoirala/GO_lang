/*Function Parameters


Function parameters are values you pass to a function when calling it. They allow functions to work with different data each time they're called.

Create a function with parameters by listing their names and types in the parentheses:

func add(x int, y int) {
    sum := x + y
    fmt.Println("Sum:", sum)
}

func main() {
    add(5, 3)
    add(10, 20)
}
When you run this program, it outputs:

Sum: 8
Sum: 30
Parameters with the same type can be written more concisely:

func greet(first, last string) {
    fmt.Println("Hello,", first, last)
}
Functions can also return values by specifying a return type:

func multiply(x, y int) int {
    return x * y
}

func main() {
    result := multiply(4, 5)
    fmt.Println("Result:", result) // Output: Result: 20
}

Challenge

In this challenge, you'll practice working with function parameters. You need to complete a function that takes two string parameters and combines them to create a greeting message.

The function createGreeting is already defined, but it's missing the parameters. Your task is to add the correct parameters and use them in the function body.*/

package main
import "fmt"

func createGreeting(greeting, name string)string{
	return greeting + ", " + name + "!"
}


func main() {
	message := createGreeting("Hello", "Gopher")
	fmt.Println(message)
	
	message = createGreeting("Welcome", "Friend")
	fmt.Println(message)
}