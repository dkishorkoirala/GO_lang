/*Declaring a Function


Let's learn how to declare a function in Go that takes multiple parameters of the same type.

First, let's create a function that adds three integers together:

func addThree(x, y, z int) int {
    return x + y + z
}

func main() {
    sum := addThree(5, 10, 15)
    fmt.Println("The sum is:", sum)
}
After executing this code, the value 30 (which is 5 + 10 + 15) will be calculated and stored in the sum variable. Then the message "The sum is: 30" will show on the screen:

The sum is: 30
Notice that we wrote x, y, z int instead of writing x int, y int, z int. This is a helpful shorthand in Go when multiple parameters share the same type - it makes our code shorter and cleaner.

Challenge

In this challenge, you'll practice declaring a function in Go. Your task is to complete the greet function that takes a name as a parameter and returns a greeting message.

The function should return a string in the format: "Hello, [name]!"*/

package main
import "fmt"

func greet(name string) string{
	return "Hello, " + name + "!"
}

func main(){
	name1 := "Alice"
    name2 := "Bob"

	message1:= greet(name1)
	message2:= greet(name2)

	fmt.Println(message1)
	fmt.Println(message2)
}