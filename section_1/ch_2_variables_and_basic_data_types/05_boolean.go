/*
Booleans


Booleans in Go store true/false values using the bool data type.

Let's declare a boolean variable using the var keyword with an explicit type:

isActive := true
fmt.Println("Is active:", isActive)
After executing this code, the value Is active: true will show on the screen. We've created a boolean variable that holds the value true.

Now, let's create another boolean variable with a false value:

isComplete := false
fmt.Println("Is complete:", isComplete)
After executing this code, the value Is complete: false will show on the screen. This demonstrates that boolean variables can hold either true or false values.

Let's see what happens when we print both boolean variables together:

isActive := true
isComplete := false
fmt.Println("Is active:", isActive)
fmt.Println("Is complete:", isComplete)
After executing this code, we'll see both values displayed one after another on the screen:

Is active: true
Is complete: false

Challenge

In this challenge, you'll work with boolean values in Go. Boolean values can only be true or false.

We've started with two boolean variables: likesGo and firstProgram. Your task is to:

Create a new boolean variable called isLearning and set it to true
Create another boolean variable called isExpert and set it to false
Then print all four boolean variables to see their values.
*/

package main
import "fmt"

func main(){
	likesGo := true
	firstProgram := false

	isLearning := true
	isExpert := false

	fmt.Println("Likes Go:", likesGo)
	fmt.Println("First program:", firstProgram)
	fmt.Println("Is learning:", isLearning)
	fmt.Println("Is expert:", isExpert)
}