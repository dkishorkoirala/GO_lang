/*Floating-Point Numbers


Let's learn about floating-point numbers in Go, which we use to store decimal values.

First, let's declare a floating-point variable with an explicit type:

var price float64 = 19.99
fmt.Println("Price:", price)
After executing this code, you'll see the following on the screen:

Price: 19.99
We've created a variable that can hold decimal numbers.

In Go, there are two floating-point types: float32 and float64. Let's see how to use float32:

var temperature float32 = 23.5
fmt.Println("Temperature:", temperature)
After executing this code, you'll see the following on the screen:

Temperature: 23.5
The difference between float32 and float64 is precision and range. float64 uses 64 bits and provides higher precision (about 15-17 decimal digits) while float32 uses 32 bits with less precision (about 6-9 decimal digits). For most applications, float64 is recommended as it's more precise.

Now, let's use Go's type inference to create a floating-point number more simply:

temperature := 23.5 // Go automatically uses float64
fmt.Println("Temperature:", temperature)
After executing this code, you'll see the following on the screen:

Temperature: 23.5
When you include a decimal point in a number, Go automatically treats it as a float64 type.

Let's try some basic math with floating-point numbers:

radius := 5.0
pi := 3.14159
area := pi * radius * radius
fmt.Println("Circle area:", area)
After executing this code, you'll see the following on the screen:

Circle area: 78.53975
Floating-point numbers allow us to perform calculations with decimal values.

Challenge

In this challenge, you'll practice working with floating-point numbers in Go.

Floating-point numbers are used to represent decimal values in Go. There are two main types: float32 and float64, with float64 being more commonly used as it provides greater precision.

Your task is to declare a floating-point variable and print it to the console.
*/

package main
import "fmt"

func main(){
	pi:= 3.14159
	temperature := 72.5

	height := 5.9
	fmt.Println("Pi approximately equals:", pi)
	fmt.Println("The temperature is:", temperature)
	fmt.Println("The height is:", height)
}