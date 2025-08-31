/*
Integers (int)


Integers (int) are whole numbers without decimal points. Go uses int to store these values.

Let's create integer variable using the var keyword:

var age int = 25
fmt.Println("Age:", age)
After executing this code, you will see Age: 25 displayed on the screen. We've created an integer variable named age with the value 25.

Now, let's create an integer using Go's short declaration syntax with :=:

score := 95
fmt.Println("Score:", score)
After executing this code, you will see Score: 95 displayed on the screen. The := operator automatically determines that score should be an integer type based on the value 95.

Challenge

In this challenge, you'll practice working with integer variables in Go.

Integers (int) are whole numbers without decimal points. They can be positive, negative, or zero.

Your tasks:

Create an integer variable named age with a value of 25
Create another integer variable named year with a value of 2023
Create a third integer variable named temperature with a value of -5
Print all three variables on separate lines using fmt.Println()
*/


package main

import ("fmt")

func main(){
	var age int = 25
	year := 2023
	temperature := -5
	fmt.Println(age)
	fmt.Println(year)
	fmt.Println(temperature)
}