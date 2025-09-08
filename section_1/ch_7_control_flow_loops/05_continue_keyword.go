/*
The `continue` Keyword


The continue keyword allows us to skip the current iteration of a loop and move directly to the next iteration. This is useful when we want to skip certain values in our loop processing.

Let's create a program that prints only odd numbers from 1 to 10:

for i := 1; i <= 10; i++ {
    if i%2 == 0 {
        continue
    }
    fmt.Println(i)
}
After executing this code, only the odd numbers will show on the screen:

1
3
5
7
9
What happened here? Our loop runs from 1 to 10. For each number, we check if it's even by using i%2 == 0 (which tests if dividing by 2 leaves no remainder). When we find an even number, the continue statement immediately jumps to the next iteration, skipping the fmt.Println(i) line. This is why only odd numbers appear in our output.

Challenge

In this challenge, you'll practice using the continue keyword in a for loop. The program should print all numbers from 1 to 10, except for numbers that are divisible by 3 (3, 6, 9). Use the continue keyword to skip these numbers.
*/

package main
import "fmt"

func main(){
	for i:= 1; i<=10;i++{
		if i %3 ==0{
			continue
		}
		fmt.Println(i)
	}
}