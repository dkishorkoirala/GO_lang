/*
For Loop - Basic


Let's learn about the basic for loop in Go, focusing on the initialization part. A for loop helps us repeat code multiple times.

First, let's create a simple loop that counts from 1 to 3:

for i := 1; i <= 3; i++ {
    fmt.Println(i)
}
After executing this code, the numbers 1, 2, and 3 will show on the screen one after another:

1
2
3
The initialization part i := 1 happens only once before the loop starts. It creates a variable i and sets it to 1. This variable i is only available inside the loop.

The middle part i <= 3 is the condition that checks if we should continue looping. The last part i++ increases the value of i by 1 after each loop.

Challenge

In this challenge, you'll practice using a for loop with initialization, condition, and post statement. Your task is to complete the for loop to print the numbers from 1 to 5.

The for loop should:
1. Initialize a variable i to 1
2. Continue as long as i is less than or equal to 5
3. Increment i by 1 after each iteration*/

package main
import "fmt"

func main(){
	for i := 1; i <=5;i++{
		fmt.Println(i)
	}
}