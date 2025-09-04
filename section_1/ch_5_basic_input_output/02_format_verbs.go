/*
Format Verbs


Let's learn about format verbs in Go. Format verbs are special placeholders that we use with fmt.Printf() to display different types of data in specific ways.

First, let's create some variables with different data types:

name := "Gopher"
age := 5
height := 1.5
isActive := true
After executing this code, we've created four variables: a string, an integer, a float, and a boolean.

Now, let's use the default format verb %v to display our string variable:

fmt.Printf("Default: %v\n", name)
After executing this code, Default: Gopher will show on the screen. The %v is replaced with the value of name.

Next, let's use the %d format verb specifically for integers:

fmt.Printf("Integer: %d\n", age)
After executing this code, Integer: 5 will show on the screen. The %d is perfect for displaying whole numbers.

Let's display our float value using the %f format verb:

fmt.Printf("Float: %f\n", height)
After executing this code, Float: 1.500000 will show on the screen. Notice that %f shows six decimal places by default.

We can control the number of decimal places by adding a precision specifier:

fmt.Printf("Float (2 decimals): %.2f\n", height)
After executing this code, Float (2 decimals): 1.50 will show on the screen. The .2 before f limits the output to 2 decimal places.

Finally, let's display our boolean value using the %t format verb:

fmt.Printf("Boolean: %t\n", isActive)
After executing this code, Boolean: true will show on the screen. The %t is specifically for displaying boolean values.

Remember: use %v for any type, %d for integers, %f for floats, and %t for booleans.

Challenge

In this challenge, you'll practice using format verbs with fmt.Printf. You have several variables of different types already defined. Your task is to print them using the appropriate format verbs (%v, %d, %f, %s, %t, etc.).
*/

package main 
import "fmt"

func main(){
	age := 25
	height := 5.9
	isStudent := true
	name := "Alex"

	fmt.Printf("Name: %v, Age: %d, Height: %f, Student: %t", name, age, height, isStudent)
}