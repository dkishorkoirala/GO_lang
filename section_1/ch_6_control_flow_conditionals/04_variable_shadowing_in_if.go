/*Variable Shadowing in `if`


Let's learn about variable shadowing in if statements. This allows you to create a variable that only exists within the if block and its related else blocks.

First, let's create a variable outside our conditional statement:

x := 10
fmt.Println("Original x value:", x)
After executing this code, we'll have a variable x with the value 10 that we can use throughout our program.

Now, let's use variable shadowing in an if statement to create a new variable that only exists within that block:

if y := x * 2; y > 15 {
    fmt.Println("y is greater than 15:", y)
    // y only exists in this block
} else {
    fmt.Println("y is not greater than 15:", y)
    // y is also available in else blocks
}
After executing this code, Go will calculate y as 20 (which is x * 2), check if it's greater than 15, and then execute the first block. The output will show on the screen:

y is greater than 15: 20
Finally, let's try to access our variables after the if statement:

// fmt.Println(y) would cause an error - y doesn't exist here
fmt.Println("x is still:", x)
After executing this code, we can still access x, but y is no longer available because it only existed within the if and else blocks. The output will show on the screen:

x is still: 10
This demonstrates how variable shadowing works in if statements - the variable y is only accessible within the conditional blocks where it was declared.

Challenge

In this challenge, you'll practice variable shadowing in an if statement. You'll see how a variable declared inside an if block can have the same name as a variable in the outer scope, but with a different value.

The code has a variable message in the outer scope. Your task is to create a shadowed variable with the same name inside the if block and then print both variables to see the shadowing effect.
*/

package main 
import "fmt"

func main(){
	message := "Hello from outer scope"
    isLoggedIn := true

    if isLoggedIn {
        message := "Hello from inner scope"
        fmt.Println(message)
    }

    fmt.Println(message)
}