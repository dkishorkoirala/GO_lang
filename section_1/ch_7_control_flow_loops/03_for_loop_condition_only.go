/*
For Loop - Condition Only


Let's learn about the condition-only for loop in Go. This type of loop works like a while loop in other languages, using just a condition without initialization or post statements.

First, let's create a simple counter program that counts from 1 to 3:

i := 1
for i <= 3 {
    fmt.Println(i)
    i++
}
After executing this code, the numbers 1, 2, and 3 will show on the screen one after another:

1
2
3
The loop works by checking if i is less than or equal to 3 before each iteration. Inside the loop, we print the current value of i and then increase it by 1 with i++. The loop continues running as long as the condition i <= 3 remains true, and stops when i becomes 4.

Unlike the standard for loop, we need to manually increment our counter variable inside the loop body, otherwise we'd create an infinite loop.

Challenge

In this challenge, you'll practice using a for loop with only a condition (no initialization or post statement). This is similar to a while loop in other languages.

We've provided a boolean variable keepRunning set to true and a counter variable count set to 0. Your task is to complete the for loop that will print numbers from 0 to 4 and then stop.

Inside the loop, we increment count and check if it reaches 5. When it does, we set keepRunning to false to exit the loop.
*/

package main
import "fmt"

func main(){
	keepRunning :=true
	count := 0
	for keepRunning{
		fmt.Println(count)
		count++

		if count >= 5{
			keepRunning = false
		}
	}
}