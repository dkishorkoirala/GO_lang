/*The `break` Keyword


The break keyword allows us to immediately exit a loop, even if the loop condition would normally continue running.

Let's create a program that counts from 1 to 10 but stops when it reaches 5:

for i := 1; i <= 10; i++ {
    if i > 5 {
        break
    }
    fmt.Println(i)
}
fmt.Println("Loop exited")
After executing this code, the numbers 1 through 5 will show on the screen, followed by the message "Loop exited":

1
2
3
4
5
Loop exited
The break statement causes the loop to immediately stop when i becomes greater than 5. This means the loop never prints the numbers 6 through 10. Instead, it jumps directly to the code after the loop and prints "Loop exited".

Challenge

In this challenge, you'll practice using the break keyword to exit a loop early.

You have a sequence of numbers from 1 to 10. Your task is to find the first number that is divisible by 3, print it, and then immediately exit the loop using break.

This way, once you find what you're looking for, you don't waste time continuing to search through the rest of the numbers.*/

package main
import "fmt"

func main(){

	for i := 1; i <= 10; i++{
		fmt.Printf("Checking number: %d\n",i)

		if i %3 == 0{
			fmt.Printf("Found it: %d!\n",i)
			break
		}
	}

	fmt.Println("Search complete")
}