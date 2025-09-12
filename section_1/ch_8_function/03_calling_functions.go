/*
Calling Functions


Calling Functions

Let's learn how to call functions in Go. Calling a function means executing a function that has been declared so we can use its functionality.

First, let's create a simple greeting function and call it with different names:

func greet(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    greet("Alex")
    greet("Taylor")
}
After executing this code, we'll see two greeting messages on the screen:

Hello, Alex
Hello, Taylor
Now, let's call a function that returns a value and store that result in a variable:

func addThree(a, b, c int) int {
    return a + b + c
}

func main() {
    result := addThree(10, 20, 30)
    fmt.Println(result)
}
After executing this code, the function will add the three numbers together and return the sum. The variable result will store the value 60, and then we'll see this output on the screen:

60
That's how we call functions in Go - simply write the function name followed by parentheses containing any required arguments. If the function returns a value, we can capture it in a variable.

Challenge

In this challenge, you'll practice calling functions in Go. We've defined three simple functions for you: sayHello, sayGoodbye, and sayThankYou. Your task is to call each of these functions in the correct order to produce the expected output.

Each function simply prints a greeting message. You just need to call them in the right sequence.*/

package main
import "fmt"
func sayHello() {
    fmt.Println("Hello, friend!")
}

func sayGoodbye() {
    fmt.Println("Goodbye, friend!")
}

func sayThankYou() {
    fmt.Println("Thank you, friend!")
}


func main(){
	sayHello()
	sayThankYou()
	sayGoodbye()
}