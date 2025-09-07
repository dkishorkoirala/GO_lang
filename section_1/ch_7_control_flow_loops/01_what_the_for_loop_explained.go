/*
What The `for` Loop Explained


The for loop is Go's only looping construct, used to repeat a block of code multiple times. Let's explore how it works step by step.

First, let's create a simple program that counts from 1 to 5 using a for loop:

for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
After executing this code, the numbers 1 through 5 will show on the screen, one per line:

1
2
3
4
5
Let's understand what's happening in our for loop:

The for loop has three main parts:

1. i := 1 - This is the initialization part where we create a variable and set its starting value

2. i <= 5 - This is the condition that's checked before each loop iteration; the loop continues as long as this is true

3. i++ - This is the post statement that runs after each loop iteration (it adds 1 to i)

The loop works by first initializing i to 1, then checking if i is less than or equal to 5. If true, it runs the code inside the loop (printing the value), then executes i++ to increase the value by 1. This process repeats until the condition becomes false (when i becomes 6).

Challenge

In this challenge, you'll practice using a basic for loop to count and print numbers.

Your task is to write a for loop that counts from 1 to 10 and prints each number on a new line.

This will help you understand how to use the for loop for simple counting tasks.
*/

package main
import "fmt"

func main(){
	for i := 1; i <=10;i++{
		fmt.Println(i)
	}
}