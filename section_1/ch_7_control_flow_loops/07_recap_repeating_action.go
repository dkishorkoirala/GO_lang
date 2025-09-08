/*Recap - Repeating Actions


Challenge

Complete the code to implement different types of loops and print the results.*/

package main
import "fmt"

func main()  {
	sum := 0
	for i := 1; i <=5;i ++{
		fmt.Println(i)
		sum += i
	}
    
	fmt.Printf("\nSum: %d\n\n",sum)

	for j := 10; j >= 1; j--{
		fmt.Println(j)
	}
}