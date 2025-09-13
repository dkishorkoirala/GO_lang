/*Function Scope Basics


In Go, variables declared inside a function are only accessible within that function. This is called function scope.

Create a function with a local variable:

func calculateTotal(price int, quantity int) int {
    tax := 0.1 * float64(price * quantity)  // local variable
    total := price * quantity + int(tax)
    return total
}

func main() {
    result := calculateTotal(10, 2)
    fmt.Println("Total:", result)
    // fmt.Println(tax)  // Error: tax is not accessible here
}
The output shows only the returned value:

Total: 22
The variable tax only exists inside the calculateTotal function and cannot be accessed from main.

Challenge

In this challenge, you'll practice understanding function scope in Go. You need to fix a simple program that has a scope issue.

The program has a global variable message and a function that tries to modify it. Your task is to fix the code so that the function properly updates the global variable.*/

package main
import "fmt"

var message string = "Hello from global scope"

func updateMessage() {
    message = "Hello from function scope"
    fmt.Println("Inside function:", message)
}

func main() {
    fmt.Println("Before function call:", message)
    updateMessage()
    fmt.Println("After function call:", message)
}