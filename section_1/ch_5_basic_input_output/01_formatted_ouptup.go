/*
Formatted Output


Let's learn how to format our output in Go using the fmt.Printf() function. This function allows us to display values in specific formats using placeholders.

First, let's create a program that formats different data types using placeholders:

name := "Gopher"
age := 5
height := 1.5
    
fmt.Printf("Name: %s, Age: %d, Height: %.1f\n", name, age, height)
After executing this code, the text "Name: Gopher, Age: 5, Height: 1.5" will show on the screen. The placeholders in our string are replaced with the actual values of our variables.

Let's understand what each placeholder does:

// %s is replaced with a string value (name)
// %d is replaced with an integer value (age)
// %.1f is replaced with a float value showing 1 decimal place (height)
// %t is replaced with a boolean value (true or false)
The fmt.Printf() function replaces each placeholder with the corresponding variable that we provide after the format string. The order of variables must match the order of placeholders.

Now let's try another example with different types of formatting:

pi := 3.14159
percentage := 0.75
letter := 'A'

fmt.Printf("Pi: %.2f, Percentage: %.1f%%, Letter: %c\n", pi, percentage*100, letter)
After executing this code, the text "Pi: 3.14, Percentage: 75.0%, Letter: A" will show on the screen. Notice how %.2f shows only 2 decimal places for pi, and we multiply the percentage by 100 to show it as a percentage value.

Challenge

In this challenge, you'll practice using formatted output in Go with the `fmt.Printf()` function.

You have information about a product that needs to be displayed in a specific format using format verbs. The variables are already defined for you.

Your task is to use `fmt.Printf()` with the appropriate format verbs to display the product information:
- Use `%s` for strings
- Use `%d` for integers  
- Use `%.2f` for floating-point numbers with 2 decimal places
- Use `%t` for boolean values

Display the product information exactly as shown in the expected output.
*/

package main 
import "fmt"

func main (){

	productName := "Wireless Headphones"
	productPrice := 79.99
	productInStock := true
	productID := 12345

	fmt.Printf("Product: %s (ID: %d)\n", productName, productID)
	fmt.Printf("Price: $%.2f\n", productPrice)
	fmt.Printf("In Stock: %t\n", productInStock)
}