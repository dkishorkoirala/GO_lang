/*
Nested Loops


Nested Loops are loops placed inside another loop. The inner loop runs completely for each iteration of the outer loop.

Let's create a multiplication table for numbers 1-3 using nested loops:

for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Printf("%d×%d=%d\t", i, j, i*j)
    }
    fmt.Println()
}
After executing this code, we get a multiplication table where each row represents one value of the outer loop variable (i) multiplied by all values of the inner loop variable (j). The output will show on the screen:

1×1=1   1×2=2   1×3=3   
2×1=2   2×2=4   2×3=6   
3×1=3   3×2=6   3×3=9   
Let's understand how this works step by step:

First, the outer loop sets i=1 and the inner loop runs completely:

i=1, j=1: prints "1×1=1"
i=1, j=2: prints "1×2=2"
i=1, j=3: prints "1×3=3"
After the inner loop finishes, we print a new line with fmt.Println().

Then the outer loop sets i=2 and the inner loop runs again:

i=2, j=1: prints "2×1=2"
i=2, j=2: prints "2×2=4"
i=2, j=3: prints "2×3=6"
Finally, the outer loop sets i=3 and the inner loop runs one last time:

i=3, j=1: prints "3×1=3"
i=3, j=2: prints "3×2=6"
i=3, j=3: prints "3×3=9"
This pattern of loops within loops is very useful for working with tables of data or creating grid-like structures.

Challenge

In this challenge, you'll practice using nested loops to print a simple pattern. The outer loop will iterate through rows, and the inner loop will iterate through columns to create a small grid of characters.

Complete the nested loops to print a 3x3 grid of asterisks (*) as shown in the expected output.
*/

package main 
import "fmt"

func main()  {
	row := 3
	cols := 3

	for i :=0; i< row; i++{
		for j := 0; j < cols; j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}