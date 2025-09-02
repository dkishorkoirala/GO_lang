/*
Augmented Assignment Operators


Augmented assignment operators combine an arithmetic operation with assignment. They're shortcuts that make your code cleaner.

Instead of writing:

score := 100
score = score + 50
Use the += operator:

score := 100
score += 50
fmt.Println(score)
Output:

150
Other augmented operators include -=, *=, /=, and %=:

count := 10
count *= 5  // Same as: count = count * 5
fmt.Println(count)
Output:

50

Challenge

In this challenge, you'll practice using augmented assignment operators in Go. Augmented assignment operators combine an arithmetic operation with assignment (like +=, -=, *=, /=, %=).

We have a variable score that starts at 10. Your task is to use augmented assignment to increase the score by 5.
*/
package main
import "fmt"

func main(){
	score := 10
	score += 5
	fmt.Println("Final score:", score)
}