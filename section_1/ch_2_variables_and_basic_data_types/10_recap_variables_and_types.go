/*
Recap - Variables and Types
Challenge

In this challenge, you'll practice working with different variable types in Go and printing them to the console.

You need to:

Define a new integer variable named 'quantity' with value 5
Print the product name with a label
Print the quantity with a label
Make sure to use the variables in your print statements rather than hardcoding values where possible.

Expected output format:

Product: Coffee Mug
Quantity: 5
*/

package main
import "fmt"

func main(){
	productName := "Coffee Mug"
	quantity := 5

	fmt.Println("Product:", productName)
	fmt.Println("Quantity:", quantity)

}