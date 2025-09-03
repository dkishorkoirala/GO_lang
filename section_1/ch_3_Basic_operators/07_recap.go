/*
Recap - Calculations

Challenge

Complete the code to perform various arithmetic operations and display the results.
*/

package main
import "fmt"

func main(){
	a := 10
	b := 3

	sum := a + b 
	difference := a - b 
	product := a * b 
	quotient := a/b 
	remainder := a % b

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)
}