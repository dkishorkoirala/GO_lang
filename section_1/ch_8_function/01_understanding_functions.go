/*Understanding Functions


Functions are reusable blocks of code that perform specific tasks. They help organize code and avoid repetition.

Let's create a simple function that adds two numbers:

func add(x int, y int) int {
    return x + y
}

func main() {
    result := add(5, 3)
    fmt.Println("5 + 3 =", result)
}
After executing this code, the number 8 will show on the screen as the result of adding 5 and 3:

5 + 3 = 8
Let's break down how to define a function:

First, we declare the function with a name and parameters:

func add(x int, y int) int {
This creates a function named add that takes two integer parameters and returns an integer value.

Next, let's write the function body that performs the calculation:

return x + y
This simple line adds the two parameters together and sends the result back to wherever the function was called.

Now, let's see how to call our function and use the result:

result := add(5, 3)
This line calls our add function with the values 5 and 3, and stores the returned value (8) in a new variable called result.

Finally, let's display the result to the user:

fmt.Println("5 + 3 =", result)
This prints the text "5 + 3 =" followed by the value stored in result, which is 8.

Challenge

In this challenge, you'll practice creating and using a simple function in Go.

Your task is to complete the greet function that takes a name as a parameter and returns a greeting message.

The function should return a string in the format: "Hello, " + name + "!"

The function should return a string in the format: "Hello, Gopher!"
*/

package main
import "fmt"

func greet(name string) string{
	return "Hello, " + name + "!"
}

func main(){
	name := "Gopher"
	message := greet(name)
	fmt.Println(message)
}